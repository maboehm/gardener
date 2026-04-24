<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# join

```mermaid
---
title: join (run)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    EnsuringShootIsNotConcurrentlyReconciledByGardenletWhenJoiningControlPlaneNode{{"Ensuring shoot is not concurrently reconciled by gardenlet when joining control plane node\n[CONDITIONAL]"}}:::conditional
    DetermineZoneConfiguration["Determine zone configuration"]
    DeterminingGardenerNodeAgentSecretContainingTheConfigurationForThisNode["Determining gardener-node-agent Secret containing the configuration for this node"]
    GeneratingETCDCertificates{{"Generating ETCD certificates\n[CONDITIONAL]"}}:::conditional
    WritingETCDFilesToDisk{{"Writing ETCD files to disk\n[CONDITIONAL]"}}:::conditional
    PreparingGardenerNodeInitConfiguration{{"Preparing gardener-node-init configuration\n[CONDITIONAL]"}}:::conditional
    ApplyingOperatingSystemConfigUsingGardenerNodeAgentsReconciliationLogic{{"Applying OperatingSystemConfig using gardener-node-agent's reconciliation logic\n[CONDITIONAL]"}}:::conditional
    WaitingForNodeToJoinTheClusterAndBecomeReady["Waiting for node to join the cluster and become ready"]
    SyncPointReadyForGardenerNodeInit(["Sync: Ready For Gardener Node Init"]):::syncpoint

    DeterminingGardenerNodeAgentSecretContainingTheConfigurationForThisNode --> SyncPointReadyForGardenerNodeInit
    EnsuringShootIsNotConcurrentlyReconciledByGardenletWhenJoiningControlPlaneNode --> SyncPointReadyForGardenerNodeInit
    DetermineZoneConfiguration --> SyncPointReadyForGardenerNodeInit
    WritingETCDFilesToDisk --> SyncPointReadyForGardenerNodeInit
    GeneratingETCDCertificates --> WritingETCDFilesToDisk
    SyncPointReadyForGardenerNodeInit --> PreparingGardenerNodeInitConfiguration
    PreparingGardenerNodeInitConfiguration --> ApplyingOperatingSystemConfigUsingGardenerNodeAgentsReconciliationLogic
    ApplyingOperatingSystemConfigUsingGardenerNodeAgentsReconciliationLogic --> WaitingForNodeToJoinTheClusterAndBecomeReady
```
