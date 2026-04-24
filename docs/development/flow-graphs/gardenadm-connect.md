<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# connect

```mermaid
---
title: connect (run)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    RetrievingShortLivedKubeconfigForGardenClusterToPrepareGardenerResources["Retrieving short-lived kubeconfig for garden cluster to prepare Gardener resources"]
    PreparingGardenerResourcesInGardenCluster["Preparing Gardener resources in garden cluster"]
    DeployingGardenletIntoSelfHostedShootCluster["Deploying gardenlet into self-hosted shoot cluster"]
    WaitingUntilGardenletIsReady["Waiting until gardenlet is ready"]

    RetrievingShortLivedKubeconfigForGardenClusterToPrepareGardenerResources --> PreparingGardenerResourcesInGardenCluster
    PreparingGardenerResourcesInGardenCluster --> DeployingGardenletIntoSelfHostedShootCluster
    DeployingGardenletIntoSelfHostedShootCluster --> WaitingUntilGardenletIsReady
```
