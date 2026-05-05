#!/usr/bin/env bash
# SPDX-FileCopyrightText: SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0

set -o errexit
set -o pipefail

COMMAND="${1:-up}"
VALID_COMMANDS=("up" "down")

SCENARIO="${SCENARIO:-default}"
declare -A SCENARIO_LEVEL=(
  [setup]=1     # Only starts the kind cluster and renders the manifests.
  [bootstrap]=2 # Like 'setup', but also runs `gardenadm bootstrap` and exports the kubeconfig for the self-hosted shoot
  [join]=3      # Like 'default', but also runs `gardenadm join` on gind-machine-1 to join it as worker node
  [connect]=4   # Like 'join', but also deploys Gardener into the self-hosted shoot and runs `gardenadm connect` to deploy gardenlet which registers the Shoot
  [full]=5      # Like 'connect', but also registers the self-hosted shoot as a seed via a ManagedSeed
)

if [[ -z "${SCENARIO_LEVEL[$SCENARIO]+x}" ]]; then
  echo "Error: Invalid scenario '${SCENARIO}'. Valid options are: ${!SCENARIO_LEVEL[*]}." >&2
  exit 1
fi

level="${SCENARIO_LEVEL[$SCENARIO]}"

generated_dir="$(dirname "$0")/../dev-setup/gardenadm/resources/generated"

up() {
  if ! kind get clusters | grep gardener-local &>/dev/null; then
    make kind-up
  fi

  controlplane_pod=""

  make gardenadm-up SCENARIO=managed-infra

  # Run `gardenadm bootstrap` and export the kubeconfig for the self-hosted shoot
  if ((level >= 2)); then
    make gardenadm # builds gardenadm binary locally
    export IMAGEVECTOR_OVERWRITE="$generated_dir/.imagevector-overwrite.yaml"
    export IMAGEVECTOR_OVERWRITE_CHARTS="$generated_dir/.imagevector-overwrite-charts.yaml"
    KUBECONFIG="$KUBECONFIG_RUNTIME_CLUSTER" "$(dirname "$0")/../bin/gardenadm" bootstrap -d "$generated_dir/managed-infra"

    controlplane_pod="$(kubectl --kubeconfig="$KUBECONFIG_RUNTIME_CLUSTER" get po -n shoot--garden--root -l app=machine -o yaml | yq '.items[]|select(.metadata.name|contains("control")).metadata.labels.machine')"
    tmp_exposure "$controlplane_pod"
    KUBECONFIG="$KUBECONFIG_RUNTIME_CLUSTER" ./hack/usage/generate-kubeconfig.sh self-hosted-shoot >"$KUBECONFIG_SELFHOSTEDSHOOT_CLUSTER"
  fi

  # Deploy Gardener into the self-hosted shoot and run `gardenadm connect` to deploy gardenlet which registers the Shoot
  if ((level >= 4)); then
    make gardenadm-up SCENARIO=connect-managed-infra # deploys gardener-operator, the 'Garden' resource, and waits for reconciliation
    connect_command="$(KUBECONFIG=$KUBECONFIG_VIRTUAL_GARDEN_CLUSTER "$(dirname "$0")/../bin/gardenadm" token create --print-connect-command --shoot-namespace garden --shoot-name root)"
    kubectl --kubeconfig="$KUBECONFIG_RUNTIME_CLUSTER" exec -it -n shoot--garden--root "$controlplane_pod" -- sh -c \
      "export IMAGEVECTOR_OVERWRITE=/var/lib/gardenadm/.imagevector-overwrite.yaml; export IMAGEVECTOR_OVERWRITE_CHARTS=/var/lib/gardenadm/.imagevector-overwrite-charts.yaml; $connect_command"
  fi

  # Register the self-hosted shoot as a seed via a ManagedSeed
  if ((level >= 5)); then
    make seed-up KUBECONFIG="$KUBECONFIG_SELFHOSTEDSHOOT_CLUSTER"
  fi
}

down() {
  make kind-down
}

# TODO(maboehm): Remove after SelfHostedShoot Exposure is working
tmp_exposure() {
  local machine="$1"

  export KUBECONFIG="$KUBECONFIG_RUNTIME_CLUSTER"
  cat <<EOF | kubectl apply -o yaml -f -
apiVersion: v1
kind: Service
metadata:
  name: control-plane
  namespace: shoot--garden--root
spec:
  selector:
    machine: $machine
  ports:
    - protocol: TCP
      port: 443
      targetPort: 443
  type: LoadBalancer
EOF
  kubectl wait -n shoot--garden--root service/control-plane --for=jsonpath='{.status.loadBalancer.ingress}'
  ip="$(kubectl get svc -n shoot--garden--root control-plane -o yaml | yq '.status.loadBalancer.ingress[0].ip')"
  patch='{"spec":{"values":["'$ip'"]},"metadata":{"annotations":{"gardener.cloud/operation":"reconcile"}}}'
  kubectl exec -it -n shoot--garden--root "$(kubectl get pod -n shoot--garden--root -l machine="$machine" -o name | cut -d/ -f2)" -- \
    sh -c "KUBECONFIG=/etc/kubernetes/admin.conf kubectl patch dnsrecord -n kube-system root-external -o yaml --type=merge -p '$patch'"
}

case "$COMMAND" in
up)
  up
  ;;

down)
  down
  ;;
*)
  echo "Error: Invalid command '${COMMAND}'. Valid options are: ${VALID_COMMANDS[*]}." >&2
  exit 1
  ;;
esac
