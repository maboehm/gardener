<!-- This file is auto-generated via `make generate`. DO NOT EDIT. -->

# Extension reconciliation

```mermaid
---
title: Extension reconciliation (reconcile)
---
flowchart TD
    classDef conditional stroke-dasharray:5 5,color:#888
    classDef syncpoint fill:#ddf4ff,stroke:#4a90d9,color:#1a5276

    DeployingExtensionInRuntimeCluster["Deploying extension in runtime cluster"]
    CheckingIfGardenIsReconciled["Checking if garden is reconciled"]
    CreatingVirtualGardenClient["Creating virtual garden-client"]
    DeployingAdmissionController["Deploying Admission Controller"]
    DeployingControllerRegistrationAndControllerDeployment["Deploying ControllerRegistration and ControllerDeployment"]

    DeployingExtensionInRuntimeCluster --> CheckingIfGardenIsReconciled
    CheckingIfGardenIsReconciled --> CreatingVirtualGardenClient
    CreatingVirtualGardenClient --> DeployingAdmissionController
    CheckingIfGardenIsReconciled --> DeployingControllerRegistrationAndControllerDeployment
```
