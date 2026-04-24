<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# bootstrap

```mermaid
---
title: bootstrap (bootstrapControlPlane)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    InitializingSecretsManagement{{"Initializing secrets management\n[CONDITIONAL]"}}:::conditional
    WritingKubeletBootstrapKubeconfigWithAFakeTokenToDiskToMakeKubeletStart{{"Writing kubelet bootstrap kubeconfig with a fake token to disk to make kubelet start\n[CONDITIONAL]"}}:::conditional
    GeneratingOperatingSystemConfigAndDeployingSecretForGardenerNodeAgent{{"Generating OperatingSystemConfig and deploying Secret for gardener-node-agent\n[CONDITIONAL]"}}:::conditional
    PersistingBootstrapSecretsAsShootStateForRetryResilience{{"Persisting bootstrap secrets as ShootState for retry resilience\n[CONDITIONAL]"}}:::conditional
    ApplyingOperatingSystemConfigUsingGardenerNodeAgentsReconciliationLogic{{"Applying OperatingSystemConfig using gardener-node-agent's reconciliation logic\n[CONDITIONAL]"}}:::conditional
    InitializingConnectionToKubernetesControlPlane["Initializing connection to Kubernetes control plane"]
    ImportingSecretsIntoControlPlane{{"Importing secrets into control plane\n[CONDITIONAL]"}}:::conditional

    InitializingSecretsManagement --> WritingKubeletBootstrapKubeconfigWithAFakeTokenToDiskToMakeKubeletStart
    InitializingSecretsManagement --> GeneratingOperatingSystemConfigAndDeployingSecretForGardenerNodeAgent
    GeneratingOperatingSystemConfigAndDeployingSecretForGardenerNodeAgent --> PersistingBootstrapSecretsAsShootStateForRetryResilience
    WritingKubeletBootstrapKubeconfigWithAFakeTokenToDiskToMakeKubeletStart --> ApplyingOperatingSystemConfigUsingGardenerNodeAgentsReconciliationLogic
    PersistingBootstrapSecretsAsShootStateForRetryResilience --> ApplyingOperatingSystemConfigUsingGardenerNodeAgentsReconciliationLogic
    ApplyingOperatingSystemConfigUsingGardenerNodeAgentsReconciliationLogic --> InitializingConnectionToKubernetesControlPlane
    PersistingBootstrapSecretsAsShootStateForRetryResilience --> ImportingSecretsIntoControlPlane
    InitializingConnectionToKubernetesControlPlane --> ImportingSecretsIntoControlPlane
```
