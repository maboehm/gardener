<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# Seed reconciliation

```mermaid
---
title: Seed reconciliation (runReconcileSeedFlow)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    DeployingMachineRelatedCustomResourceDefinitions["Deploying machine-related custom resource definitions"]
    DeployingExtensionsRelatedCustomResourceDefinitions["Deploying extensions-related custom resource definitions"]
    DeployingETCDRelatedCustomResourceDefinitions{{"Deploying ETCD-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DeployingIstioRelatedCustomResourceDefinitions{{"Deploying Istio-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DeployingVPARelatedCustomResourceDefinitions{{"Deploying VPA-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DeployingLoggingRelatedCustomResourceDefinitions{{"Deploying logging-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DeployingPrometheusRelatedCustomResourceDefinitions{{"Deploying Prometheus-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DeployingPersesRelatedCustomResourceDefinitions{{"Deploying Perses-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DeployingVictoriaRelatedCustomResourceDefinitions{{"Deploying Victoria-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DeployOpenTelemetryRelatedCustomResourceDefinitions{{"Deploy OpenTelemetry-related custom resource definitions\n[CONDITIONAL]"}}:::conditional
    DeployingVPAForGardenlet["Deploying VPA for gardenlet"]
    WaitingForCustomResourceDefinitionsForIstio{{"Waiting for custom resource definitions for Istio\n[CONDITIONAL]"}}:::conditional
    DeployingAndWaitingForGardenerResourceManagerToBeHealthy{{"Deploying and waiting for gardener-resource-manager to be healthy\n[CONDITIONAL]"}}:::conditional
    DeployingSystemResources["Deploying system resources"]
    DeployingReferencedResources["Deploying referenced resources"]
    WaitingUntilRequiredExtensionsAreReady["Waiting until required extensions are ready"]
    DeployingClusterIdentity{{"Deploying cluster-identity\n[CONDITIONAL]"}}:::conditional
    DeployingExtensionResources["Deploying extension resources"]
    WaitingUntilExtensionResourcesAreReady["Waiting until extension resources are ready"]
    CleaningUpOrphanExposureClassHandlerResources["Cleaning up orphan ExposureClass handler resources"]
    DeployingIstio["Deploying Istio"]
    WaitingUntilIstioLoadBalancerIsReadyAndManagedIngressDNSRecordIsReconciled["Waiting until istio LoadBalancer is ready and managed ingress DNS record is reconciled"]
    DeployingClusterAutoscalerResources["Deploying cluster-autoscaler resources"]
    DeployingDependencyWatchdogWeeder["Deploying dependency-watchdog-weeder"]
    DeployingDependencyWatchdogProber["Deploying dependency-watchdog-prober"]
    RenewingGardenAccessSecrets{{"Renewing garden access secrets\n[CONDITIONAL]"}}:::conditional
    RenewingWorkloadIdentityTokens{{"Renewing workload identity tokens\n[CONDITIONAL]"}}:::conditional
    RenewingGardenKubeconfig{{"Renewing garden kubeconfig\n[CONDITIONAL]"}}:::conditional
    ReconcilingKubeApiserverService["Reconciling kube-apiserver service"]
    ReconcilingKubeApiserverIngress["Reconciling kube-apiserver ingress"]
    ReconcilingKubernetesVerticalPodAutoscaler{{"Reconciling Kubernetes vertical pod autoscaler\n[CONDITIONAL]"}}:::conditional
    DeployingETCDDruid{{"Deploying ETCD Druid\n[CONDITIONAL]"}}:::conditional
    DeployingKubeStateMetrics["Deploying kube-state-metrics"]
    DeployingOpenTelemetryOperator{{"Deploying OpenTelemetry Operator\n[CONDITIONAL]"}}:::conditional
    DeployingOpenTelemetryCollector{{"Deploying OpenTelemetry Collector\n[CONDITIONAL]"}}:::conditional
    DeployingFluentOperator{{"Deploying Fluent Operator\n[CONDITIONAL]"}}:::conditional
    DeployingFluentBit{{"Deploying Fluent Bit\n[CONDITIONAL]"}}:::conditional
    DeployingFluentOperatorCustomResources{{"Deploying Fluent Operator custom resources\n[CONDITIONAL]"}}:::conditional
    DeployingPlutono["Deploying Plutono"]
    WaitingUntilPlutonoIsReady["Waiting until Plutono is ready"]
    DeployingVictoriaLogs{{"Deploying VictoriaLogs\n[CONDITIONAL]"}}:::conditional
    DeployingVali{{"Deploying Vali\n[CONDITIONAL]"}}:::conditional
    DeployingPrometheusOperator{{"Deploying Prometheus Operator\n[CONDITIONAL]"}}:::conditional
    DeployingCachePrometheus["Deploying cache Prometheus"]
    DeployingSeedPrometheus["Deploying seed Prometheus"]
    DeployingAggregatePrometheus["Deploying aggregate Prometheus"]
    WaitingUntilAggregatePrometheusIsReady["Waiting until aggregate Prometheus is ready"]
    DeployingAlertmanager["Deploying Alertmanager"]
    DeployingIstioBasicAuthServer["Deploying istio-basic-auth-server"]
    DeployingPersesOperator{{"Deploying Perses Operator\n[CONDITIONAL]"}}:::conditional
    DeployingVictoriaOperator{{"Deploying Victoria Operator\n[CONDITIONAL]"}}:::conditional
    DeletingStaleExtensionResources["Deleting stale extension resources"]
    WaitingUntilStaleExtensionResourcesAreDeleted["Waiting until stale extension resources are deleted"]
    DeployingBackupBucketForSeed{{"Deploying BackupBucket for seed\n[CONDITIONAL]"}}:::conditional
    SyncPointCRDs(["Sync: CRDs"]):::syncpoint
    SyncPointReadyForSystemComponents(["Sync: Ready For System Components"]):::syncpoint

    DeployingMachineRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingExtensionsRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingETCDRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingIstioRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingVPARelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingLoggingRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingPrometheusRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingPersesRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingVictoriaRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployOpenTelemetryRelatedCustomResourceDefinitions --> SyncPointCRDs
    DeployingAndWaitingForGardenerResourceManagerToBeHealthy --> SyncPointReadyForSystemComponents
    DeployingClusterIdentity --> SyncPointReadyForSystemComponents
    CleaningUpOrphanExposureClassHandlerResources --> SyncPointReadyForSystemComponents
    WaitingUntilExtensionResourcesAreReady --> SyncPointReadyForSystemComponents
    SyncPointCRDs --> DeployingVPAForGardenlet
    DeployingIstioRelatedCustomResourceDefinitions --> WaitingForCustomResourceDefinitionsForIstio
    SyncPointCRDs --> DeployingAndWaitingForGardenerResourceManagerToBeHealthy
    WaitingForCustomResourceDefinitionsForIstio --> DeployingAndWaitingForGardenerResourceManagerToBeHealthy
    DeployingAndWaitingForGardenerResourceManagerToBeHealthy --> DeployingSystemResources
    DeployingSystemResources --> DeployingReferencedResources
    DeployingReferencedResources --> WaitingUntilRequiredExtensionsAreReady
    WaitingUntilRequiredExtensionsAreReady --> DeployingClusterIdentity
    WaitingUntilRequiredExtensionsAreReady --> DeployingExtensionResources
    DeployingExtensionResources --> WaitingUntilExtensionResourcesAreReady
    WaitingUntilRequiredExtensionsAreReady --> CleaningUpOrphanExposureClassHandlerResources
    SyncPointReadyForSystemComponents --> DeployingIstio
    DeployingIstio --> WaitingUntilIstioLoadBalancerIsReadyAndManagedIngressDNSRecordIsReconciled
    SyncPointReadyForSystemComponents --> DeployingClusterAutoscalerResources
    SyncPointReadyForSystemComponents --> DeployingDependencyWatchdogWeeder
    SyncPointReadyForSystemComponents --> DeployingDependencyWatchdogProber
    SyncPointReadyForSystemComponents --> RenewingGardenAccessSecrets
    SyncPointReadyForSystemComponents --> RenewingWorkloadIdentityTokens
    SyncPointReadyForSystemComponents --> RenewingGardenKubeconfig
    SyncPointReadyForSystemComponents --> ReconcilingKubeApiserverService
    SyncPointReadyForSystemComponents --> ReconcilingKubeApiserverIngress
    SyncPointReadyForSystemComponents --> ReconcilingKubernetesVerticalPodAutoscaler
    SyncPointReadyForSystemComponents --> DeployingETCDDruid
    SyncPointReadyForSystemComponents --> DeployingKubeStateMetrics
    SyncPointReadyForSystemComponents --> DeployingOpenTelemetryOperator
    SyncPointReadyForSystemComponents --> DeployingOpenTelemetryCollector
    SyncPointReadyForSystemComponents --> DeployingFluentOperator
    SyncPointReadyForSystemComponents --> DeployingFluentBit
    DeployingFluentOperator --> DeployingFluentBit
    SyncPointReadyForSystemComponents --> DeployingFluentOperatorCustomResources
    DeployingFluentOperator --> DeployingFluentOperatorCustomResources
    SyncPointReadyForSystemComponents --> DeployingPlutono
    DeployingPlutono --> WaitingUntilPlutonoIsReady
    SyncPointReadyForSystemComponents --> DeployingVictoriaLogs
    SyncPointReadyForSystemComponents --> DeployingVali
    SyncPointReadyForSystemComponents --> DeployingPrometheusOperator
    SyncPointReadyForSystemComponents --> DeployingCachePrometheus
    SyncPointReadyForSystemComponents --> DeployingSeedPrometheus
    SyncPointReadyForSystemComponents --> DeployingAggregatePrometheus
    DeployingAggregatePrometheus --> WaitingUntilAggregatePrometheusIsReady
    SyncPointReadyForSystemComponents --> DeployingAlertmanager
    SyncPointReadyForSystemComponents --> DeployingIstioBasicAuthServer
    WaitingUntilAggregatePrometheusIsReady --> DeployingIstioBasicAuthServer
    WaitingUntilPlutonoIsReady --> DeployingIstioBasicAuthServer
    SyncPointReadyForSystemComponents --> DeployingPersesOperator
    SyncPointReadyForSystemComponents --> DeployingVictoriaOperator
    SyncPointReadyForSystemComponents --> DeletingStaleExtensionResources
    DeletingStaleExtensionResources --> WaitingUntilStaleExtensionResourcesAreDeleted
    SyncPointReadyForSystemComponents --> DeployingBackupBucketForSeed
```
