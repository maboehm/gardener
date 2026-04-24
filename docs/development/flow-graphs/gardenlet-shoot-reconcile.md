<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# Shoot cluster

```mermaid
---
title: Shoot cluster (runReconcileShootFlow)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    DeployingShootNamespaceInSeed["Deploying Shoot namespace in Seed"]
    EnsuringShootClusterIdentity["Ensuring Shoot cluster identity"]
    DeployingCloudProviderAccountSecret{{"Deploying cloud provider account secret\n[CONDITIONAL]"}}:::conditional
    ReconcileIstioInternalLoadBalancingConfigMap["Reconcile Istio internal load balancing ConfigMap"]
    InitializingSecretsManagement["Initializing secrets management"]
    DeployingInitialShootLoggingStackInSeed["Deploying initial shoot logging stack in Seed"]
    DeployingReferencedResources["Deploying referenced resources"]
    DeployingShootInfrastructure{{"Deploying Shoot infrastructure\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootInfrastructureHasBeenReconciled{{"Waiting until shoot infrastructure has been reconciled\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesAPIServerServiceInTheSeedCluster["Deploying Kubernetes API server service in the Seed cluster"]
    WaitingUntilKubernetesAPIServerServiceInTheSeedClusterHasReportedReadiness{{"Waiting until Kubernetes API server service in the Seed cluster has reported readiness\n[CONDITIONAL]"}}:::conditional
    EnsuringAdvertisedAddressesForTheShoot["Ensuring advertised addresses for the Shoot"]
    DeployingInternalDomainDNSRecord{{"Deploying internal domain DNS record\n[CONDITIONAL]"}}:::conditional
    DeployingExternalDomainDNSRecord{{"Deploying external domain DNS record\n[CONDITIONAL]"}}:::conditional
    DeployingSourceBackupEntry{{"Deploying source backup entry\n[CONDITIONAL]"}}:::conditional
    WaitingUntilTheSourceBackupEntryHasBeenReconciled{{"Waiting until the source backup entry has been reconciled\n[CONDITIONAL]"}}:::conditional
    DeployingBackupEntry{{"Deploying backup entry\n[CONDITIONAL]"}}:::conditional
    WaitingUntilTheBackupEntryHasBeenReconciled{{"Waiting until the backup entry has been reconciled\n[CONDITIONAL]"}}:::conditional
    CopyingEtcdBackupsToNewSeedsBackupBucket{{"Copying etcd backups to new seed's backup bucket\n[CONDITIONAL]"}}:::conditional
    WaitingUntilEtcdBackupsAreCopied{{"Waiting until etcd backups are copied\n[CONDITIONAL]"}}:::conditional
    DestroyingCopyEtcdBackupsTaskResource{{"Destroying copy etcd backups task resource\n[CONDITIONAL]"}}:::conditional
    DeployingMainAndEventsEtcd["Deploying main and events etcd"]
    DestroyingSourceBackupEntry{{"Destroying source backup entry\n[CONDITIONAL]"}}:::conditional
    WaitingUntilSourceBackupEntryHasBeenDeleted{{"Waiting until source backup entry has been deleted\n[CONDITIONAL]"}}:::conditional
    WaitingUntilMainAndEventEtcdReportReadiness{{"Waiting until main and event etcd report readiness\n[CONDITIONAL]"}}:::conditional
    DeployingExtensionResourcesBeforeKubeApiserver{{"Deploying extension resources before kube-apiserver\n[CONDITIONAL]"}}:::conditional
    WaitingUntilExtensionResourcesHandledBeforeKubeApiserverAreReady{{"Waiting until extension resources handled before kube-apiserver are ready\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesAPIServer["Deploying Kubernetes API server"]
    WaitingUntilKubernetesAPIServerRolledOut{{"Waiting until Kubernetes API server rolled out\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesAPIServerServiceSNISettingsInTheSeedCluster["Deploying Kubernetes API server service SNI settings in the Seed cluster"]
    ScalingMainAndEventsEtcdAfterKubeApiserverIsReady{{"Scaling main and events etcd after kube-apiserver is ready\n[CONDITIONAL]"}}:::conditional
    WaitingUntilMainAndEventsEtcdScaledUpAfterKubeApiserverIsReady{{"Waiting until main and events etcd scaled up after kube-apiserver is ready\n[CONDITIONAL]"}}:::conditional
    DeployingGardenerResourceManager["Deploying gardener-resource-manager"]
    WaitingUntilGardenerResourceManagerReportsReadiness{{"Waiting until gardener-resource-manager reports readiness\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesAPIServerWithNodeAgentAuthorizer{{"Deploying Kubernetes API server with node-agent-authorizer\n[CONDITIONAL]"}}:::conditional
    WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut{{"Waiting until Kubernetes API server with node-agent-authorizer rolled out\n[CONDITIONAL]"}}:::conditional
    RenewingShootAccessSecretsAfterCreationOfNewServiceAccountSigningKey{{"Renewing shoot access secrets after creation of new ServiceAccount signing key\n[CONDITIONAL]"}}:::conditional
    DeployingShootControlPlaneComponents{{"Deploying shoot control plane components\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootControlPlaneHasBeenReconciled{{"Waiting until shoot control plane has been reconciled\n[CONDITIONAL]"}}:::conditional
    DeployingShootLoggingStackInSeed["Deploying shoot logging stack in Seed"]
    DeployingShootNamespacesSystemComponent["Deploying shoot namespaces system component"]
    WaitingUntilShootNamespacesHaveBeenReconciled{{"Waiting until shoot namespaces have been reconciled\n[CONDITIONAL]"}}:::conditional
    DeployingVpnSeedServer{{"Deploying vpn-seed-server\n[CONDITIONAL]"}}:::conditional
    DeployingGardenerShootAccessResources["Deploying Gardener shoot access resources"]
    InitializingConnectionToShoot["Initializing connection to Shoot"]
    SyncPublicServiceAccountSigningKeysToGardenCluster{{"Sync public service account signing keys to Garden cluster\n[CONDITIONAL]"}}:::conditional
    LabelingResourcesAfterModificationOfEncryptionConfigOrToEncryptThemWithNewETCDEncryptionKey{{"Labeling resources after modification of encryption config or to encrypt them with new ETCD encryption key\n[CONDITIONAL]"}}:::conditional
    SnapshottingETCDAfterModificationOfEncryptionConfigOrResourcesAreReEncryptedWithNewETCDEncryptionKey{{"Snapshotting ETCD after modification of encryption config or resources are re-encrypted with new ETCD encryption key\n[CONDITIONAL]"}}:::conditional
    RemovingLabelFromResourcesAfterModificationOfEncryptionConfigOrRotationOfETCDEncryptionKey{{"Removing label from resources after modification of encryption config or rotation of ETCD encryption key\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesScheduler{{"Deploying Kubernetes scheduler\n[CONDITIONAL]"}}:::conditional
    ReconcilingKubernetesVerticalPodAutoscaler{{"Reconciling Kubernetes vertical pod autoscaler\n[CONDITIONAL]"}}:::conditional
    DeployingDependencyWatchdogShootAccessResources{{"Deploying dependency-watchdog shoot access resources\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesControllerManager["Deploying Kubernetes controller manager"]
    WaitingUntilKubeControllerManagerReportsReadiness{{"Waiting until kube-controller-manager reports readiness\n[CONDITIONAL]"}}:::conditional
    CreatingNewServiceAccountSecretsAfterCreationOfNewSigningKey{{"Creating new ServiceAccount secrets after creation of new signing key\n[CONDITIONAL]"}}:::conditional
    DeletingOldServiceAccountSecretsAfterRotationOfSigningKey{{"Deleting old ServiceAccount secrets after rotation of signing key\n[CONDITIONAL]"}}:::conditional
    DeletingBastions{{"Deleting Bastions\n[CONDITIONAL]"}}:::conditional
    DeployingOperatingSystemSpecificConfigurationForShootWorkers{{"Deploying operating system specific configuration for shoot workers\n[CONDITIONAL]"}}:::conditional
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled{{"Waiting until operating system configurations for worker nodes have been reconciled\n[CONDITIONAL]"}}:::conditional
    DeleteStaleOperatingSystemConfigResources{{"Delete stale operating system config resources\n[CONDITIONAL]"}}:::conditional
    WaitingUntilStaleOperatingSystemConfigResourcesAreDeleted{{"Waiting until stale operating system config resources are deleted\n[CONDITIONAL]"}}:::conditional
    DeployingShootNetworkPlugin{{"Deploying shoot network plugin\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootNetworkPluginHasBeenReconciled{{"Waiting until shoot network plugin has been reconciled\n[CONDITIONAL]"}}:::conditional
    CheckCoreDNSMigration{{"Check coreDNS migration\n[CONDITIONAL]"}}:::conditional
    DeployingShootClusterIdentity{{"Deploying shoot cluster identity\n[CONDITIONAL]"}}:::conditional
    DeployingShootSystemResources{{"Deploying shoot system resources\n[CONDITIONAL]"}}:::conditional
    PopulatingStaticManifestsFromSeedToShoot["Populating static manifests from seed to shoot"]
    DeployingCoreDNSSystemComponent{{"Deploying CoreDNS system component\n[CONDITIONAL]"}}:::conditional
    ReconcileNodeLocalDnsSystemComponent{{"Reconcile node-local-dns system component\n[CONDITIONAL]"}}:::conditional
    DeployingMetricsServerSystemComponent{{"Deploying metrics-server system component\n[CONDITIONAL]"}}:::conditional
    DeployingVpnShootSystemComponent{{"Deploying vpn-shoot system component\n[CONDITIONAL]"}}:::conditional
    DeployingNodeProblemDetectorSystemComponent{{"Deploying node-problem-detector system component\n[CONDITIONAL]"}}:::conditional
    DeployingKubeProxySystemComponent{{"Deploying kube-proxy system component\n[CONDITIONAL]"}}:::conditional
    DeletingStaleKubeProxyDaemonSets{{"Deleting stale kube-proxy DaemonSets\n[CONDITIONAL]"}}:::conditional
    DeletingKubeProxySystemComponent{{"Deleting kube-proxy system component\n[CONDITIONAL]"}}:::conditional
    DeployingApiserverProxy{{"Deploying apiserver-proxy\n[CONDITIONAL]"}}:::conditional
    DeployingBlackboxExporter{{"Deploying blackbox-exporter\n[CONDITIONAL]"}}:::conditional
    DeployingNodeExporter{{"Deploying node-exporter\n[CONDITIONAL]"}}:::conditional
    DeployingAddonKubernetesDashboard{{"Deploying addon Kubernetes Dashboard\n[CONDITIONAL]"}}:::conditional
    DeployingAddonNginxIngressController{{"Deploying addon Nginx Ingress Controller\n[CONDITIONAL]"}}:::conditional
    DeployingManagedResourcesForTheGardenerNodeAgent{{"Deploying managed resources for the gardener-node-agent\n[CONDITIONAL]"}}:::conditional
    ScalingDownClusterAutoscaler{{"Scaling down cluster autoscaler\n[CONDITIONAL]"}}:::conditional
    DeployingMachineControllerManager{{"Deploying machine-controller-manager\n[CONDITIONAL]"}}:::conditional
    ConfiguringShootWorkerPools{{"Configuring shoot worker pools\n[CONDITIONAL]"}}:::conditional
    WaitingUntilWorkerResourceStatusIsUpdatedWithLatestMachineDeployments{{"Waiting until worker resource status is updated with latest machine deployments\n[CONDITIONAL]"}}:::conditional
    DeployingExtensionResourcesAfterWorkers{{"Deploying extension resources after workers\n[CONDITIONAL]"}}:::conditional
    DeployingClusterAutoscaler{{"Deploying cluster autoscaler\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootWorkerNodesHaveBeenReconciled{{"Waiting until shoot worker nodes have been reconciled\n[CONDITIONAL]"}}:::conditional
    CheckingIfWeHaveDualStackPodCIDRsInNodes{{"Checking if we have dual-stack pod CIDRs in nodes\n[CONDITIONAL]"}}:::conditional
    WaitingUntilExtensionResourcesHandledAfterWorkersAreReady{{"Waiting until extension resources handled after workers are ready\n[CONDITIONAL]"}}:::conditional
    ScalingDownMachineControllerManager{{"Scaling down machine-controller-manager\n[CONDITIONAL]"}}:::conditional
    ReconcilingPlutonoForShootInSeedForTheLoggingStack["Reconciling Plutono for Shoot in Seed for the logging stack"]
    WaitingUntilNginxIngressLoadBalancerIsReady{{"Waiting until nginx ingress LoadBalancer is ready\n[CONDITIONAL]"}}:::conditional
    DeployingNginxIngressDNSRecord{{"Deploying nginx ingress DNS record\n[CONDITIONAL]"}}:::conditional
    WaitingUntilTheKubernetesAPIServerCanConnectToTheShootWorkers{{"Waiting until the Kubernetes API server can connect to the Shoot workers\n[CONDITIONAL]"}}:::conditional
    WaitingUntilAllShootWorkerNodesHaveUpdatedTheOperatingSystemConfig{{"Waiting until all shoot worker nodes have updated the operating system config\n[CONDITIONAL]"}}:::conditional
    ReconcilingShootAlertmanager["Reconciling Shoot Alertmanager"]
    WaitingUntilShootAlertmanagerIsReconciled["Waiting until Shoot Alertmanager is reconciled"]
    ReconcilingShootPrometheus["Reconciling Shoot Prometheus"]
    WaitingUntilShootPrometheusIsReconciled["Waiting until Shoot Prometheus is reconciled"]
    DeployingControlPlaneBlackboxExporter["Deploying control plane blackbox-exporter"]
    ReconcilingKubeStateMetricsForShootInSeedForTheMonitoringStack{{"Reconciling kube-state-metrics for Shoot in Seed for the monitoring stack\n[CONDITIONAL]"}}:::conditional
    ReconcilingPlutonoForShootInSeedForTheMonitoringStack["Reconciling Plutono for Shoot in Seed for the monitoring stack"]
    WaitingUntilPlutonoForShootInSeedIsReconciled["Waiting until Plutono for Shoot in Seed is reconciled"]
    DeployingIstioBasicAuthServer["Deploying istio-basic-auth-server"]
    HibernatingControlPlane{{"Hibernating control plane\n[CONDITIONAL]"}}:::conditional
    HibernatingExtensionResourcesAfterKubeApiserverHibernation{{"Hibernating extension resources after kube-apiserver hibernation\n[CONDITIONAL]"}}:::conditional
    WaitingUntilExtensionResourcesHibernatedAfterKubeApiserverHibernationAreReady{{"Waiting until extension resources hibernated after kube-apiserver hibernation are ready\n[CONDITIONAL]"}}:::conditional
    DestroyingIngressDomainDNSRecordIfHibernated{{"Destroying ingress domain DNS record if hibernated\n[CONDITIONAL]"}}:::conditional
    DestroyingExternalDomainDNSRecordIfHibernated{{"Destroying external domain DNS record if hibernated\n[CONDITIONAL]"}}:::conditional
    DestroyingInternalDomainDNSRecordIfHibernated{{"Destroying internal domain DNS record if hibernated\n[CONDITIONAL]"}}:::conditional
    DeletingStaleExtensionResources["Deleting stale extension resources"]
    WaitingUntilStaleExtensionResourcesAreDeleted{{"Waiting until stale extension resources are deleted\n[CONDITIONAL]"}}:::conditional
    DeployingContainerRuntimeResources{{"Deploying container runtime resources\n[CONDITIONAL]"}}:::conditional
    WaitingUntilContainerRuntimeResourcesAreReady{{"Waiting until container runtime resources are ready\n[CONDITIONAL]"}}:::conditional
    DeletingStaleContainerRuntimeResources{{"Deleting stale container runtime resources\n[CONDITIONAL]"}}:::conditional
    WaitingUntilStaleContainerRuntimeResourcesAreDeleted{{"Waiting until stale container runtime resources are deleted\n[CONDITIONAL]"}}:::conditional
    RestartingControlPlanePods{{"Restarting control plane pods\n[CONDITIONAL]"}}:::conditional
    SyncPointAllSystemComponentsDeployed(["Sync: All System Components Deployed"]):::syncpoint

    WaitingUntilShootNetworkPluginHasBeenReconciled --> SyncPointAllSystemComponentsDeployed
    DeployingApiserverProxy --> SyncPointAllSystemComponentsDeployed
    DeployingShootSystemResources --> SyncPointAllSystemComponentsDeployed
    DeployingCoreDNSSystemComponent --> SyncPointAllSystemComponentsDeployed
    DeployingNodeExporter --> SyncPointAllSystemComponentsDeployed
    ReconcileNodeLocalDnsSystemComponent --> SyncPointAllSystemComponentsDeployed
    DeployingMetricsServerSystemComponent --> SyncPointAllSystemComponentsDeployed
    DeployingVpnShootSystemComponent --> SyncPointAllSystemComponentsDeployed
    DeployingNodeProblemDetectorSystemComponent --> SyncPointAllSystemComponentsDeployed
    DeployingKubeProxySystemComponent --> SyncPointAllSystemComponentsDeployed
    DeployingBlackboxExporter --> SyncPointAllSystemComponentsDeployed
    DeployingAddonKubernetesDashboard --> SyncPointAllSystemComponentsDeployed
    DeployingAddonNginxIngressController --> SyncPointAllSystemComponentsDeployed
    DeployingShootNamespaceInSeed --> EnsuringShootClusterIdentity
    DeployingShootNamespaceInSeed --> DeployingCloudProviderAccountSecret
    DeployingShootNamespaceInSeed --> ReconcileIstioInternalLoadBalancingConfigMap
    DeployingShootNamespaceInSeed --> InitializingSecretsManagement
    ReconcileIstioInternalLoadBalancingConfigMap --> InitializingSecretsManagement
    DeployingShootNamespaceInSeed --> DeployingInitialShootLoggingStackInSeed
    InitializingSecretsManagement --> DeployingInitialShootLoggingStackInSeed
    DeployingShootNamespaceInSeed --> DeployingReferencedResources
    InitializingSecretsManagement --> DeployingShootInfrastructure
    DeployingCloudProviderAccountSecret --> DeployingShootInfrastructure
    DeployingReferencedResources --> DeployingShootInfrastructure
    DeployingShootInfrastructure --> WaitingUntilShootInfrastructureHasBeenReconciled
    DeployingShootNamespaceInSeed --> DeployingKubernetesAPIServerServiceInTheSeedCluster
    EnsuringShootClusterIdentity --> DeployingKubernetesAPIServerServiceInTheSeedCluster
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeployingKubernetesAPIServerServiceInTheSeedCluster
    DeployingKubernetesAPIServerServiceInTheSeedCluster --> WaitingUntilKubernetesAPIServerServiceInTheSeedClusterHasReportedReadiness
    InitializingSecretsManagement --> EnsuringAdvertisedAddressesForTheShoot
    WaitingUntilKubernetesAPIServerServiceInTheSeedClusterHasReportedReadiness --> EnsuringAdvertisedAddressesForTheShoot
    DeployingReferencedResources --> DeployingInternalDomainDNSRecord
    WaitingUntilKubernetesAPIServerServiceInTheSeedClusterHasReportedReadiness --> DeployingInternalDomainDNSRecord
    DeployingReferencedResources --> DeployingExternalDomainDNSRecord
    WaitingUntilKubernetesAPIServerServiceInTheSeedClusterHasReportedReadiness --> DeployingExternalDomainDNSRecord
    DeployingShootNamespaceInSeed --> DeployingSourceBackupEntry
    DeployingSourceBackupEntry --> WaitingUntilTheSourceBackupEntryHasBeenReconciled
    DeployingShootNamespaceInSeed --> DeployingBackupEntry
    WaitingUntilTheSourceBackupEntryHasBeenReconciled --> DeployingBackupEntry
    DeployingBackupEntry --> WaitingUntilTheBackupEntryHasBeenReconciled
    InitializingSecretsManagement --> CopyingEtcdBackupsToNewSeedsBackupBucket
    DeployingCloudProviderAccountSecret --> CopyingEtcdBackupsToNewSeedsBackupBucket
    WaitingUntilTheBackupEntryHasBeenReconciled --> CopyingEtcdBackupsToNewSeedsBackupBucket
    WaitingUntilTheSourceBackupEntryHasBeenReconciled --> CopyingEtcdBackupsToNewSeedsBackupBucket
    CopyingEtcdBackupsToNewSeedsBackupBucket --> WaitingUntilEtcdBackupsAreCopied
    WaitingUntilEtcdBackupsAreCopied --> DestroyingCopyEtcdBackupsTaskResource
    InitializingSecretsManagement --> DeployingMainAndEventsEtcd
    DeployingCloudProviderAccountSecret --> DeployingMainAndEventsEtcd
    WaitingUntilTheBackupEntryHasBeenReconciled --> DeployingMainAndEventsEtcd
    WaitingUntilEtcdBackupsAreCopied --> DeployingMainAndEventsEtcd
    DeployingMainAndEventsEtcd --> DestroyingSourceBackupEntry
    DestroyingSourceBackupEntry --> WaitingUntilSourceBackupEntryHasBeenDeleted
    DeployingMainAndEventsEtcd --> WaitingUntilMainAndEventEtcdReportReadiness
    InitializingSecretsManagement --> DeployingExtensionResourcesBeforeKubeApiserver
    DeployingCloudProviderAccountSecret --> DeployingExtensionResourcesBeforeKubeApiserver
    DeployingReferencedResources --> DeployingExtensionResourcesBeforeKubeApiserver
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeployingExtensionResourcesBeforeKubeApiserver
    DeployingExtensionResourcesBeforeKubeApiserver --> WaitingUntilExtensionResourcesHandledBeforeKubeApiserverAreReady
    InitializingSecretsManagement --> DeployingKubernetesAPIServer
    DeployingMainAndEventsEtcd --> DeployingKubernetesAPIServer
    WaitingUntilMainAndEventEtcdReportReadiness --> DeployingKubernetesAPIServer
    WaitingUntilKubernetesAPIServerServiceInTheSeedClusterHasReportedReadiness --> DeployingKubernetesAPIServer
    WaitingUntilExtensionResourcesHandledBeforeKubeApiserverAreReady --> DeployingKubernetesAPIServer
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeployingKubernetesAPIServer
    DeployingKubernetesAPIServer --> WaitingUntilKubernetesAPIServerRolledOut
    WaitingUntilKubernetesAPIServerRolledOut --> DeployingKubernetesAPIServerServiceSNISettingsInTheSeedCluster
    WaitingUntilMainAndEventEtcdReportReadiness --> ScalingMainAndEventsEtcdAfterKubeApiserverIsReady
    WaitingUntilKubernetesAPIServerRolledOut --> ScalingMainAndEventsEtcdAfterKubeApiserverIsReady
    ScalingMainAndEventsEtcdAfterKubeApiserverIsReady --> WaitingUntilMainAndEventsEtcdScaledUpAfterKubeApiserverIsReady
    WaitingUntilKubernetesAPIServerRolledOut --> DeployingGardenerResourceManager
    DeployingGardenerResourceManager --> WaitingUntilGardenerResourceManagerReportsReadiness
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingKubernetesAPIServerWithNodeAgentAuthorizer
    DeployingKubernetesAPIServerWithNodeAgentAuthorizer --> WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut
    WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut --> RenewingShootAccessSecretsAfterCreationOfNewServiceAccountSigningKey
    WaitingUntilGardenerResourceManagerReportsReadiness --> RenewingShootAccessSecretsAfterCreationOfNewServiceAccountSigningKey
    WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut --> DeployingShootControlPlaneComponents
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingShootControlPlaneComponents
    DeployingShootControlPlaneComponents --> WaitingUntilShootControlPlaneHasBeenReconciled
    DeployingInitialShootLoggingStackInSeed --> DeployingShootLoggingStackInSeed
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingShootLoggingStackInSeed
    DeployingGardenerResourceManager --> DeployingShootNamespacesSystemComponent
    WaitingUntilGardenerResourceManagerReportsReadiness --> WaitingUntilShootNamespacesHaveBeenReconciled
    DeployingShootNamespacesSystemComponent --> WaitingUntilShootNamespacesHaveBeenReconciled
    InitializingSecretsManagement --> DeployingVpnSeedServer
    DeployingShootNamespaceInSeed --> DeployingVpnSeedServer
    WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut --> DeployingVpnSeedServer
    InitializingSecretsManagement --> DeployingGardenerShootAccessResources
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingGardenerShootAccessResources
    WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut --> InitializingConnectionToShoot
    DeployingInternalDomainDNSRecord --> InitializingConnectionToShoot
    DeployingGardenerShootAccessResources --> InitializingConnectionToShoot
    InitializingConnectionToShoot --> SyncPublicServiceAccountSigningKeysToGardenCluster
    InitializingConnectionToShoot --> LabelingResourcesAfterModificationOfEncryptionConfigOrToEncryptThemWithNewETCDEncryptionKey
    LabelingResourcesAfterModificationOfEncryptionConfigOrToEncryptThemWithNewETCDEncryptionKey --> SnapshottingETCDAfterModificationOfEncryptionConfigOrResourcesAreReEncryptedWithNewETCDEncryptionKey
    InitializingConnectionToShoot --> RemovingLabelFromResourcesAfterModificationOfEncryptionConfigOrRotationOfETCDEncryptionKey
    SnapshottingETCDAfterModificationOfEncryptionConfigOrResourcesAreReEncryptedWithNewETCDEncryptionKey --> RemovingLabelFromResourcesAfterModificationOfEncryptionConfigOrRotationOfETCDEncryptionKey
    InitializingSecretsManagement --> DeployingKubernetesScheduler
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingKubernetesScheduler
    InitializingSecretsManagement --> ReconcilingKubernetesVerticalPodAutoscaler
    WaitingUntilGardenerResourceManagerReportsReadiness --> ReconcilingKubernetesVerticalPodAutoscaler
    InitializingSecretsManagement --> DeployingDependencyWatchdogShootAccessResources
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingDependencyWatchdogShootAccessResources
    InitializingSecretsManagement --> DeployingKubernetesControllerManager
    DeployingCloudProviderAccountSecret --> DeployingKubernetesControllerManager
    WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut --> DeployingKubernetesControllerManager
    DeployingKubernetesControllerManager --> WaitingUntilKubeControllerManagerReportsReadiness
    InitializingConnectionToShoot --> CreatingNewServiceAccountSecretsAfterCreationOfNewSigningKey
    WaitingUntilKubeControllerManagerReportsReadiness --> CreatingNewServiceAccountSecretsAfterCreationOfNewSigningKey
    InitializingConnectionToShoot --> DeletingOldServiceAccountSecretsAfterRotationOfSigningKey
    WaitingUntilKubeControllerManagerReportsReadiness --> DeletingOldServiceAccountSecretsAfterRotationOfSigningKey
    DeployingReferencedResources --> DeletingBastions
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeletingBastions
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeletingBastions
    DeployingReferencedResources --> DeployingOperatingSystemSpecificConfigurationForShootWorkers
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeployingOperatingSystemSpecificConfigurationForShootWorkers
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingOperatingSystemSpecificConfigurationForShootWorkers
    DeletingBastions --> DeployingOperatingSystemSpecificConfigurationForShootWorkers
    DeployingOperatingSystemSpecificConfigurationForShootWorkers --> WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled
    DeployingOperatingSystemSpecificConfigurationForShootWorkers --> DeleteStaleOperatingSystemConfigResources
    DeleteStaleOperatingSystemConfigResources --> WaitingUntilStaleOperatingSystemConfigResourcesAreDeleted
    DeployingReferencedResources --> DeployingShootNetworkPlugin
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingShootNetworkPlugin
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> DeployingShootNetworkPlugin
    DeployingKubernetesScheduler --> DeployingShootNetworkPlugin
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingShootNetworkPlugin
    DeployingShootNetworkPlugin --> WaitingUntilShootNetworkPluginHasBeenReconciled
    WaitingUntilShootNetworkPluginHasBeenReconciled --> CheckCoreDNSMigration
    DeployingGardenerResourceManager --> DeployingShootClusterIdentity
    EnsuringShootClusterIdentity --> DeployingShootClusterIdentity
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> DeployingShootClusterIdentity
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingShootSystemResources
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingShootSystemResources
    InitializingConnectionToShoot --> DeployingShootSystemResources
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> DeployingShootSystemResources
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingShootSystemResources
    WaitingUntilGardenerResourceManagerReportsReadiness --> PopulatingStaticManifestsFromSeedToShoot
    WaitingUntilShootNamespacesHaveBeenReconciled --> PopulatingStaticManifestsFromSeedToShoot
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingCoreDNSSystemComponent
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingCoreDNSSystemComponent
    InitializingConnectionToShoot --> DeployingCoreDNSSystemComponent
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> DeployingCoreDNSSystemComponent
    DeployingKubernetesScheduler --> DeployingCoreDNSSystemComponent
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingCoreDNSSystemComponent
    WaitingUntilShootControlPlaneHasBeenReconciled --> ReconcileNodeLocalDnsSystemComponent
    DeployingGardenerResourceManager --> ReconcileNodeLocalDnsSystemComponent
    InitializingConnectionToShoot --> ReconcileNodeLocalDnsSystemComponent
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> ReconcileNodeLocalDnsSystemComponent
    DeployingKubernetesScheduler --> ReconcileNodeLocalDnsSystemComponent
    WaitingUntilShootNamespacesHaveBeenReconciled --> ReconcileNodeLocalDnsSystemComponent
    WaitingUntilShootNetworkPluginHasBeenReconciled --> ReconcileNodeLocalDnsSystemComponent
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingMetricsServerSystemComponent
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingMetricsServerSystemComponent
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> DeployingMetricsServerSystemComponent
    DeployingKubernetesScheduler --> DeployingMetricsServerSystemComponent
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingMetricsServerSystemComponent
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingVpnShootSystemComponent
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingVpnShootSystemComponent
    DeployingGardenerResourceManager --> DeployingVpnShootSystemComponent
    DeployingKubernetesScheduler --> DeployingVpnShootSystemComponent
    DeployingVpnSeedServer --> DeployingVpnShootSystemComponent
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingVpnShootSystemComponent
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingNodeProblemDetectorSystemComponent
    DeployingGardenerResourceManager --> DeployingNodeProblemDetectorSystemComponent
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> DeployingNodeProblemDetectorSystemComponent
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingNodeProblemDetectorSystemComponent
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingKubeProxySystemComponent
    DeployingGardenerResourceManager --> DeployingKubeProxySystemComponent
    InitializingConnectionToShoot --> DeployingKubeProxySystemComponent
    EnsuringShootClusterIdentity --> DeployingKubeProxySystemComponent
    DeployingKubernetesScheduler --> DeployingKubeProxySystemComponent
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingKubeProxySystemComponent
    DeployingKubeProxySystemComponent --> DeletingStaleKubeProxyDaemonSets
    DeployingGardenerResourceManager --> DeletingKubeProxySystemComponent
    InitializingConnectionToShoot --> DeletingKubeProxySystemComponent
    EnsuringShootClusterIdentity --> DeletingKubeProxySystemComponent
    DeployingKubernetesScheduler --> DeletingKubeProxySystemComponent
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingApiserverProxy
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingApiserverProxy
    InitializingConnectionToShoot --> DeployingApiserverProxy
    EnsuringShootClusterIdentity --> DeployingApiserverProxy
    DeployingKubernetesScheduler --> DeployingApiserverProxy
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingApiserverProxy
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingBlackboxExporter
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingBlackboxExporter
    InitializingConnectionToShoot --> DeployingBlackboxExporter
    EnsuringShootClusterIdentity --> DeployingBlackboxExporter
    DeployingKubernetesScheduler --> DeployingBlackboxExporter
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingBlackboxExporter
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingNodeExporter
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingNodeExporter
    InitializingConnectionToShoot --> DeployingNodeExporter
    EnsuringShootClusterIdentity --> DeployingNodeExporter
    DeployingKubernetesScheduler --> DeployingNodeExporter
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingNodeExporter
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingAddonKubernetesDashboard
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingAddonKubernetesDashboard
    InitializingConnectionToShoot --> DeployingAddonKubernetesDashboard
    EnsuringShootClusterIdentity --> DeployingAddonKubernetesDashboard
    DeployingKubernetesScheduler --> DeployingAddonKubernetesDashboard
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingAddonKubernetesDashboard
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingAddonNginxIngressController
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingAddonNginxIngressController
    InitializingConnectionToShoot --> DeployingAddonNginxIngressController
    EnsuringShootClusterIdentity --> DeployingAddonNginxIngressController
    DeployingKubernetesScheduler --> DeployingAddonNginxIngressController
    WaitingUntilShootNamespacesHaveBeenReconciled --> DeployingAddonNginxIngressController
    DeployingGardenerResourceManager --> DeployingManagedResourcesForTheGardenerNodeAgent
    EnsuringShootClusterIdentity --> DeployingManagedResourcesForTheGardenerNodeAgent
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> DeployingManagedResourcesForTheGardenerNodeAgent
    DeployingManagedResourcesForTheGardenerNodeAgent --> ScalingDownClusterAutoscaler
    DeployingCloudProviderAccountSecret --> DeployingMachineControllerManager
    DeployingReferencedResources --> DeployingMachineControllerManager
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeployingMachineControllerManager
    InitializingConnectionToShoot --> DeployingMachineControllerManager
    WaitingUntilOperatingSystemConfigurationsForWorkerNodesHaveBeenReconciled --> DeployingMachineControllerManager
    WaitingUntilShootNetworkPluginHasBeenReconciled --> DeployingMachineControllerManager
    CreatingNewServiceAccountSecretsAfterCreationOfNewSigningKey --> DeployingMachineControllerManager
    ScalingDownClusterAutoscaler --> DeployingMachineControllerManager
    DeployingMachineControllerManager --> ConfiguringShootWorkerPools
    ConfiguringShootWorkerPools --> WaitingUntilWorkerResourceStatusIsUpdatedWithLatestMachineDeployments
    WaitingUntilWorkerResourceStatusIsUpdatedWithLatestMachineDeployments --> DeployingExtensionResourcesAfterWorkers
    WaitingUntilWorkerResourceStatusIsUpdatedWithLatestMachineDeployments --> DeployingClusterAutoscaler
    DeployingManagedResourcesForTheGardenerNodeAgent --> DeployingClusterAutoscaler
    ConfiguringShootWorkerPools --> WaitingUntilShootWorkerNodesHaveBeenReconciled
    WaitingUntilWorkerResourceStatusIsUpdatedWithLatestMachineDeployments --> WaitingUntilShootWorkerNodesHaveBeenReconciled
    DeployingManagedResourcesForTheGardenerNodeAgent --> WaitingUntilShootWorkerNodesHaveBeenReconciled
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> CheckingIfWeHaveDualStackPodCIDRsInNodes
    DeployingExtensionResourcesAfterWorkers --> WaitingUntilExtensionResourcesHandledAfterWorkersAreReady
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> ScalingDownMachineControllerManager
    DeployingShootLoggingStackInSeed --> ReconcilingPlutonoForShootInSeedForTheLoggingStack
    InitializingConnectionToShoot --> WaitingUntilNginxIngressLoadBalancerIsReady
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> WaitingUntilNginxIngressLoadBalancerIsReady
    EnsuringShootClusterIdentity --> WaitingUntilNginxIngressLoadBalancerIsReady
    WaitingUntilNginxIngressLoadBalancerIsReady --> DeployingNginxIngressDNSRecord
    SyncPointAllSystemComponentsDeployed --> WaitingUntilTheKubernetesAPIServerCanConnectToTheShootWorkers
    WaitingUntilShootNetworkPluginHasBeenReconciled --> WaitingUntilTheKubernetesAPIServerCanConnectToTheShootWorkers
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> WaitingUntilTheKubernetesAPIServerCanConnectToTheShootWorkers
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> WaitingUntilAllShootWorkerNodesHaveUpdatedTheOperatingSystemConfig
    WaitingUntilTheKubernetesAPIServerCanConnectToTheShootWorkers --> WaitingUntilAllShootWorkerNodesHaveUpdatedTheOperatingSystemConfig
    InitializingConnectionToShoot --> ReconcilingShootAlertmanager
    WaitingUntilTheKubernetesAPIServerCanConnectToTheShootWorkers --> ReconcilingShootAlertmanager
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> ReconcilingShootAlertmanager
    WaitingUntilShootInfrastructureHasBeenReconciled --> ReconcilingShootAlertmanager
    ReconcilingShootAlertmanager --> WaitingUntilShootAlertmanagerIsReconciled
    InitializingConnectionToShoot --> ReconcilingShootPrometheus
    WaitingUntilTheKubernetesAPIServerCanConnectToTheShootWorkers --> ReconcilingShootPrometheus
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> ReconcilingShootPrometheus
    WaitingUntilShootInfrastructureHasBeenReconciled --> ReconcilingShootPrometheus
    ReconcilingShootPrometheus --> WaitingUntilShootPrometheusIsReconciled
    InitializingConnectionToShoot --> DeployingControlPlaneBlackboxExporter
    WaitingUntilTheKubernetesAPIServerCanConnectToTheShootWorkers --> DeployingControlPlaneBlackboxExporter
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> DeployingControlPlaneBlackboxExporter
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeployingControlPlaneBlackboxExporter
    ReconcilingShootPrometheus --> ReconcilingKubeStateMetricsForShootInSeedForTheMonitoringStack
    ReconcilingShootAlertmanager --> ReconcilingKubeStateMetricsForShootInSeedForTheMonitoringStack
    ReconcilingShootPrometheus --> ReconcilingPlutonoForShootInSeedForTheMonitoringStack
    ReconcilingShootAlertmanager --> ReconcilingPlutonoForShootInSeedForTheMonitoringStack
    ReconcilingPlutonoForShootInSeedForTheLoggingStack --> WaitingUntilPlutonoForShootInSeedIsReconciled
    ReconcilingPlutonoForShootInSeedForTheMonitoringStack --> WaitingUntilPlutonoForShootInSeedIsReconciled
    WaitingUntilShootAlertmanagerIsReconciled --> DeployingIstioBasicAuthServer
    WaitingUntilShootPrometheusIsReconciled --> DeployingIstioBasicAuthServer
    WaitingUntilPlutonoForShootInSeedIsReconciled --> DeployingIstioBasicAuthServer
    InitializingConnectionToShoot --> HibernatingControlPlane
    ReconcilingShootPrometheus --> HibernatingControlPlane
    ReconcilingShootAlertmanager --> HibernatingControlPlane
    DeployingShootLoggingStackInSeed --> HibernatingControlPlane
    DeployingClusterAutoscaler --> HibernatingControlPlane
    WaitingUntilShootWorkerNodesHaveBeenReconciled --> HibernatingControlPlane
    WaitingUntilMainAndEventsEtcdScaledUpAfterKubeApiserverIsReady --> HibernatingControlPlane
    HibernatingControlPlane --> HibernatingExtensionResourcesAfterKubeApiserverHibernation
    HibernatingExtensionResourcesAfterKubeApiserverHibernation --> WaitingUntilExtensionResourcesHibernatedAfterKubeApiserverHibernationAreReady
    HibernatingControlPlane --> DestroyingIngressDomainDNSRecordIfHibernated
    HibernatingControlPlane --> DestroyingExternalDomainDNSRecordIfHibernated
    HibernatingControlPlane --> DestroyingInternalDomainDNSRecordIfHibernated
    InitializingConnectionToShoot --> DeletingStaleExtensionResources
    DeletingStaleExtensionResources --> WaitingUntilStaleExtensionResourcesAreDeleted
    DeployingReferencedResources --> DeployingContainerRuntimeResources
    InitializingConnectionToShoot --> DeployingContainerRuntimeResources
    DeployingContainerRuntimeResources --> WaitingUntilContainerRuntimeResourcesAreReady
    InitializingConnectionToShoot --> DeletingStaleContainerRuntimeResources
    DeletingStaleContainerRuntimeResources --> WaitingUntilStaleContainerRuntimeResourcesAreDeleted
    DeployingKubernetesControllerManager --> RestartingControlPlanePods
    DeployingShootControlPlaneComponents --> RestartingControlPlanePods
```
