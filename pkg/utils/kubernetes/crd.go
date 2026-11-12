// SPDX-FileCopyrightText: Contributors to the Gardener project
//
// SPDX-License-Identifier: Apache-2.0

package kubernetes

import (
	"context"
	"errors"
	"fmt"
	"time"

	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"
	"k8s.io/apimachinery/pkg/runtime/serializer/json"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/rest"
	watchtools "k8s.io/client-go/tools/watch"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/gardener/gardener/pkg/utils/flow"
	"github.com/gardener/gardener/pkg/utils/kubernetes/health"
	"github.com/gardener/gardener/pkg/utils/retry"
)

var (
	// WaitTimeout specifies the total time to wait for CRDs to become ready or to be deleted. Exposed for testing.
	// While waiting for CRD readiness is parallelized (see WaitUntilCRDManifestsReady below), the controllers
	// responsible for populating the "readiness" status into the CRD only have one worker each (e.g., see
	// https://github.com/kubernetes/apiextensions-apiserver/blob/376adbc0c7f0bc548dbbf2ad7c4f3e53840aa08f/pkg/controller/establish/establishing_controller.go#L88-L89).
	// Therefore, we need to wait for a longer time here  (basically proportional to the
	// amount of CRDs) in case we create a lot of CRDs in parallel (which happens at Garden or Seed creation), since
	// they are processed sequentially.
	WaitTimeout = 2 * time.Minute
	// IntervalWait specifies the interval between tries to confirm for CRDs to become ready or to be deleted. Exposed for testing.
	IntervalWait = 1 * time.Second
)

// WaitUntilCRDManifestsReady takes names of CRDs and waits for them to get ready with a timeout, specified by `WaitTimeout`
func WaitUntilCRDManifestsReady(ctx context.Context, c client.Client, crdNames ...string) error {
	var fns []flow.TaskFn
	for _, crdName := range crdNames {
		fns = append(fns, func(ctx context.Context) error {
			timeoutCtx, cancel := context.WithTimeout(ctx, WaitTimeout)
			defer cancel()

			return retry.Until(timeoutCtx, IntervalWait, func(ctx context.Context) (done bool, err error) {
				crd := &apiextensionsv1.CustomResourceDefinition{}

				if err := c.Get(ctx, client.ObjectKey{Name: crdName}, crd); err != nil {
					if client.IgnoreNotFound(err) == nil {
						return retry.MinorError(err)
					}
					return retry.SevereError(err)
				}

				if err := health.CheckCustomResourceDefinition(crd); err != nil {
					return retry.MinorError(err)
				}
				return retry.Ok()
			})
		})
	}
	return flow.Parallel(fns...)(ctx)
}

// NewUncachedWatchClient returns a client that talks directly to the API server (no cache) and is able to watch. It
// reuses the scheme and REST mapper of the given client. It is meant to be passed to the CRD deployers, so that they
// do not need an informer for CustomResourceDefinitions and can watch them instead of polling.
func NewUncachedWatchClient(restConfig *rest.Config, c client.Client) (client.WithWatch, error) {
	return client.NewWithWatch(restConfig, client.Options{Scheme: c.Scheme(), Mapper: c.RESTMapper()})
}

// WaitUntilCRDManifestsReadyWithWatch behaves like WaitUntilCRDManifestsReady but does not poll. Instead, it opens a
// watch for each CRD (restricted to its name via a field selector) and returns once the CRD is ready. The given client
// must talk directly to the API server (i.e., it must not be backed by an informer cache).
func WaitUntilCRDManifestsReadyWithWatch(ctx context.Context, c client.WithWatch, crdNames ...string) error {
	var fns []flow.TaskFn
	for _, crdName := range crdNames {
		fns = append(fns, func(ctx context.Context) error {
			timeoutCtx, cancel := context.WithTimeout(ctx, WaitTimeout)
			defer cancel()

			return waitUntilCRDReadyWithWatch(timeoutCtx, c, crdName)
		})
	}
	return flow.Parallel(fns...)(ctx)
}

