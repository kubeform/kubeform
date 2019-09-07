---
title: DevTestLab
menu:
  docs_v0.0.1:
    identifier: devtestlab-azurerm.kubeform.com
    name: DevTestLab
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DevTestLab
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DevTestLab` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevTestLabSpec](#DevTestLabSpec)***||
| `status` | ***[DevTestLabStatus](#DevTestLabStatus)***||
## DevTestLabSpec
##### (Appears on:[DevTestLab](#DevTestLab), [DevTestLabStatus](#DevTestLabStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `artifactsStorageAccountID` | ***string***| ***(Optional)*** |
| `defaultPremiumStorageAccountID` | ***string***| ***(Optional)*** |
| `defaultStorageAccountID` | ***string***| ***(Optional)*** |
| `keyVaultID` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `premiumDataDiskStorageAccountID` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `storageType` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uniqueIdentifier` | ***string***| ***(Optional)*** |
## DevTestLabStatus
##### (Appears on:[DevTestLab](#DevTestLab))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevTestLabSpec](#DevTestLabSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
