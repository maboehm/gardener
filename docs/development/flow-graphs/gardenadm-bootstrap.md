<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# bootstrap

```mermaid
---
title: bootstrap (run)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    DeployingControlPlaneNamespace["Deploying control plane namespace"]
    DeployingCloudProviderAccountSecret{{"Deploying cloud provider account secret\n[CONDITIONAL]"}}:::conditional
    ReconcilingCustomResourceDefinitions["Reconciling CustomResourceDefinitions"]
    EnsuringCustomResourceDefinitionsAreReady["Ensuring CustomResourceDefinitions are ready"]
    ReconcilingExtensionsgardenercloudV1alpha1ClusterResource["Reconciling extensions.gardener.cloud/v1alpha1.Cluster resource"]
    InitializingInternalStateOfGardenerSecretsManager["Initializing internal state of Gardener secrets manager"]
    DeployingPriorityClassForGardenerResourceManager["Deploying PriorityClass for gardener-resource-manager"]
    DeployingGardenerResourceManager["Deploying gardener-resource-manager"]
    WaitingUntilGardenerResourceManagerReportsReadiness["Waiting until gardener-resource-manager reports readiness"]
    DeployingSeedSystemResources["Deploying seed system resources"]
    DeployingExtensionControllers["Deploying extension controllers"]
    WaitingUntilExtensionControllersReportReadiness["Waiting until extension controllers report readiness"]
    DeployingNetworkPolicies["Deploying network policies"]
    DeployingShootInfrastructure{{"Deploying Shoot infrastructure\n[CONDITIONAL]"}}:::conditional
    WaitingUntilShootInfrastructureHasBeenReconciled{{"Waiting until Shoot infrastructure has been reconciled\n[CONDITIONAL]"}}:::conditional
    DeployingOperatingSystemConfigForControlPlaneMachines["Deploying OperatingSystemConfig for control plane machines"]
    WaitingUntilOperatingSystemConfigForControlPlaneMachinesHasBeenReconciled["Waiting until OperatingSystemConfig for control plane machines has been reconciled"]
    DeployingMachineControllerManager["Deploying machine-controller-manager"]
    DeployingControlPlaneMachines{{"Deploying control plane machines\n[CONDITIONAL]"}}:::conditional
    WaitingUntilControlPlaneMachinesHaveBeenDeployed{{"Waiting until control plane machines have been deployed\n[CONDITIONAL]"}}:::conditional
    ListingControlPlaneMachines["Listing control plane machines"]
    ScalingDownMachineControllerManager["Scaling down machine-controller-manager"]
    DeployingDNSRecordPointingToTheFirstControlPlaneMachine{{"Deploying DNSRecord pointing to the first control plane machine\n[CONDITIONAL]"}}:::conditional
    PreparingExtensionResourcesForMigrationToSelfHostedShoot["Preparing extension resources for migration to self-hosted shoot"]
    CompilingShootState["Compiling ShootState"]
    DeployingAndConnectingToBastionHost["Deploying and connecting to bastion host"]
    ConnectingToTheFirstControlPlaneMachine["Connecting to the first control plane machine"]
    CopyingManifestsToTheFirstControlPlaneMachine["Copying manifests to the first control plane machine"]
    BootstrappingControlPlaneOnTheFirstControlPlaneMachine["Bootstrapping control plane on the first control plane machine"]
    FetchingKubeconfigFromTheFirstControlPlaneMachine["Fetching kubeconfig from the first control plane machine"]
    SyncPointBootstrapped(["Sync: Bootstrapped"]):::syncpoint

    DeployingNetworkPolicies --> SyncPointBootstrapped
    WaitingUntilGardenerResourceManagerReportsReadiness --> SyncPointBootstrapped
    WaitingUntilExtensionControllersReportReadiness --> SyncPointBootstrapped
    DeployingControlPlaneNamespace --> DeployingCloudProviderAccountSecret
    ReconcilingCustomResourceDefinitions --> EnsuringCustomResourceDefinitionsAreReady
    EnsuringCustomResourceDefinitionsAreReady --> ReconcilingExtensionsgardenercloudV1alpha1ClusterResource
    ReconcilingExtensionsgardenercloudV1alpha1ClusterResource --> InitializingInternalStateOfGardenerSecretsManager
    DeployingControlPlaneNamespace --> DeployingPriorityClassForGardenerResourceManager
    InitializingInternalStateOfGardenerSecretsManager --> DeployingPriorityClassForGardenerResourceManager
    DeployingControlPlaneNamespace --> DeployingGardenerResourceManager
    InitializingInternalStateOfGardenerSecretsManager --> DeployingGardenerResourceManager
    DeployingPriorityClassForGardenerResourceManager --> DeployingGardenerResourceManager
    DeployingGardenerResourceManager --> WaitingUntilGardenerResourceManagerReportsReadiness
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingSeedSystemResources
    WaitingUntilGardenerResourceManagerReportsReadiness --> DeployingExtensionControllers
    DeployingExtensionControllers --> WaitingUntilExtensionControllersReportReadiness
    DeployingGardenerResourceManager --> DeployingNetworkPolicies
    DeployingExtensionControllers --> DeployingNetworkPolicies
    SyncPointBootstrapped --> DeployingShootInfrastructure
    DeployingShootInfrastructure --> WaitingUntilShootInfrastructureHasBeenReconciled
    SyncPointBootstrapped --> DeployingOperatingSystemConfigForControlPlaneMachines
    DeployingOperatingSystemConfigForControlPlaneMachines --> WaitingUntilOperatingSystemConfigForControlPlaneMachinesHasBeenReconciled
    SyncPointBootstrapped --> DeployingMachineControllerManager
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeployingControlPlaneMachines
    WaitingUntilOperatingSystemConfigForControlPlaneMachinesHasBeenReconciled --> DeployingControlPlaneMachines
    DeployingMachineControllerManager --> DeployingControlPlaneMachines
    DeployingControlPlaneMachines --> WaitingUntilControlPlaneMachinesHaveBeenDeployed
    WaitingUntilControlPlaneMachinesHaveBeenDeployed --> ListingControlPlaneMachines
    WaitingUntilControlPlaneMachinesHaveBeenDeployed --> ScalingDownMachineControllerManager
    ListingControlPlaneMachines --> DeployingDNSRecordPointingToTheFirstControlPlaneMachine
    ScalingDownMachineControllerManager --> PreparingExtensionResourcesForMigrationToSelfHostedShoot
    DeployingDNSRecordPointingToTheFirstControlPlaneMachine --> PreparingExtensionResourcesForMigrationToSelfHostedShoot
    PreparingExtensionResourcesForMigrationToSelfHostedShoot --> CompilingShootState
    WaitingUntilShootInfrastructureHasBeenReconciled --> DeployingAndConnectingToBastionHost
    ListingControlPlaneMachines --> ConnectingToTheFirstControlPlaneMachine
    DeployingAndConnectingToBastionHost --> ConnectingToTheFirstControlPlaneMachine
    ConnectingToTheFirstControlPlaneMachine --> CopyingManifestsToTheFirstControlPlaneMachine
    CompilingShootState --> CopyingManifestsToTheFirstControlPlaneMachine
    DeployingDNSRecordPointingToTheFirstControlPlaneMachine --> BootstrappingControlPlaneOnTheFirstControlPlaneMachine
    CopyingManifestsToTheFirstControlPlaneMachine --> BootstrappingControlPlaneOnTheFirstControlPlaneMachine
    BootstrappingControlPlaneOnTheFirstControlPlaneMachine --> FetchingKubeconfigFromTheFirstControlPlaneMachine
```