func waitUntilCRDReadyWithWatch(ctx context.Context, c client.WithWatch, crdName string) error {
	var lastErr error

	selector := client.MatchingFieldsSelector{Selector: fields.OneTermEqualSelector("metadata.name", crdName)}

	for {
		// List first (a single object thanks to the field selector) to get the current state and the resourceVersion to
		// start the watch from.
		crdList := &apiextensionsv1.CustomResourceDefinitionList{}
		if err := c.List(ctx, crdList, selector); err != nil {
			return fmt.Errorf("failed listing CRD %q: %w", crdName, err)
		}

		lastErr = apierrors.NewNotFound(apiextensionsv1.Resource("customresourcedefinitions"), crdName)
		if len(crdList.Items) > 0 {
			// short-circuit if the CRD is already ready
			if lastErr = health.CheckCustomResourceDefinition(&crdList.Items[0]); lastErr == nil {
				return nil
			}
		}

		w, err := c.Watch(ctx, &apiextensionsv1.CustomResourceDefinitionList{}, selector, &client.ListOptions{Raw: &metav1.ListOptions{ResourceVersion: crdList.ResourceVersion}})
		if err != nil {
			return fmt.Errorf("failed watching CRD %q: %w", crdName, err)
		}

		_, err = watchtools.UntilWithoutRetry(ctx, w, func(event watch.Event) (bool, error) {
			switch event.Type {
			case watch.Added, watch.Modified:
				crd, ok := event.Object.(*apiextensionsv1.CustomResourceDefinition)
				if !ok {
					return false, fmt.Errorf("unexpected object type %T", event.Object)
				}
				lastErr = health.CheckCustomResourceDefinition(crd)
				return lastErr == nil, nil
			case watch.Deleted:
				lastErr = apierrors.NewNotFound(apiextensionsv1.Resource("customresourcedefinitions"), crdName)
				return false, nil
			case watch.Error:
				return false, apierrors.FromObject(event.Object)
			}
			return false, nil
		})
		w.Stop()

		switch {
		case err == nil:
			return nil
		case errors.Is(err, watchtools.ErrWatchClosed):
			// The server closed the watch (e.g., due to a timeout), start a new one unless the context is done.
			if ctx.Err() != nil {
				return timeoutError(ctx, lastErr)
			}
		case errors.Is(err, context.DeadlineExceeded), errors.Is(err, context.Canceled), wait.Interrupted(err):
			return timeoutError(ctx, lastErr)
		default:
			return err
		}
	}
}

func timeoutError(ctx context.Context, lastErr error) error {
	return fmt.Errorf("watch failed with %w, last error: %w", ctx.Err(), lastErr)
}

// WaitUntilCRDManifestsDestroyed takes CRD names and waits for them to be gone with a timeout, specified by `WaitTimeout`.
func WaitUntilCRDManifestsDestroyed(ctx context.Context, c client.Client, crdNames ...string) error {
	var fns []flow.TaskFn

	for _, resourceName := range crdNames {
		crd := &apiextensionsv1.CustomResourceDefinition{
			ObjectMeta: metav1.ObjectMeta{
				Name: resourceName,
			},
		}

		fns = append(fns, func(ctx context.Context) error {
			timeoutCtx, cancel := context.WithTimeout(ctx, WaitTimeout)
			defer cancel()
			return WaitUntilResourceDeleted(timeoutCtx, c, crd, IntervalWait)
		})
	}
	return flow.Parallel(fns...)(ctx)
}

var (
	crdScheme *runtime.Scheme
	crdCodec  runtime.Codec
)

func init() {
	crdScheme = runtime.NewScheme()
	utilruntime.Must(apiextensionsv1.AddToScheme(crdScheme))
	ser := json.NewSerializerWithOptions(json.DefaultMetaFactory, crdScheme, crdScheme, json.SerializerOptions{
		Yaml:   true,
		Pretty: false,
		Strict: false,
	})
	versions := schema.GroupVersions([]schema.GroupVersion{apiextensionsv1.SchemeGroupVersion})
	crdCodec = serializer.NewCodecFactory(crdScheme).CodecForVersions(ser, ser, versions, versions)
}

// DecodeCRD decodes a CRD from a YAML string.
func DecodeCRD(crdYAML string) (*apiextensionsv1.CustomResourceDefinition, error) {
	obj, err := runtime.Decode(crdCodec, []byte(crdYAML))
	if err != nil {
		return nil, fmt.Errorf("failed to decode CRD: %w", err)
	}
	crd, ok := obj.(*apiextensionsv1.CustomResourceDefinition)
	if !ok {
		return nil, fmt.Errorf("expected *apiextensionsv1.CustomResourceDefinition, got %T", obj)
	}
	return crd, nil
}
