<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# Garden reconciliation

```mermaid
---
title: Garden reconciliation (reconcile)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    GeneratingGenericTokenKubeconfig["Generating generic token kubeconfig"]
    ReconcileIstioInternalLoadBalancingConfigMap["Reconcile Istio internal load balancing ConfigMap"]
    GeneratingObservabilityIngressPassword["Generating observability ingress password"]
    DeployingVPAForGardenerOperator["Deploying VPA for gardener-operator"]
    DeployingServiceMonitorForGardenerOperator["Deploying ServiceMonitor for gardener-operator"]
    DeployingNginxIngressController["Deploying nginx-ingress controller"]
    ReconcilingKubernetesVerticalPodAutoscaler["Reconciling Kubernetes vertical pod autoscaler"]
    DeployingETCDDruid["Deploying ETCD Druid"]
    DeployingIstio["Deploying Istio"]
    WaitingForETCDDruidToBeReady["Waiting for ETCD Druid to be ready"]
    ReconcilingDNSRecordsForVirtualGardenClusterAndIngressController{{"Reconciling DNSRecords for virtual garden cluster and ingress controller\n[CONDITIONAL]"}}:::conditional
    ReconcilingMainETCDBackupBucket{{"Reconciling main ETCD backup bucket\n[CONDITIONAL]"}}:::conditional
    ReconcilingMainETCDBackupEntry{{"Reconciling main ETCD backup entry\n[CONDITIONAL]"}}:::conditional
    DeployingMainAndEventsETCDsOfVirtualGarden["Deploying main and events ETCDs of virtual garden"]
    WaitingUntilMainAndEventETCDsReportReadiness["Waiting until main and event ETCDs report readiness"]
    DeployingExtensionResourcesBeforeKubeApiserver["Deploying extension resources before kube-apiserver"]
    WaitingUntilExtensionResourcesHandledBeforeKubeApiserverAreReady["Waiting until extension resources handled before kube-apiserver are ready"]
    DeployingAndWaitingForKubeApiserverServiceInTheRuntimeCluster["Deploying and waiting for kube-apiserver service in the runtime cluster"]
    DeployingKubernetesAPIServer["Deploying Kubernetes API Server"]
    WaitingUntilKubernetesAPIServerRolledOut["Waiting until Kubernetes API server rolled out"]
    DeployingKubernetesAPIServerServiceSNI["Deploying Kubernetes API server service SNI"]
    DeployingKubernetesControllerManager["Deploying Kubernetes Controller Manager"]
    DeployingGardenerResourceManagerForVirtualGarden["Deploying gardener-resource-manager for virtual garden"]
    WaitingUntilGardenerResourceManagerForVirtualGardenRolledOut["Waiting until gardener-resource-manager for virtual garden rolled out"]
    DeployingGardenerAPIServer["Deploying Gardener API Server"]
    WaitingUntilGardenerAPIServerRolledOut["Waiting until Gardener API server rolled out"]
    DeployingGardenerAdmissionController["Deploying Gardener Admission Controller"]
    DeployingGardenerControllerManager["Deploying Gardener Controller Manager"]
    DeployingGardenerScheduler["Deploying Gardener Scheduler"]
    DeployingGardenerDiscoveryServer["Deploying Gardener Discovery Server"]
    DeployingVirtualSystemResources["Deploying virtual system resources"]
    DeployingResourcesForGardenerOperatorAccessToVirtualGarden["Deploying resources for gardener-operator access to virtual garden"]
    RenewingVirtualGardenAccessSecretsAfterCreationOfNewServiceAccountSigningKey{{"Renewing virtual garden access secrets after creation of new ServiceAccount signing key\n[CONDITIONAL]"}}:::conditional
    InitializingConnectionToVirtualGardenCluster["Initializing connection to virtual garden cluster"]
    DeployGardenerInfoConfigMap["Deploy gardener-info ConfigMap"]
    DeployingExtensionResources["Deploying extension resources"]
    WaitingUntilExtensionResourcesAreReady["Waiting until extension resources are ready"]
    DeletingStaleExtensionResources["Deleting stale extension resources"]
    WaitingUntilStaleExtensionResourcesAreDeleted["Waiting until stale extension resources are deleted"]
    ReconcilingGardenerDashboard["Reconciling Gardener Dashboard"]
    ReconcilingGardenerDashboardWebTerminalControllerManager["Reconciling Gardener Dashboard web terminal controller manager"]
    AnnotateSeedsToTriggerRenewalOfTheirGardenAccessSecrets{{"Annotate seeds to trigger renewal of their garden access secrets\n[CONDITIONAL]"}}:::conditional
    CheckIfAllSeedsFinishedTheRenewalOfTheirGardenAccessSecrets{{"Check if all seeds finished the renewal of their garden access secrets\n[CONDITIONAL]"}}:::conditional
    AnnotateSeedsToTriggerRenewalOfWorkloadIdentityTokens{{"Annotate seeds to trigger renewal of workload identity tokens\n[CONDITIONAL]"}}:::conditional
    CheckIfAllSeedsFinishedTheRenewalOfTheirWorkloadIdentityTokens{{"Check if all seeds finished the renewal of their workload identity tokens\n[CONDITIONAL]"}}:::conditional
    AnnotateSeedsToTriggerRenewalOfTheirGardenletKubeconfig{{"Annotate seeds to trigger renewal of their gardenlet kubeconfig\n[CONDITIONAL]"}}:::conditional
    CheckIfAllSeedsFinishedTheRenewalOfTheirGardenletKubeconfig{{"Check if all seeds finished the renewal of their gardenlet kubeconfig\n[CONDITIONAL]"}}:::conditional
    LabelingEncryptedResourcesAfterModificationOfEncryptionConfigOrToReEncryptThemWithNewETCDEncryptionKey{{"Labeling encrypted resources after modification of encryption config or to re-encrypt them with new ETCD encryption key\n[CONDITIONAL]"}}:::conditional
    SnapshottingETCDAfterModificationOfEncryptionConfigOrResourcesAreReEncryptedWithNewETCDEncryptionKey{{"Snapshotting ETCD after modification of encryption config or resources are re-encrypted with new ETCD encryption key\n[CONDITIONAL]"}}:::conditional
    RemovingLabelFromReEncryptedResourcesAfterModificationOfEncryptionConfigOrRotationOfETCDEncryptionKey{{"Removing label from re-encrypted resources after modification of encryption config or rotation of ETCD encryption key\n[CONDITIONAL]"}}:::conditional
    DeployingFluentOperator["Deploying fluent-operator"]
    DeployingFluentOperatorCustomResources["Deploying fluent-operator CustomResources"]
    DeployingFluentBit["Deploying fluent-bit"]
    DeployingVictoriaLogs{{"Deploying VictoriaLogs\n[CONDITIONAL]"}}:::conditional
    DeployingVali{{"Deploying Vali\n[CONDITIONAL]"}}:::conditional
    DeployingPrometheusOperator["Deploying prometheus-operator"]
    DeployingOpenTelemetryOperator["Deploying OpenTelemetry Operator"]
    DeployingAlertmanager["Deploying Alertmanager"]
    WaitingUntilAlertmanagerIsReady["Waiting until Alertmanager is ready"]
    DeployingGardenPrometheus["Deploying Garden Prometheus"]
    WaitingUntilGardenPrometheusIsReady["Waiting until Garden Prometheus is ready"]
    DeployingLongTermPrometheus["Deploying long-term Prometheus"]
    WaitingUntilLongTermPrometheusIsReady["Waiting until long-term Prometheus is ready"]
    DeployingOpenTelemetryCollector["Deploying OpenTelemetry Collector"]
    DeployingBlackboxExporter["Deploying blackbox-exporter"]
    DeployingKubeStateMetrics["Deploying Kube State Metrics"]
    DeployingGardenerMetricsExporter["Deploying Gardener Metrics Exporter"]
    DeployingPlutono["Deploying Plutono"]
    WaitingUntilPlutonoIsReady["Waiting until Plutono is ready"]
    DeployingIstioBasicAuthServer["Deploying istio-basic-auth-server"]
    DeployingPersesOperator["Deploying perses-operator"]
    DeployingVictoriaOperator["Deploying victoria-operator"]
    SyncPointSystemComponents(["Sync: System Components"]):::syncpoint

    GeneratingGenericTokenKubeconfig --> SyncPointSystemComponents
    GeneratingObservabilityIngressPassword --> SyncPointSystemComponents
    ReconcilingKubernetesVerticalPodAutoscaler --> SyncPointSystemComponents
    DeployingETCDDruid --> SyncPointSystemComponents
    DeployingIstio --> SyncPointSystemComponents
    DeployingNginxIngressController --> SyncPointSystemComponents
    ReconcileIstioInternalLoadBalancingConfigMap --> SyncPointSystemComponents
    DeployingETCDDruid --> WaitingForETCDDruidToBeReady
    SyncPointSystemComponents --> ReconcilingDNSRecordsForVirtualGardenClusterAndIngressController
    ReconcilingMainETCDBackupBucket --> ReconcilingMainETCDBackupEntry
    WaitingForETCDDruidToBeReady --> DeployingMainAndEventsETCDsOfVirtualGarden
    ReconcilingMainETCDBackupEntry --> DeployingMainAndEventsETCDsOfVirtualGarden
    DeployingMainAndEventsETCDsOfVirtualGarden --> WaitingUntilMainAndEventETCDsReportReadiness
    DeployingExtensionResourcesBeforeKubeApiserver --> WaitingUntilExtensionResourcesHandledBeforeKubeApiserverAreReady
    SyncPointSystemComponents --> DeployingAndWaitingForKubeApiserverServiceInTheRuntimeCluster
    WaitingUntilMainAndEventETCDsReportReadiness --> DeployingKubernetesAPIServer
    WaitingUntilExtensionResourcesHandledBeforeKubeApiserverAreReady --> DeployingKubernetesAPIServer
    DeployingKubernetesAPIServer --> WaitingUntilKubernetesAPIServerRolledOut
    WaitingUntilKubernetesAPIServerRolledOut --> DeployingKubernetesAPIServerServiceSNI
    WaitingUntilKubernetesAPIServerRolledOut --> DeployingKubernetesControllerManager
    WaitingUntilKubernetesAPIServerRolledOut --> DeployingGardenerResourceManagerForVirtualGarden
    DeployingGardenerResourceManagerForVirtualGarden --> WaitingUntilGardenerResourceManagerForVirtualGardenRolledOut
    WaitingUntilMainAndEventETCDsReportReadiness --> DeployingGardenerAPIServer
    WaitingUntilKubernetesAPIServerRolledOut --> DeployingGardenerAPIServer
    WaitingUntilGardenerResourceManagerForVirtualGardenRolledOut --> DeployingGardenerAPIServer
    DeployingGardenerAPIServer --> WaitingUntilGardenerAPIServerRolledOut
    WaitingUntilGardenerAPIServerRolledOut --> DeployingGardenerAdmissionController
    WaitingUntilGardenerAPIServerRolledOut --> DeployingGardenerControllerManager
    WaitingUntilGardenerAPIServerRolledOut --> DeployingGardenerScheduler
    WaitingUntilGardenerAPIServerRolledOut --> DeployingGardenerDiscoveryServer
    DeployingGardenerResourceManagerForVirtualGarden --> DeployingVirtualSystemResources
    WaitingUntilGardenerResourceManagerForVirtualGardenRolledOut --> DeployingResourcesForGardenerOperatorAccessToVirtualGarden
    DeployingKubernetesControllerManager --> RenewingVirtualGardenAccessSecretsAfterCreationOfNewServiceAccountSigningKey
    DeployingResourcesForGardenerOperatorAccessToVirtualGarden --> RenewingVirtualGardenAccessSecretsAfterCreationOfNewServiceAccountSigningKey
    DeployingGardenerAPIServer --> RenewingVirtualGardenAccessSecretsAfterCreationOfNewServiceAccountSigningKey
    DeployingGardenerAdmissionController --> RenewingVirtualGardenAccessSecretsAfterCreationOfNewServiceAccountSigningKey
    DeployingGardenerControllerManager --> RenewingVirtualGardenAccessSecretsAfterCreationOfNewServiceAccountSigningKey
    DeployingGardenerScheduler --> RenewingVirtualGardenAccessSecretsAfterCreationOfNewServiceAccountSigningKey
    DeployingAndWaitingForKubeApiserverServiceInTheRuntimeCluster --> InitializingConnectionToVirtualGardenCluster
    DeployingResourcesForGardenerOperatorAccessToVirtualGarden --> InitializingConnectionToVirtualGardenCluster
    RenewingVirtualGardenAccessSecretsAfterCreationOfNewServiceAccountSigningKey --> InitializingConnectionToVirtualGardenCluster
    WaitingUntilGardenerAPIServerRolledOut --> DeployGardenerInfoConfigMap
    InitializingConnectionToVirtualGardenCluster --> DeployGardenerInfoConfigMap
    InitializingConnectionToVirtualGardenCluster --> DeployingExtensionResources
    DeployingExtensionResources --> WaitingUntilExtensionResourcesAreReady
    InitializingConnectionToVirtualGardenCluster --> DeletingStaleExtensionResources
    DeletingStaleExtensionResources --> WaitingUntilStaleExtensionResourcesAreDeleted
    WaitingUntilGardenerAPIServerRolledOut --> ReconcilingGardenerDashboard
    InitializingConnectionToVirtualGardenCluster --> ReconcilingGardenerDashboard
    WaitingUntilGardenerAPIServerRolledOut --> ReconcilingGardenerDashboardWebTerminalControllerManager
    InitializingConnectionToVirtualGardenCluster --> AnnotateSeedsToTriggerRenewalOfTheirGardenAccessSecrets
    AnnotateSeedsToTriggerRenewalOfTheirGardenAccessSecrets --> CheckIfAllSeedsFinishedTheRenewalOfTheirGardenAccessSecrets
    InitializingConnectionToVirtualGardenCluster --> AnnotateSeedsToTriggerRenewalOfWorkloadIdentityTokens
    WaitingUntilGardenerAPIServerRolledOut --> AnnotateSeedsToTriggerRenewalOfWorkloadIdentityTokens
    CheckIfAllSeedsFinishedTheRenewalOfTheirGardenAccessSecrets --> AnnotateSeedsToTriggerRenewalOfWorkloadIdentityTokens
    AnnotateSeedsToTriggerRenewalOfWorkloadIdentityTokens --> CheckIfAllSeedsFinishedTheRenewalOfTheirWorkloadIdentityTokens
    CheckIfAllSeedsFinishedTheRenewalOfTheirWorkloadIdentityTokens --> AnnotateSeedsToTriggerRenewalOfTheirGardenletKubeconfig
    AnnotateSeedsToTriggerRenewalOfTheirGardenletKubeconfig --> CheckIfAllSeedsFinishedTheRenewalOfTheirGardenletKubeconfig
    InitializingConnectionToVirtualGardenCluster --> LabelingEncryptedResourcesAfterModificationOfEncryptionConfigOrToReEncryptThemWithNewETCDEncryptionKey
    WaitingUntilGardenerAPIServerRolledOut --> LabelingEncryptedResourcesAfterModificationOfEncryptionConfigOrToReEncryptThemWithNewETCDEncryptionKey
    LabelingEncryptedResourcesAfterModificationOfEncryptionConfigOrToReEncryptThemWithNewETCDEncryptionKey --> SnapshottingETCDAfterModificationOfEncryptionConfigOrResourcesAreReEncryptedWithNewETCDEncryptionKey
    InitializingConnectionToVirtualGardenCluster --> RemovingLabelFromReEncryptedResourcesAfterModificationOfEncryptionConfigOrRotationOfETCDEncryptionKey
    WaitingUntilGardenerAPIServerRolledOut --> RemovingLabelFromReEncryptedResourcesAfterModificationOfEncryptionConfigOrRotationOfETCDEncryptionKey
    SnapshottingETCDAfterModificationOfEncryptionConfigOrResourcesAreReEncryptedWithNewETCDEncryptionKey --> RemovingLabelFromReEncryptedResourcesAfterModificationOfEncryptionConfigOrRotationOfETCDEncryptionKey
    GeneratingObservabilityIngressPassword --> DeployingAlertmanager
    DeployingAlertmanager --> WaitingUntilAlertmanagerIsReady
    WaitingUntilGardenerAPIServerRolledOut --> DeployingGardenPrometheus
    InitializingConnectionToVirtualGardenCluster --> DeployingGardenPrometheus
    DeployingGardenPrometheus --> WaitingUntilGardenPrometheusIsReady
    DeployingGardenPrometheus --> DeployingLongTermPrometheus
    DeployingLongTermPrometheus --> WaitingUntilLongTermPrometheusIsReady
    DeployingOpenTelemetryOperator --> DeployingOpenTelemetryCollector
    WaitingUntilKubernetesAPIServerRolledOut --> DeployingBlackboxExporter
    DeployingGardenPrometheus --> DeployingBlackboxExporter
    SyncPointSystemComponents --> DeployingKubeStateMetrics
    WaitingUntilKubernetesAPIServerRolledOut --> DeployingGardenerMetricsExporter
    WaitingUntilGardenerAPIServerRolledOut --> DeployingGardenerMetricsExporter
    GeneratingObservabilityIngressPassword --> DeployingPlutono
    DeployingPlutono --> WaitingUntilPlutonoIsReady
    WaitingUntilAlertmanagerIsReady --> DeployingIstioBasicAuthServer
    WaitingUntilGardenPrometheusIsReady --> DeployingIstioBasicAuthServer
    WaitingUntilLongTermPrometheusIsReady --> DeployingIstioBasicAuthServer
    WaitingUntilPlutonoIsReady --> DeployingIstioBasicAuthServer
```
