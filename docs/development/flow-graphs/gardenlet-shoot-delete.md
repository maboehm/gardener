<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# Shoot cluster deletion

```mermaid
---
title: Shoot cluster deletion (runDeleteShootFlow)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    DeployingShootNamespaceInSeed{{"Deploying Shoot namespace in Seed\n[CONDITIONAL]"}}:::conditional
    EnsuringShootClusterIdentity["Ensuring Shoot cluster identity"]
    DeployingCloudProviderAccountSecret{{"Deploying cloud provider account secret\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesAPIServerServiceInTheSeedCluster{{"Deploying Kubernetes API server service in the Seed cluster\n[CONDITIONAL]"}}:::conditional
    WaitingUntilKubernetesAPILoadBalancerInTheSeedClusterHasReportedReadiness{{"Waiting until Kubernetes API LoadBalancer in the Seed cluster has reported readiness\n[CONDITIONAL]"}}:::conditional
    ReconcileIstioInternalLoadBalancingConfigMap{{"Reconcile Istio internal load balancing ConfigMap\n[CONDITIONAL]"}}:::conditional
    InitializingSecretsManagement{{"Initializing secrets management\n[CONDITIONAL]"}}:::conditional
    EnsuringAdvertisedAddressesForTheShoot["Ensuring advertised addresses for the Shoot"]
    DeployingReferencedResources{{"Deploying referenced resources\n[CONDITIONAL]"}}:::conditional
    DeployingInternalDomainDNSRecord{{"Deploying internal domain DNS record\n[CONDITIONAL]"}}:::conditional
    DeployingMainAndEventsEtcd{{"Deploying main and events etcd\n[CONDITIONAL]"}}:::conditional
    ScalingUpEtcdMainAndEvent{{"Scaling up etcd main and event\n[CONDITIONAL]"}}:::conditional
    WaitingUntilMainAndEventEtcdReportReadiness{{"Waiting until main and event etcd report readiness\n[CONDITIONAL]"}}:::conditional
    DeployingShootControlPlane{{"Deploying Shoot control plane\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootControlPlaneHasBeenReconciled{{"Waiting until Shoot control plane has been reconciled\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesAPIServer{{"Deploying Kubernetes API server\n[CONDITIONAL]"}}:::conditional
    ScalingUpKubernetesAPIServer{{"Scaling up Kubernetes API server\n[CONDITIONAL]"}}:::conditional
    WaitingUntilKubernetesAPIServerReportsReadiness{{"Waiting until Kubernetes API server reports readiness\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesAPIServerServiceSNISettingsInTheSeedCluster{{"Deploying Kubernetes API server service SNI settings in the Seed cluster\n[CONDITIONAL]"}}:::conditional
    SettingGardenerResourceManagerReplicasTo2{{"Setting gardener-resource-manager replicas to 2\n[CONDITIONAL]"}}:::conditional
    DeployingGardenerResourceManager{{"Deploying gardener-resource-manager\n[CONDITIONAL]"}}:::conditional
    WaitingUntilGardenerResourceManagerReportsReadiness{{"Waiting until gardener-resource-manager reports readiness\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesAPIServerWithNodeAgentAuthorizer{{"Deploying Kubernetes API server with node-agent-authorizer\n[CONDITIONAL]"}}:::conditional
    WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut{{"Waiting until Kubernetes API server with node-agent-authorizer rolled out\n[CONDITIONAL]"}}:::conditional
    DeployingGardenerShootAccessResources{{"Deploying Gardener shoot access resources\n[CONDITIONAL]"}}:::conditional
    InitializingConnectionToShoot{{"Initializing connection to Shoot\n[CONDITIONAL]"}}:::conditional
    DeployingKubernetesControllerManager{{"Deploying Kubernetes controller manager\n[CONDITIONAL]"}}:::conditional
    ScalingUpKubernetesControllerManager{{"Scaling up Kubernetes controller manager\n[CONDITIONAL]"}}:::conditional
    DeletingClusterAutoscaler{{"Deleting cluster autoscaler\n[CONDITIONAL]"}}:::conditional
    CleaningUpWebhooks{{"Cleaning up webhooks\n[CONDITIONAL]"}}:::conditional
    WaitingUntilKubeControllerManagerIsActive{{"Waiting until kube-controller-manager is active\n[CONDITIONAL]"}}:::conditional
    CleaningExtendedAPIGroups{{"Cleaning extended API groups\n[CONDITIONAL]"}}:::conditional
    CleaningKubernetesResources{{"Cleaning Kubernetes resources\n[CONDITIONAL]"}}:::conditional
    DeletingMetricsServer{{"Deleting metrics-server\n[CONDITIONAL]"}}:::conditional
    DestroyingShootNetworkPlugin{{"Destroying shoot network plugin\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootNetworkPluginHasBeenDestroyed{{"Waiting until shoot network plugin has been destroyed\n[CONDITIONAL]"}}:::conditional
    DeployingMachineControllerManager{{"Deploying machine-controller-manager\n[CONDITIONAL]"}}:::conditional
    DestroyingShootWorkers{{"Destroying shoot workers\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootWorkerNodesHaveBeenTerminated{{"Waiting until shoot worker nodes have been terminated\n[CONDITIONAL]"}}:::conditional
    DeletingMachineControllerManager{{"Deleting machine-controller-manager\n[CONDITIONAL]"}}:::conditional
    DeletingOperatingSystemConfigResources{{"Deleting operating system config resources\n[CONDITIONAL]"}}:::conditional
    WaitingUntilAllOperatingSystemConfigResourcesAreDeleted{{"Waiting until all operating system config resources are deleted\n[CONDITIONAL]"}}:::conditional
    DeletingManagedResources["Deleting managed resources"]
    DeletingDWDManagedResourceAndSecrets{{"Deleting DWD managed resource and secrets\n[CONDITIONAL]"}}:::conditional
    WaitingUntilManagedResourcesHaveBeenDeleted["Waiting until managed resources have been deleted"]
    DeletingExtensionResourcesBeforeKubeApiserver["Deleting extension resources before kube-apiserver"]
    WaitingUntilExtensionResourcesThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted["Waiting until extension resources that should be handled before kube-apiserver have been deleted"]
    DeletingStaleExtensionResources["Deleting stale extension resources"]
    WaitingUntilAllStaleExtensionResourcesHaveBeenDeleted["Waiting until all stale extension resources have been deleted"]
    DeletingContainerRuntimeResources{{"Deleting container runtime resources\n[CONDITIONAL]"}}:::conditional
    WaitingUntilStaleContainerRuntimeResourcesAreDeleted{{"Waiting until stale container runtime resources are deleted\n[CONDITIONAL]"}}:::conditional
    DestroyingShootControlPlane{{"Destroying shoot control plane\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootControlPlaneHasBeenDestroyed{{"Waiting until shoot control plane has been destroyed\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootManagedResourcesHaveBeenDeleted["Waiting until shoot managed resources have been deleted"]
    DeletingKubernetesAPIServer["Deleting Kubernetes API server"]
    WaitingUntilKubernetesAPIServerHasBeenDeleted["Waiting until Kubernetes API server has been deleted"]
    DeletingExtensionResourcesAfterKubeApiserver["Deleting extension resources after kube-apiserver"]
    WaitingUntilExtensionResourcesThatShouldBeHandledAfterKubeApiserverHaveBeenDeleted["Waiting until extension resources that should be handled after kube-apiserver have been deleted"]
    WaitingUntilAllExtensionResourcesHaveBeenDeleted["Waiting until all extension resources have been deleted"]
    DestroyingKubernetesAPIServerServiceSNI["Destroying Kubernetes API server service SNI"]
    DestroyingKubernetesAPIServerService["Destroying Kubernetes API server service"]
    DestroyingGardenerResourceManager["Destroying gardener-resource-manager"]
    DeletePublicServiceAccountSigningKeysFromGardenCluster["Delete public service account signing keys from Garden cluster"]
    DestroyingNginxIngressDNSRecord{{"Destroying nginx ingress DNS record\n[CONDITIONAL]"}}:::conditional
    DestroyingShootInfrastructure{{"Destroying shoot infrastructure\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootInfrastructureHasBeenDeleted{{"Waiting until shoot infrastructure has been deleted\n[CONDITIONAL]"}}:::conditional
    DestroyingExternalDomainDNSRecord{{"Destroying external domain DNS record\n[CONDITIONAL]"}}:::conditional
    DeletingShootAlertmanager["Deleting Shoot Alertmanager"]
    DeletingShootPrometheus["Deleting Shoot Prometheus"]
    DestroyingControlPlaneBlackboxExporter["Destroying control plane blackbox-exporter"]
    DeletingPlutonoInSeed["Deleting Plutono in Seed"]
    DeletingIstioBasicAuthServerInSeed["Deleting istio-basic-auth-server in Seed"]
    DeletingLoggingStackInSeed["Deleting logging stack in Seed"]
    DestroyingInternalDomainDNSRecord{{"Destroying internal domain DNS record\n[CONDITIONAL]"}}:::conditional
    DeletingReferencedResources["Deleting referenced resources"]
    DestroyingMainAndEventsEtcd["Destroying main and events etcd"]
    WaitingUntilMainAndEventEtcdHaveBeenDestroyed["Waiting until main and event etcd have been destroyed"]
    DeletingShootNamespaceInSeed["Deleting shoot namespace in Seed"]
    WaitingUntilShootNamespaceInSeedHasBeenDeleted["Waiting until shoot namespace in Seed has been deleted"]
    DeletingShootState["Deleting Shoot State"]
    SyncPointReadyForCleanup(["Sync: Ready For Cleanup"]):::syncpoint
    SyncPointCleanedKubernetesResources(["Sync: Cleaned Kubernetes Resources"]):::syncpoint
    SyncPointCleaned(["Sync: Cleaned"]):::syncpoint
    SyncPointControlPlaneDown(["Sync: Control Plane Down"]):::syncpoint
    SyncPointObservabilityDown(["Sync: Observability Down"]):::syncpoint

    InitializingConnectionToShoot --> SyncPointReadyForCleanup
    CleaningExtendedAPIGroups --> SyncPointReadyForCleanup
    DeployingShootControlPlane --> SyncPointReadyForCleanup
    DeployingKubernetesControllerManager --> SyncPointReadyForCleanup
    WaitingUntilKubeControllerManagerIsActive --> SyncPointReadyForCleanup
    CleaningUpWebhooks --> SyncPointCleanedKubernetesResources
    CleaningExtendedAPIGroups --> SyncPointCleanedKubernetesResources
    CleaningKubernetesResources --> SyncPointCleanedKubernetesResources
    DeletingMetricsServer --> SyncPointCleanedKubernetesResources
    SyncPointCleanedKubernetesResources --> SyncPointCleaned
    DeletingOperatingSystemConfigResources --> SyncPointCleaned
    WaitingUntilShootWorkerNodesHaveBeenTerminated --> SyncPointCleaned
    WaitingUntilManagedResourcesHaveBeenDeleted --> SyncPointCleaned
    DestroyingShootNetworkPlugin --> SyncPointCleaned
    WaitingUntilShootNetworkPluginHasBeenDestroyed --> SyncPointCleaned
    WaitingUntilExtensionResourcesThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted --> SyncPointCleaned
    WaitingUntilAllStaleExtensionResourcesHaveBeenDeleted --> SyncPointCleaned
    WaitingUntilStaleContainerRuntimeResourcesAreDeleted --> SyncPointCleaned
    WaitingUntilKubernetesAPIServerHasBeenDeleted --> SyncPointControlPlaneDown
    WaitingUntilShootControlPlaneHasBeenDestroyed --> SyncPointControlPlaneDown
    WaitingUntilExtensionResourcesThatShouldBeHandledAfterKubeApiserverHaveBeenDeleted --> SyncPointControlPlaneDown
    WaitingUntilAllExtensionResourcesHaveBeenDeleted --> SyncPointControlPlaneDown
    DestroyingNginxIngressDNSRecord --> SyncPointControlPlaneDown
    DestroyingExternalDomainDNSRecord --> SyncPointControlPlaneDown
    WaitingUntilShootInfrastructureHasBeenDeleted --> SyncPointControlPlaneDown
    DeletingShootAlertmanager --> SyncPointObservabilityDown
    DeletingShootPrometheus --> SyncPointObservabilityDown
    DestroyingControlPlaneBlackboxExporter --> SyncPointObservabilityDown
    DeletingPlutonoInSeed --> SyncPointObservabilityDown
    DeletingIstioBasicAuthServerInSeed --> SyncPointObservabilityDown
    DeletingLoggingStackInSeed --> SyncPointObservabilityDown
    DeployingShootNamespaceInSeed --> EnsuringShootClusterIdentity
    DeployingShootNamespaceInSeed --> DeployingCloudProviderAccountSecret
    DeployingShootNamespaceInSeed --> DeployingKubernetesAPIServerServiceInTheSeedCluster
    EnsuringShootClusterIdentity --> DeployingKubernetesAPIServerServiceInTheSeedCluster
    DeployingKubernetesAPIServerServiceInTheSeedCluster --> WaitingUntilKubernetesAPILoadBalancerInTheSeedClusterHasReportedReadiness
    DeployingShootNamespaceInSeed --> ReconcileIstioInternalLoadBalancingConfigMap
    DeployingShootNamespaceInSeed --> InitializingSecretsManagement
    ReconcileIstioInternalLoadBalancingConfigMap --> InitializingSecretsManagement
    InitializingSecretsManagement --> EnsuringAdvertisedAddressesForTheShoot
    WaitingUntilKubernetesAPILoadBalancerInTheSeedClusterHasReportedReadiness --> EnsuringAdvertisedAddressesForTheShoot
    DeployingShootNamespaceInSeed --> DeployingReferencedResources
    DeployingReferencedResources --> DeployingInternalDomainDNSRecord
    WaitingUntilKubernetesAPILoadBalancerInTheSeedClusterHasReportedReadiness --> DeployingInternalDomainDNSRecord
    InitializingSecretsManagement --> DeployingMainAndEventsEtcd
    DeployingCloudProviderAccountSecret --> DeployingMainAndEventsEtcd
    DeployingMainAndEventsEtcd --> ScalingUpEtcdMainAndEvent
    ScalingUpEtcdMainAndEvent --> WaitingUntilMainAndEventEtcdReportReadiness
    InitializingSecretsManagement --> DeployingShootControlPlane
    DeployingCloudProviderAccountSecret --> DeployingShootControlPlane
    EnsuringShootClusterIdentity --> DeployingShootControlPlane
    DeployingShootControlPlane --> WaitingUntilShootControlPlaneHasBeenReconciled
    InitializingSecretsManagement --> DeployingKubernetesAPIServer
    DeployingMainAndEventsEtcd --> DeployingKubernetesAPIServer
    WaitingUntilMainAndEventEtcdReportReadiness --> DeployingKubernetesAPIServer
    WaitingUntilKubernetesAPILoadBalancerInTheSeedClusterHasReportedReadiness --> DeployingKubernetesAPIServer
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingKubernetesAPIServer
    DeployingKubernetesAPIServer --> ScalingUpKubernetesAPIServer
    DeployingKubernetesAPIServer --> WaitingUntilKubernetesAPIServerReportsReadiness
    ScalingUpKubernetesAPIServer --> WaitingUntilKubernetesAPIServerReportsReadiness
    WaitingUntilKubernetesAPIServerReportsReadiness --> DeployingKubernetesAPIServerServiceSNISettingsInTheSeedCluster
    WaitingUntilKubernetesAPIServerReportsReadiness --> SettingGardenerResourceManagerReplicasTo2
    SettingGardenerResourceManagerReplicasTo2 --> DeployingGardenerResourceManager
    DeployingGardenerResourceManager --> WaitingUntilGardenerResourceManagerReportsReadiness
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingKubernetesAPIServerWithNodeAgentAuthorizer
    DeployingKubernetesAPIServerWithNodeAgentAuthorizer --> WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut
    InitializingSecretsManagement --> DeployingGardenerShootAccessResources
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingGardenerShootAccessResources
    WaitingUntilKubernetesAPIServerWithNodeAgentAuthorizerRolledOut --> DeployingGardenerShootAccessResources
    DeployingCloudProviderAccountSecret --> InitializingConnectionToShoot
    WaitingUntilKubernetesAPIServerReportsReadiness --> InitializingConnectionToShoot
    DeployingInternalDomainDNSRecord --> InitializingConnectionToShoot
    DeployingGardenerShootAccessResources --> InitializingConnectionToShoot
    InitializingSecretsManagement --> DeployingKubernetesControllerManager
    DeployingCloudProviderAccountSecret --> DeployingKubernetesControllerManager
    WaitingUntilShootControlPlaneHasBeenReconciled --> DeployingKubernetesControllerManager
    InitializingConnectionToShoot --> DeployingKubernetesControllerManager
    DeployingKubernetesControllerManager --> ScalingUpKubernetesControllerManager
    InitializingConnectionToShoot --> DeletingClusterAutoscaler
    InitializingConnectionToShoot --> CleaningUpWebhooks
    DeployingGardenerResourceManager --> CleaningUpWebhooks
    InitializingConnectionToShoot --> WaitingUntilKubeControllerManagerIsActive
    CleaningUpWebhooks --> WaitingUntilKubeControllerManagerIsActive
    DeployingShootControlPlane --> WaitingUntilKubeControllerManagerIsActive
    DeployingKubernetesControllerManager --> WaitingUntilKubeControllerManagerIsActive
    InitializingConnectionToShoot --> CleaningExtendedAPIGroups
    DeletingClusterAutoscaler --> CleaningExtendedAPIGroups
    WaitingUntilKubeControllerManagerIsActive --> CleaningExtendedAPIGroups
    SyncPointReadyForCleanup --> CleaningKubernetesResources
    SyncPointReadyForCleanup --> DeletingMetricsServer
    SyncPointCleanedKubernetesResources --> DestroyingShootNetworkPlugin
    DestroyingShootNetworkPlugin --> WaitingUntilShootNetworkPluginHasBeenDestroyed
    SyncPointCleanedKubernetesResources --> DeployingMachineControllerManager
    DeployingMachineControllerManager --> DestroyingShootWorkers
    DestroyingShootWorkers --> WaitingUntilShootWorkerNodesHaveBeenTerminated
    WaitingUntilShootWorkerNodesHaveBeenTerminated --> DeletingMachineControllerManager
    WaitingUntilShootWorkerNodesHaveBeenTerminated --> DeletingOperatingSystemConfigResources
    DeletingOperatingSystemConfigResources --> WaitingUntilAllOperatingSystemConfigResourcesAreDeleted
    SyncPointCleanedKubernetesResources --> DeletingManagedResources
    WaitingUntilShootWorkerNodesHaveBeenTerminated --> DeletingManagedResources
    DeletingManagedResources --> DeletingDWDManagedResourceAndSecrets
    DeletingDWDManagedResourceAndSecrets --> WaitingUntilManagedResourcesHaveBeenDeleted
    CleaningKubernetesResources --> DeletingExtensionResourcesBeforeKubeApiserver
    WaitingUntilAllOperatingSystemConfigResourcesAreDeleted --> DeletingExtensionResourcesBeforeKubeApiserver
    WaitingUntilManagedResourcesHaveBeenDeleted --> DeletingExtensionResourcesBeforeKubeApiserver
    DeletingExtensionResourcesBeforeKubeApiserver --> WaitingUntilExtensionResourcesThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted
    CleaningKubernetesResources --> DeletingStaleExtensionResources
    WaitingUntilManagedResourcesHaveBeenDeleted --> DeletingStaleExtensionResources
    DeletingStaleExtensionResources --> WaitingUntilAllStaleExtensionResourcesHaveBeenDeleted
    InitializingConnectionToShoot --> DeletingContainerRuntimeResources
    SyncPointCleanedKubernetesResources --> DeletingContainerRuntimeResources
    DeletingContainerRuntimeResources --> WaitingUntilStaleContainerRuntimeResourcesAreDeleted
    SyncPointCleaned --> DestroyingShootControlPlane
    DestroyingShootControlPlane --> WaitingUntilShootControlPlaneHasBeenDestroyed
    WaitingUntilShootControlPlaneHasBeenDestroyed --> WaitingUntilShootManagedResourcesHaveBeenDeleted
    SyncPointCleaned --> DeletingKubernetesAPIServer
    WaitingUntilShootControlPlaneHasBeenDestroyed --> DeletingKubernetesAPIServer
    WaitingUntilShootManagedResourcesHaveBeenDeleted --> DeletingKubernetesAPIServer
    DeletingKubernetesAPIServer --> WaitingUntilKubernetesAPIServerHasBeenDeleted
    WaitingUntilKubernetesAPIServerHasBeenDeleted --> DeletingExtensionResourcesAfterKubeApiserver
    DeletingExtensionResourcesAfterKubeApiserver --> WaitingUntilExtensionResourcesThatShouldBeHandledAfterKubeApiserverHaveBeenDeleted
    WaitingUntilKubernetesAPIServerHasBeenDeleted --> WaitingUntilAllExtensionResourcesHaveBeenDeleted
    WaitingUntilKubernetesAPIServerHasBeenDeleted --> DestroyingKubernetesAPIServerServiceSNI
    WaitingUntilKubernetesAPIServerHasBeenDeleted --> DestroyingKubernetesAPIServerService
    DestroyingKubernetesAPIServerServiceSNI --> DestroyingKubernetesAPIServerService
    WaitingUntilKubernetesAPIServerHasBeenDeleted --> DestroyingGardenerResourceManager
    WaitingUntilKubernetesAPIServerHasBeenDeleted --> DeletePublicServiceAccountSigningKeysFromGardenCluster
    SyncPointCleaned --> DestroyingNginxIngressDNSRecord
    SyncPointCleaned --> DestroyingShootInfrastructure
    WaitingUntilShootControlPlaneHasBeenDestroyed --> DestroyingShootInfrastructure
    DestroyingShootInfrastructure --> WaitingUntilShootInfrastructureHasBeenDeleted
    SyncPointCleaned --> DestroyingExternalDomainDNSRecord
    WaitingUntilKubernetesAPIServerHasBeenDeleted --> DestroyingExternalDomainDNSRecord
    SyncPointControlPlaneDown --> DeletingShootAlertmanager
    SyncPointControlPlaneDown --> DeletingShootPrometheus
    SyncPointControlPlaneDown --> DestroyingControlPlaneBlackboxExporter
    SyncPointControlPlaneDown --> DeletingPlutonoInSeed
    DeletingShootAlertmanager --> DeletingIstioBasicAuthServerInSeed
    DeletingShootPrometheus --> DeletingIstioBasicAuthServerInSeed
    DeletingPlutonoInSeed --> DeletingIstioBasicAuthServerInSeed
    SyncPointControlPlaneDown --> DeletingLoggingStackInSeed
    SyncPointObservabilityDown --> DestroyingInternalDomainDNSRecord
    SyncPointObservabilityDown --> DeletingReferencedResources
    SyncPointObservabilityDown --> DestroyingMainAndEventsEtcd
    SyncPointObservabilityDown --> WaitingUntilMainAndEventEtcdHaveBeenDestroyed
    DestroyingMainAndEventsEtcd --> WaitingUntilMainAndEventEtcdHaveBeenDestroyed
    SyncPointObservabilityDown --> DeletingShootNamespaceInSeed
    DestroyingInternalDomainDNSRecord --> DeletingShootNamespaceInSeed
    DeletingReferencedResources --> DeletingShootNamespaceInSeed
    WaitingUntilMainAndEventEtcdHaveBeenDestroyed --> DeletingShootNamespaceInSeed
    DeletingShootNamespaceInSeed --> WaitingUntilShootNamespaceInSeedHasBeenDeleted
    DeletingShootNamespaceInSeed --> DeletingShootState
```
