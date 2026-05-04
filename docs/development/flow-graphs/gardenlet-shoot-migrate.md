<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# Shoot cluster preparation for migration

```mermaid
---
title: Shoot cluster preparation for migration (runMigrateShootFlow)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    DeployingShootNamespaceInSeed{{"Deploying Shoot namespace in Seed\n[CONDITIONAL]"}}:::conditional
    ReconcileIstioInternalLoadBalancingConfigMap{{"Reconcile Istio internal load balancing ConfigMap\n[CONDITIONAL]"}}:::conditional
    InitializingSecretsManagement{{"Initializing secrets management\n[CONDITIONAL]"}}:::conditional
    DeployingMainAndEventsEtcd{{"Deploying main and events etcd\n[CONDITIONAL]"}}:::conditional
    ScalingEtcdUp{{"Scaling etcd up\n[CONDITIONAL]"}}:::conditional
    WaitingUntilMainAndEventEtcdReportReadiness{{"Waiting until main and event etcd report readiness\n[CONDITIONAL]"}}:::conditional
    ScalingKubernetesAPIServerUpAndWaitingUntilReady{{"Scaling Kubernetes API Server up and waiting until ready\n[CONDITIONAL]"}}:::conditional
    DeployingGardenerResourceManager{{"Deploying gardener-resource-manager\n[CONDITIONAL]"}}:::conditional
    EnsuringThatTheGardenerResourceManagerIsScaledTo1{{"Ensuring that the gardener-resource-manager is scaled to 1\n[CONDITIONAL]"}}:::conditional
    ConfiguringManagedResourcesObjectsToBeKeptInTheShoot{{"Configuring Managed Resources objects to be kept in the Shoot\n[CONDITIONAL]"}}:::conditional
    DeletingAllManagedResourcesFromTheShootsNamespace["Deleting all Managed Resources from the Shoot's namespace"]
    WaitingUntilManagedResourcesAreDeleted["Waiting until ManagedResources are deleted"]
    DeletingMachineControllerManager{{"Deleting machine-controller-manager\n[CONDITIONAL]"}}:::conditional
    WaitingUntilMachineControllerManagerHasBeenDeleted{{"Waiting until machine-controller-manager has been deleted\n[CONDITIONAL]"}}:::conditional
    MigratingExtensionResources["Migrating extension resources"]
    WaitingUntilExtensionResourcesHaveBeenMigrated["Waiting until extension resources have been migrated"]
    MigratingExtensionsBeforeKubeApiserver["Migrating extensions before kube-apiserver"]
    WaitingUntilExtensionsThatShouldBeHandledBeforeKubeApiserverHaveBeenMigrated["Waiting until extensions that should be handled before kube-apiserver have been migrated"]
    PersistingShootStateInGardenCluster["Persisting ShootState in garden cluster"]
    DeletingExtensionResourcesFromTheShootNamespace["Deleting extension resources from the Shoot namespace"]
    WaitingUntilExtensionResourcesHaveBeenDeleted["Waiting until extension resources have been deleted"]
    ShallowDeletingMachineResourcesFromTheShootNamespace["Shallow-deleting machine resources from the Shoot namespace"]
    WaitingUntilMachineResourcesHaveBeenDeleted["Waiting until machine resources have been deleted"]
    DeletingExtensionsBeforeKubeApiserver["Deleting extensions before kube-apiserver"]
    WaitingUntilExtensionsThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted["Waiting until extensions that should be handled before kube-apiserver have been deleted"]
    DeletingStaleExtensions["Deleting stale extensions"]
    WaitingUntilAllStaleExtensionsHaveBeenDeleted["Waiting until all stale extensions have been deleted"]
    MigratingShootControlPlane{{"Migrating shoot control plane\n[CONDITIONAL]"}}:::conditional
    DeletingShootControlPlane{{"Deleting shoot control plane\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootControlPlaneHasBeenDeleted{{"Waiting until shoot control plane has been deleted\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootManagedResourcesHaveBeenDeleted{{"Waiting until shoot managed resources have been deleted\n[CONDITIONAL]"}}:::conditional
    DeletingKubeApiserverDeployment["Deleting kube-apiserver deployment"]
    WaitingUntilKubeApiserverHasBeenDeleted["Waiting until kube-apiserver has been deleted"]
    MigratingExtensionsAfterKubeApiserver["Migrating extensions after kube-apiserver"]
    WaitingUntilExtensionsThatShouldBeHandledAfterKubeApiserverHaveBeenMigrated["Waiting until extensions that should be handled after kube-apiserver have been migrated"]
    DeletingExtensionsAfterKubeApiserver["Deleting extensions after kube-apiserver"]
    WaitingUntilExtensionsThatShouldBeHandledAfterKubeApiserverHaveBeenDeleted["Waiting until extensions that should be handled after kube-apiserver have been deleted"]
    WaitingUntilAllExtensionsHaveBeenDeleted["Waiting until all extensions have been deleted"]
    MigratingShootInfrastructure{{"Migrating shoot infrastructure\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootInfrastructureHasBeenMigrated{{"Waiting until shoot infrastructure has been migrated\n[CONDITIONAL]"}}:::conditional
    DeletingShootInfrastructure{{"Deleting shoot infrastructure\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootInfrastructureHasBeenDeleted{{"Waiting until shoot infrastructure has been deleted\n[CONDITIONAL]"}}:::conditional
    MigratingNginxIngressDNSRecord["Migrating nginx ingress DNS record"]
    MigratingExternalDomainDNSRecord["Migrating external domain DNS record"]
    MigratingInternalDomainDNSRecord["Migrating internal domain DNS record"]
    DeletingDNSRecordsFromTheShootNamespace{{"Deleting DNSRecords from the Shoot namespace\n[CONDITIONAL]"}}:::conditional
    CreatingETCDSnapshot{{"Creating ETCD Snapshot\n[CONDITIONAL]"}}:::conditional
    MigratingBackupEntryToNewSeed["Migrating BackupEntry to new seed"]
    WaitingForBackupEntryToBeMigratedToNewSeed["Waiting for BackupEntry to be migrated to new seed"]
    DestroyingMainAndEventsEtcd["Destroying main and events etcd"]
    WaitingUntilMainAndEventEtcdHaveBeenDestroyed["Waiting until main and event etcd have been destroyed"]
    DeletingShootNamespaceInSeed["Deleting shoot namespace in Seed"]
    WaitingUntilShootNamespaceInSeedHasBeenDeleted["Waiting until shoot namespace in Seed has been deleted"]
    SyncPoint(["Sync: sync Point"]):::syncpoint

    WaitingUntilExtensionsThatShouldBeHandledAfterKubeApiserverHaveBeenDeleted --> SyncPoint
    WaitingUntilMachineResourcesHaveBeenDeleted --> SyncPoint
    WaitingUntilAllExtensionsHaveBeenDeleted --> SyncPoint
    WaitingUntilShootInfrastructureHasBeenDeleted --> SyncPoint
    DeployingShootNamespaceInSeed --> ReconcileIstioInternalLoadBalancingConfigMap
    DeployingShootNamespaceInSeed --> InitializingSecretsManagement
    ReconcileIstioInternalLoadBalancingConfigMap --> InitializingSecretsManagement
    InitializingSecretsManagement --> DeployingMainAndEventsEtcd
    DeployingMainAndEventsEtcd --> ScalingEtcdUp
    DeployingMainAndEventsEtcd --> WaitingUntilMainAndEventEtcdReportReadiness
    ScalingEtcdUp --> WaitingUntilMainAndEventEtcdReportReadiness
    DeployingMainAndEventsEtcd --> ScalingKubernetesAPIServerUpAndWaitingUntilReady
    ScalingEtcdUp --> ScalingKubernetesAPIServerUpAndWaitingUntilReady
    InitializingSecretsManagement --> ScalingKubernetesAPIServerUpAndWaitingUntilReady
    ScalingKubernetesAPIServerUpAndWaitingUntilReady --> DeployingGardenerResourceManager
    DeployingGardenerResourceManager --> EnsuringThatTheGardenerResourceManagerIsScaledTo1
    EnsuringThatTheGardenerResourceManagerIsScaledTo1 --> ConfiguringManagedResourcesObjectsToBeKeptInTheShoot
    ConfiguringManagedResourcesObjectsToBeKeptInTheShoot --> DeletingAllManagedResourcesFromTheShootsNamespace
    EnsuringThatTheGardenerResourceManagerIsScaledTo1 --> DeletingAllManagedResourcesFromTheShootsNamespace
    DeletingAllManagedResourcesFromTheShootsNamespace --> WaitingUntilManagedResourcesAreDeleted
    WaitingUntilManagedResourcesAreDeleted --> DeletingMachineControllerManager
    DeletingMachineControllerManager --> WaitingUntilMachineControllerManagerHasBeenDeleted
    WaitingUntilMachineControllerManagerHasBeenDeleted --> MigratingExtensionResources
    MigratingExtensionResources --> WaitingUntilExtensionResourcesHaveBeenMigrated
    WaitingUntilManagedResourcesAreDeleted --> MigratingExtensionsBeforeKubeApiserver
    MigratingExtensionsBeforeKubeApiserver --> WaitingUntilExtensionsThatShouldBeHandledBeforeKubeApiserverHaveBeenMigrated
    WaitingUntilExtensionResourcesHaveBeenMigrated --> PersistingShootStateInGardenCluster
    PersistingShootStateInGardenCluster --> DeletingExtensionResourcesFromTheShootNamespace
    DeletingExtensionResourcesFromTheShootNamespace --> WaitingUntilExtensionResourcesHaveBeenDeleted
    PersistingShootStateInGardenCluster --> ShallowDeletingMachineResourcesFromTheShootNamespace
    ShallowDeletingMachineResourcesFromTheShootNamespace --> WaitingUntilMachineResourcesHaveBeenDeleted
    WaitingUntilExtensionResourcesHaveBeenDeleted --> DeletingExtensionsBeforeKubeApiserver
    WaitingUntilExtensionsThatShouldBeHandledBeforeKubeApiserverHaveBeenMigrated --> DeletingExtensionsBeforeKubeApiserver
    DeletingExtensionsBeforeKubeApiserver --> WaitingUntilExtensionsThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted
    WaitingUntilExtensionResourcesHaveBeenMigrated --> DeletingStaleExtensions
    DeletingStaleExtensions --> WaitingUntilAllStaleExtensionsHaveBeenDeleted
    WaitingUntilExtensionResourcesHaveBeenDeleted --> MigratingShootControlPlane
    WaitingUntilExtensionsThatShouldBeHandledBeforeKubeApiserverHaveBeenDeleted --> MigratingShootControlPlane
    WaitingUntilAllStaleExtensionsHaveBeenDeleted --> MigratingShootControlPlane
    MigratingShootControlPlane --> DeletingShootControlPlane
    DeletingShootControlPlane --> WaitingUntilShootControlPlaneHasBeenDeleted
    WaitingUntilShootControlPlaneHasBeenDeleted --> WaitingUntilShootManagedResourcesHaveBeenDeleted
    WaitingUntilManagedResourcesAreDeleted --> DeletingKubeApiserverDeployment
    WaitingUntilMainAndEventEtcdReportReadiness --> DeletingKubeApiserverDeployment
    WaitingUntilShootControlPlaneHasBeenDeleted --> DeletingKubeApiserverDeployment
    WaitingUntilShootManagedResourcesHaveBeenDeleted --> DeletingKubeApiserverDeployment
    DeletingKubeApiserverDeployment --> WaitingUntilKubeApiserverHasBeenDeleted
    WaitingUntilKubeApiserverHasBeenDeleted --> MigratingExtensionsAfterKubeApiserver
    MigratingExtensionsAfterKubeApiserver --> WaitingUntilExtensionsThatShouldBeHandledAfterKubeApiserverHaveBeenMigrated
    WaitingUntilExtensionsThatShouldBeHandledAfterKubeApiserverHaveBeenMigrated --> DeletingExtensionsAfterKubeApiserver
    DeletingExtensionsAfterKubeApiserver --> WaitingUntilExtensionsThatShouldBeHandledAfterKubeApiserverHaveBeenDeleted
    WaitingUntilExtensionsThatShouldBeHandledAfterKubeApiserverHaveBeenMigrated --> WaitingUntilAllExtensionsHaveBeenDeleted
    WaitingUntilKubeApiserverHasBeenDeleted --> MigratingShootInfrastructure
    MigratingShootInfrastructure --> WaitingUntilShootInfrastructureHasBeenMigrated
    WaitingUntilShootInfrastructureHasBeenMigrated --> DeletingShootInfrastructure
    DeletingShootInfrastructure --> WaitingUntilShootInfrastructureHasBeenDeleted
    WaitingUntilKubeApiserverHasBeenDeleted --> MigratingNginxIngressDNSRecord
    WaitingUntilKubeApiserverHasBeenDeleted --> MigratingExternalDomainDNSRecord
    WaitingUntilKubeApiserverHasBeenDeleted --> MigratingInternalDomainDNSRecord
    SyncPoint --> DeletingDNSRecordsFromTheShootNamespace
    MigratingNginxIngressDNSRecord --> DeletingDNSRecordsFromTheShootNamespace
    MigratingExternalDomainDNSRecord --> DeletingDNSRecordsFromTheShootNamespace
    MigratingInternalDomainDNSRecord --> DeletingDNSRecordsFromTheShootNamespace
    SyncPoint --> CreatingETCDSnapshot
    WaitingUntilKubeApiserverHasBeenDeleted --> CreatingETCDSnapshot
    SyncPoint --> MigratingBackupEntryToNewSeed
    CreatingETCDSnapshot --> MigratingBackupEntryToNewSeed
    MigratingBackupEntryToNewSeed --> WaitingForBackupEntryToBeMigratedToNewSeed
    SyncPoint --> DestroyingMainAndEventsEtcd
    CreatingETCDSnapshot --> DestroyingMainAndEventsEtcd
    WaitingForBackupEntryToBeMigratedToNewSeed --> DestroyingMainAndEventsEtcd
    DestroyingMainAndEventsEtcd --> WaitingUntilMainAndEventEtcdHaveBeenDestroyed
    SyncPoint --> DeletingShootNamespaceInSeed
    WaitingForBackupEntryToBeMigratedToNewSeed --> DeletingShootNamespaceInSeed
    DeletingExtensionResourcesFromTheShootNamespace --> DeletingShootNamespaceInSeed
    DeletingDNSRecordsFromTheShootNamespace --> DeletingShootNamespaceInSeed
    WaitingUntilManagedResourcesAreDeleted --> DeletingShootNamespaceInSeed
    WaitingUntilMainAndEventEtcdHaveBeenDestroyed --> DeletingShootNamespaceInSeed
    DeletingShootNamespaceInSeed --> WaitingUntilShootNamespaceInSeedHasBeenDeleted
```
