<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# Garden deletion

```mermaid
---
title: Garden deletion (delete)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    DestroyingPlutono["Destroying Plutono"]
    DestroyingGardenerMetricsExporter["Destroying Gardener Metrics Exporter"]
    DestroyingKubeStateMetrics["Destroying Kube State Metrics"]
    DestroyingAlertmanager["Destroying Alertmanager"]
    DestroyingLongTermPrometheus["Destroying long-term Prometheus"]
    DestroyingGardenPrometheus["Destroying Garden Prometheus"]
    DestroyingIstioBasicAuthServer["Destroying istio-basic-auth-server"]
    DestroyingOpenTelemetryCollector["Destroying OpenTelemetry Collector"]
    DestroyingBlackboxExporter["Destroying blackbox-exporter"]
    DestroyingVPAForGardenerOperator["Destroying VPA for gardener-operator"]
    DestroyingGardenerDiscoveryServer["Destroying Gardener Discovery Server"]
    DestroyingGardenerDashboardWebTerminalControllerManager["Destroying Gardener Dashboard web terminal controller manager"]
    DestroyingGardenerDashboard["Destroying Gardener Dashboard"]
    DestroyingGardenerScheduler["Destroying Gardener Scheduler"]
    DestroyingExtensionResourcesBeforeKubeApiserver["Destroying extension resources before kube-apiserver"]
    WaitingUntilExtensionResourcesThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted["Waiting until extension resources that should be handled before kube-apiserver have been deleted"]
    DestroyingGardenerControllerManager["Destroying Gardener Controller Manager"]
    DestroyingGardenerAdmissionController["Destroying Gardener Admission Controller"]
    DestroyingGardenerAPIServer["Destroying Gardener API Server"]
    DestroyingVirtualSystemResources["Destroying virtual system resources"]
    EnsuringOnlyVirtualGardenManagedResourcesWhichAreRequiredForCleanupExist["Ensuring only virtual garden ManagedResources which are required for cleanup exist"]
    DestroyingGardenerVirtualGardenAccessResources["Destroying Gardener virtual garden access resources"]
    DestroyingKubernetesControllerManagerServer["Destroying Kubernetes Controller Manager Server"]
    EnsuringNoVirtualGardenManagedResourcesExistAnymore["Ensuring no virtual garden ManagedResources exist anymore"]
    DestroyingGardenerResourceManagerForVirtualGarden["Destroying gardener-resource-manager for virtual garden"]
    DestroyingKubernetesAPIServerServiceSNI["Destroying Kubernetes API server service SNI"]
    DestroyingKubernetesAPIServerService["Destroying Kubernetes API Server service"]
    DestroyingKubernetesAPIServer["Destroying Kubernetes API Server"]
    DestroyingMainAndEventsETCDsOfVirtualGarden["Destroying main and events ETCDs of virtual garden"]
    CleaningUpGenericTokenKubeconfig["Cleaning up generic token kubeconfig"]
    CleaningUpIstioInternalLoadBalancingConfigMap["Cleaning up Istio internal load balancing ConfigMap"]
    InvalidateClientForVirtualGarden["Invalidate client for virtual garden"]
    DestroyingExtensionResources["Destroying extension resources"]
    WaitingUntilExtensionResourcesHaveBeenDeleted["Waiting until extension resources have been deleted"]
    DestroyingDNSRecordsForVirtualGardenClusterAndIngressController{{"Destroying DNSRecords for virtual garden cluster and ingress controller\n[CONDITIONAL]"}}:::conditional
    DestroyingMainETCDBackupBucket{{"Destroying main ETCD backup bucket\n[CONDITIONAL]"}}:::conditional
    DestroyingETCDDruid{{"Destroying ETCD Druid\n[CONDITIONAL]"}}:::conditional
    DestroyingIstio["Destroying Istio"]
    DestroyingKubernetesVerticalPodAutoscaler["Destroying Kubernetes vertical pod autoscaler"]
    DestroyingNginxIngressController["Destroying nginx-ingress controller"]
    DestroyingPrometheusOperator["Destroying prometheus-operator"]
    DestroyingOpenTelemetryOperator["Destroying OpenTelemetry Operator"]
    DestroyingFluentOperatorCustomResources["Destroying fluent-operator custom resources"]
    DestroyingFluentBit["Destroying fluent-bit"]
    DestroyingFluentOperator["Destroying fluent-operator"]
    DestroyingVali["Destroying Vali"]
    DestroyingVictoriaLogs["Destroying VictoriaLogs"]
    DestroyingPersesOperator["Destroying perses-operator"]
    DestroyingVictoriaOperator["Destroying victoria-operator"]
    ResettingRequiredVirtualConditionOnExtensionsSinceVirtualClusterHasBeenDestroyed["Resetting RequiredVirtual condition on extensions since virtual cluster has been destroyed"]
    DestroyingRuntimeSystemResources["Destroying runtime system resources"]
    EnsuringNoManagedResourcesExistAnymore["Ensuring no ManagedResources exist anymore"]
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted{{"Destroying and waiting for gardener-resource-manager to be deleted\n[CONDITIONAL]"}}:::conditional
    DestroyingCustomResourceDefinitionForExtensions{{"Destroying custom resource definition for extensions\n[CONDITIONAL]"}}:::conditional
    DestroyingCustomResourceDefinitionForPrometheusOperator{{"Destroying custom resource definition for prometheus-operator\n[CONDITIONAL]"}}:::conditional
    DestroyingCustomResourceDefinitionForOpentelemetryOperator{{"Destroying custom resource definition for opentelemetry-operator\n[CONDITIONAL]"}}:::conditional
    DestroyingCustomResourceDefinitionForFluentOperator{{"Destroying custom resource definition for fluent-operator\n[CONDITIONAL]"}}:::conditional
    DestroyingCustomResourceDefinitionForIstio["Destroying custom resource definition for Istio"]
    DestroyingCustomResourceDefinitionForVPA{{"Destroying custom resource definition for VPA\n[CONDITIONAL]"}}:::conditional
    DestroyingETCDRelatedCustomResourceDefinitions{{"Destroying ETCD-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingCustomResourceDefinitionForPersesOperator{{"Destroying custom resource definition for perses-operator\n[CONDITIONAL]"}}:::conditional
    DestroyingCustomResourceDefinitionForVictoriaOperator{{"Destroying custom resource definition for victoria-operator\n[CONDITIONAL]"}}:::conditional
    CleaningUpSecrets["Cleaning up secrets"]
    CleaningUpGarbageCollectableConfigMapsAndSecrets["Cleaning up garbage-collectable ConfigMaps and Secrets"]
    SyncPointVirtualGardenManagedResourcesDestroyed(["Sync: Virtual Garden Managed Resources Destroyed"]):::syncpoint
    SyncPointCleanupRelevantVirtualManagedResourcesDestroyed(["Sync: Cleanup Relevant Virtual Managed Resources Destroyed"]):::syncpoint
    SyncPointVirtualGardenControlPlaneDestroyed(["Sync: Virtual Garden Control Plane Destroyed"]):::syncpoint
    SyncPointCleanedUp(["Sync: Cleaned Up"]):::syncpoint

    DestroyingGardenerDiscoveryServer --> SyncPointVirtualGardenManagedResourcesDestroyed
    DestroyingGardenerDashboardWebTerminalControllerManager --> SyncPointVirtualGardenManagedResourcesDestroyed
    DestroyingGardenerDashboard --> SyncPointVirtualGardenManagedResourcesDestroyed
    DestroyingGardenerScheduler --> SyncPointVirtualGardenManagedResourcesDestroyed
    DestroyingGardenerControllerManager --> SyncPointVirtualGardenManagedResourcesDestroyed
    DestroyingGardenerAdmissionController --> SyncPointVirtualGardenManagedResourcesDestroyed
    DestroyingGardenerAPIServer --> SyncPointVirtualGardenManagedResourcesDestroyed
    DestroyingVirtualSystemResources --> SyncPointVirtualGardenManagedResourcesDestroyed
    DestroyingGardenerVirtualGardenAccessResources --> SyncPointCleanupRelevantVirtualManagedResourcesDestroyed
    DestroyingKubernetesControllerManagerServer --> SyncPointCleanupRelevantVirtualManagedResourcesDestroyed
    CleaningUpGenericTokenKubeconfig --> SyncPointVirtualGardenControlPlaneDestroyed
    CleaningUpIstioInternalLoadBalancingConfigMap --> SyncPointVirtualGardenControlPlaneDestroyed
    DestroyingGardenerResourceManagerForVirtualGarden --> SyncPointVirtualGardenControlPlaneDestroyed
    DestroyingKubernetesAPIServerServiceSNI --> SyncPointVirtualGardenControlPlaneDestroyed
    DestroyingKubernetesAPIServerService --> SyncPointVirtualGardenControlPlaneDestroyed
    DestroyingKubernetesAPIServer --> SyncPointVirtualGardenControlPlaneDestroyed
    DestroyingMainAndEventsETCDsOfVirtualGarden --> SyncPointVirtualGardenControlPlaneDestroyed
    InvalidateClientForVirtualGarden --> SyncPointVirtualGardenControlPlaneDestroyed
    WaitingUntilExtensionResourcesHaveBeenDeleted --> SyncPointCleanedUp
    ResettingRequiredVirtualConditionOnExtensionsSinceVirtualClusterHasBeenDestroyed --> SyncPointCleanedUp
    DestroyingDNSRecordsForVirtualGardenClusterAndIngressController --> SyncPointCleanedUp
    DestroyingMainETCDBackupBucket --> SyncPointCleanedUp
    DestroyingETCDDruid --> SyncPointCleanedUp
    DestroyingIstio --> SyncPointCleanedUp
    DestroyingKubernetesVerticalPodAutoscaler --> SyncPointCleanedUp
    DestroyingNginxIngressController --> SyncPointCleanedUp
    DestroyingFluentOperatorCustomResources --> SyncPointCleanedUp
    DestroyingFluentBit --> SyncPointCleanedUp
    DestroyingFluentOperator --> SyncPointCleanedUp
    DestroyingVali --> SyncPointCleanedUp
    DestroyingPrometheusOperator --> SyncPointCleanedUp
    DestroyingOpenTelemetryOperator --> SyncPointCleanedUp
    DestroyingBlackboxExporter --> SyncPointCleanedUp
    DestroyingVPAForGardenerOperator --> SyncPointCleanedUp
    DestroyingPersesOperator --> SyncPointCleanedUp
    DestroyingVictoriaOperator --> SyncPointCleanedUp
    DestroyingVictoriaLogs --> SyncPointCleanedUp
    DestroyingAlertmanager --> DestroyingIstioBasicAuthServer
    DestroyingLongTermPrometheus --> DestroyingIstioBasicAuthServer
    DestroyingGardenPrometheus --> DestroyingIstioBasicAuthServer
    DestroyingPlutono --> DestroyingIstioBasicAuthServer
    DestroyingExtensionResourcesBeforeKubeApiserver --> WaitingUntilExtensionResourcesThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted
    WaitingUntilExtensionResourcesThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted --> DestroyingGardenerControllerManager
    WaitingUntilExtensionResourcesThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted --> DestroyingGardenerAdmissionController
    WaitingUntilExtensionResourcesThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted --> DestroyingGardenerAPIServer
    DestroyingGardenerAPIServer --> DestroyingVirtualSystemResources
    SyncPointVirtualGardenManagedResourcesDestroyed --> EnsuringOnlyVirtualGardenManagedResourcesWhichAreRequiredForCleanupExist
    EnsuringOnlyVirtualGardenManagedResourcesWhichAreRequiredForCleanupExist --> DestroyingGardenerVirtualGardenAccessResources
    EnsuringOnlyVirtualGardenManagedResourcesWhichAreRequiredForCleanupExist --> DestroyingKubernetesControllerManagerServer
    SyncPointCleanupRelevantVirtualManagedResourcesDestroyed --> EnsuringNoVirtualGardenManagedResourcesExistAnymore
    SyncPointVirtualGardenManagedResourcesDestroyed --> DestroyingGardenerResourceManagerForVirtualGarden
    EnsuringNoVirtualGardenManagedResourcesExistAnymore --> DestroyingGardenerResourceManagerForVirtualGarden
    DestroyingGardenerResourceManagerForVirtualGarden --> DestroyingKubernetesAPIServerServiceSNI
    DestroyingKubernetesAPIServerServiceSNI --> DestroyingKubernetesAPIServerService
    DestroyingGardenerResourceManagerForVirtualGarden --> DestroyingKubernetesAPIServer
    DestroyingKubernetesAPIServer --> DestroyingMainAndEventsETCDsOfVirtualGarden
    DestroyingKubernetesAPIServer --> CleaningUpGenericTokenKubeconfig
    DestroyingGardenerResourceManagerForVirtualGarden --> CleaningUpGenericTokenKubeconfig
    DestroyingKubernetesAPIServer --> CleaningUpIstioInternalLoadBalancingConfigMap
    DestroyingGardenerResourceManagerForVirtualGarden --> CleaningUpIstioInternalLoadBalancingConfigMap
    DestroyingKubernetesAPIServer --> InvalidateClientForVirtualGarden
    DestroyingGardenerResourceManagerForVirtualGarden --> InvalidateClientForVirtualGarden
    DestroyingExtensionResources --> WaitingUntilExtensionResourcesHaveBeenDeleted
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingDNSRecordsForVirtualGardenClusterAndIngressController
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingMainETCDBackupBucket
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingETCDDruid
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingIstio
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingKubernetesVerticalPodAutoscaler
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingNginxIngressController
    DestroyingAlertmanager --> DestroyingPrometheusOperator
    DestroyingGardenPrometheus --> DestroyingPrometheusOperator
    DestroyingLongTermPrometheus --> DestroyingPrometheusOperator
    DestroyingOpenTelemetryCollector --> DestroyingOpenTelemetryOperator
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingOpenTelemetryOperator
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingFluentOperatorCustomResources
    SyncPointVirtualGardenControlPlaneDestroyed --> DestroyingFluentBit
    DestroyingFluentOperatorCustomResources --> DestroyingFluentOperator
    DestroyingFluentBit --> DestroyingFluentOperator
    DestroyingFluentOperatorCustomResources --> DestroyingVali
    DestroyingFluentOperatorCustomResources --> DestroyingVictoriaLogs
    DestroyingVictoriaLogs --> DestroyingVictoriaOperator
    SyncPointVirtualGardenControlPlaneDestroyed --> ResettingRequiredVirtualConditionOnExtensionsSinceVirtualClusterHasBeenDestroyed
    SyncPointCleanedUp --> DestroyingRuntimeSystemResources
    DestroyingRuntimeSystemResources --> EnsuringNoManagedResourcesExistAnymore
    EnsuringNoManagedResourcesExistAnymore --> DestroyingAndWaitingForGardenerResourceManagerToBeDeleted
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingCustomResourceDefinitionForExtensions
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingCustomResourceDefinitionForPrometheusOperator
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingCustomResourceDefinitionForOpentelemetryOperator
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingCustomResourceDefinitionForFluentOperator
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingCustomResourceDefinitionForIstio
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingCustomResourceDefinitionForVPA
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingETCDRelatedCustomResourceDefinitions
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingCustomResourceDefinitionForPersesOperator
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> DestroyingCustomResourceDefinitionForVictoriaOperator
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> CleaningUpSecrets
    DestroyingAndWaitingForGardenerResourceManagerToBeDeleted --> CleaningUpGarbageCollectableConfigMapsAndSecrets
```
