---
title: KeyVaultAccessPolicy
menu:
  docs_{{ .version }}:
    identifier: keyvaultaccesspolicy-azurerm.kubeform.com
    name: KeyVaultAccessPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KeyVaultAccessPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KeyVaultAccessPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KeyVaultAccessPolicySpec](#KeyVaultAccessPolicySpec)***||
| `status` | ***[KeyVaultAccessPolicyStatus](#KeyVaultAccessPolicyStatus)***||
## KeyVaultAccessPolicySpec
##### (Appears on:[KeyVaultAccessPolicy](#KeyVaultAccessPolicy), [KeyVaultAccessPolicyStatus](#KeyVaultAccessPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationID` | ***string***| ***(Optional)*** |
| `certificatePermissions` | ***[]string***| ***(Optional)*** |
| `keyPermissions` | ***[]string***| ***(Optional)*** |
| `keyVaultID` | ***string***| ***(Optional)*** |
| `objectID` | ***string***||
| `resourceGroupName` | ***string***| ***(Optional)*** Deprecated|
| `secretPermissions` | ***[]string***| ***(Optional)*** |
| `storagePermissions` | ***[]string***| ***(Optional)*** |
| `tenantID` | ***string***||
| `vaultName` | ***string***| ***(Optional)*** Deprecated|
## KeyVaultAccessPolicyStatus
##### (Appears on:[KeyVaultAccessPolicy](#KeyVaultAccessPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KeyVaultAccessPolicySpec](#KeyVaultAccessPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
