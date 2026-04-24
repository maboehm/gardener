<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# Seed deletion

```mermaid
---
title: Seed deletion (runDeleteSeedFlow)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    DestroyingClusterIdentity{{"Destroying cluster-identity\n[CONDITIONAL]"}}:::conditional
    DestroyingManagedIngressDNSRecordIfExisting["Destroying managed ingress DNS record (if existing)"]
    DestroyingCachePrometheus["Destroying cache Prometheus"]
    DestroyingSeedPrometheus["Destroying seed Prometheus"]
    DestroyingAggregatePrometheus["Destroying aggregate Prometheus"]
    DestroyingAlertManager["Destroying AlertManager"]
    DestroyingClusterAutoscalerResources["Destroying cluster-autoscaler resources"]
    DestroyingNginxIngress["Destroying nginx-ingress"]
    DestroyDependencyWatchdogWeeder["Destroy dependency-watchdog-weeder"]
    DestroyDependencyWatchdogProber["Destroy dependency-watchdog-prober"]
    DestroyKubeApiserverIngress["Destroy kube-apiserver ingress"]
    DestroyKubeApiserverService["Destroy kube-apiserver service"]
    DestroyIstio["Destroy Istio"]
    DestroyFluentOperatorCustomResources["Destroy Fluent Operator Custom Resources"]
    DestroyingPlutono["Destroying plutono"]
    DestroyingIstioBasicAuthServer["Destroying istio basic auth server"]
    DestroyingEtcdDruid{{"Destroying etcd druid\n[CONDITIONAL]"}}:::conditional
    DestroyKubernetesVerticalPodAutoscaler{{"Destroy Kubernetes vertical pod autoscaler\n[CONDITIONAL]"}}:::conditional
    DestroyKubeStateMetrics["Destroy kube-state-metrics"]
    DestroyPrometheusOperator{{"Destroy Prometheus Operator\n[CONDITIONAL]"}}:::conditional
    DestroyingOpenTelemetryCollector{{"Destroying OpenTelemetry Collector\n[CONDITIONAL]"}}:::conditional
    DestroyOpenTelemetryOperator{{"Destroy OpenTelemetry Operator\n[CONDITIONAL]"}}:::conditional
    DestroyFluentBit{{"Destroy Fluent Bit\n[CONDITIONAL]"}}:::conditional
    DestroyFluentOperator{{"Destroy Fluent Operator\n[CONDITIONAL]"}}:::conditional
    DestroyVali{{"Destroy Vali\n[CONDITIONAL]"}}:::conditional
    DestroyVictoriaLogs{{"Destroy VictoriaLogs\n[CONDITIONAL]"}}:::conditional
    DestroyingVPAForGardenlet["Destroying VPA for gardenlet"]
    DestroyPersesOperator{{"Destroy Perses Operator\n[CONDITIONAL]"}}:::conditional
    DestroyVictoriaOperator{{"Destroy Victoria Operator\n[CONDITIONAL]"}}:::conditional
    DeletingExtensionResources["Deleting extension resources"]
    WaitingUntilExtensionResourcesHaveBeenDeleted["Waiting until extension resources have been deleted"]
    EnsuringAllControllerInstallationsAreGone["Ensuring all ControllerInstallations are gone"]
    DeletingReferencedResources["Deleting referenced resources"]
    DestroyingMachineControllerManagerCustomResourceDefinitions["Destroying machine-controller-manager custom resource definitions"]
    DestroyingExtensionsRelatedCustomResourceDefinitions{{"Destroying extensions-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingETCDRelatedCustomResourceDefinitions{{"Destroying ETCD-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingIstioCustomResourceDefinitions{{"Destroying Istio custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingVPARelatedCustomResourceDefinitions{{"Destroying VPA-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingFluentOperatorCustomResourceDefinitions{{"Destroying Fluent Operator custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyOpenTelemetryCustomResourceDefinitions{{"Destroy OpenTelemetry custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingPrometheusRelatedCustomResourceDefinitions{{"Destroying Prometheus-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingPersesOperatorCustomResourceDefinitions{{"Destroying Perses Operator custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingVictoriaOperatorCustomResourceDefinitions{{"Destroying Victoria Operator custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DestroyingSystemResources["Destroying system resources"]
    EnsuringAllManagedResourcesAreGone{{"Ensuring all ManagedResources are gone\n[CONDITIONAL]"}}:::conditional
    DestroyingGardenerResourceManager{{"Destroying gardener-resource-manager\n[CONDITIONAL]"}}:::conditional
    SyncPointCleanedUp(["Sync: Cleaned Up"]):::syncpoint
    DestroyCRDs(["Sync: destroy CRDs"]):::syncpoint

    DestroyingManagedIngressDNSRecordIfExisting --> SyncPointCleanedUp
    DestroyingClusterIdentity --> SyncPointCleanedUp
    DestroyingCachePrometheus --> SyncPointCleanedUp
    DestroyingSeedPrometheus --> SyncPointCleanedUp
    DestroyingOpenTelemetryCollector --> SyncPointCleanedUp
    DestroyingAggregatePrometheus --> SyncPointCleanedUp
    DestroyingAlertManager --> SyncPointCleanedUp
    DestroyingNginxIngress --> SyncPointCleanedUp
    DestroyingClusterAutoscalerResources --> SyncPointCleanedUp
    DestroyDependencyWatchdogWeeder --> SyncPointCleanedUp
    DestroyDependencyWatchdogProber --> SyncPointCleanedUp
    DestroyKubeApiserverIngress --> SyncPointCleanedUp
    DestroyKubeApiserverService --> SyncPointCleanedUp
    DestroyIstio --> SyncPointCleanedUp
    DestroyFluentOperatorCustomResources --> SyncPointCleanedUp
    DestroyPrometheusOperator --> SyncPointCleanedUp
    DestroyOpenTelemetryOperator --> SyncPointCleanedUp
    DestroyingPlutono --> SyncPointCleanedUp
    DestroyingIstioBasicAuthServer --> SyncPointCleanedUp
    DestroyKubeStateMetrics --> SyncPointCleanedUp
    DestroyingEtcdDruid --> SyncPointCleanedUp
    DestroyKubernetesVerticalPodAutoscaler --> SyncPointCleanedUp
    DestroyFluentBit --> SyncPointCleanedUp
    DestroyFluentOperator --> SyncPointCleanedUp
    DestroyVali --> SyncPointCleanedUp
    DestroyVictoriaLogs --> SyncPointCleanedUp
    DestroyingVPAForGardenlet --> SyncPointCleanedUp
    DestroyPersesOperator --> SyncPointCleanedUp
    DestroyVictoriaOperator --> SyncPointCleanedUp
    WaitingUntilExtensionResourcesHaveBeenDeleted --> SyncPointCleanedUp
    DestroyingMachineControllerManagerCustomResourceDefinitions --> DestroyCRDs
    DestroyingExtensionsRelatedCustomResourceDefinitions --> DestroyCRDs
    DestroyingIstioCustomResourceDefinitions --> DestroyCRDs
    DestroyingVPARelatedCustomResourceDefinitions --> DestroyCRDs
    DestroyingETCDRelatedCustomResourceDefinitions --> DestroyCRDs
    DestroyingFluentOperatorCustomResourceDefinitions --> DestroyCRDs
    DestroyingPrometheusRelatedCustomResourceDefinitions --> DestroyCRDs
    DestroyingPersesOperatorCustomResourceDefinitions --> DestroyCRDs
    DestroyingVictoriaOperatorCustomResourceDefinitions --> DestroyCRDs
    DestroyOpenTelemetryCustomResourceDefinitions --> DestroyCRDs
    DestroyingAggregatePrometheus --> DestroyingIstioBasicAuthServer
    DestroyingPlutono --> DestroyingIstioBasicAuthServer
    DestroyingOpenTelemetryCollector --> DestroyOpenTelemetryOperator
    DestroyFluentOperatorCustomResources --> DestroyFluentOperator
    DestroyFluentBit --> DestroyFluentOperator
    DestroyFluentOperatorCustomResources --> DestroyVali
    DestroyFluentOperatorCustomResources --> DestroyVictoriaLogs
    DestroyVictoriaLogs --> DestroyVictoriaOperator
    DeletingExtensionResources --> WaitingUntilExtensionResourcesHaveBeenDeleted
    SyncPointCleanedUp --> EnsuringAllControllerInstallationsAreGone
    EnsuringAllControllerInstallationsAreGone --> DeletingReferencedResources
    EnsuringAllControllerInstallationsAreGone --> DestroyingMachineControllerManagerCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyingExtensionsRelatedCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyingETCDRelatedCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyingIstioCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyingVPARelatedCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyingFluentOperatorCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyOpenTelemetryCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyingPrometheusRelatedCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyingPersesOperatorCustomResourceDefinitions
    EnsuringAllControllerInstallationsAreGone --> DestroyingVictoriaOperatorCustomResourceDefinitions
    DestroyCRDs --> DestroyingSystemResources
    DestroyingSystemResources --> EnsuringAllManagedResourcesAreGone
    EnsuringAllManagedResourcesAreGone --> DestroyingGardenerResourceManager
```
