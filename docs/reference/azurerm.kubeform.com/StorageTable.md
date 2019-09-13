---
title: StorageTable
menu:
  docs_{{ .version }}:
    identifier: storagetable-azurerm.kubeform.com
    name: StorageTable
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageTable
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageTable` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageTableSpec](#StorageTableSpec)***||
| `status` | ***[StorageTableStatus](#StorageTableStatus)***||
## StorageTableSpec
##### (Appears on:[StorageTable](#StorageTable), [StorageTableStatus](#StorageTableStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acl` | ***[[]StorageTableSpecAcl](#StorageTableSpecAcl)***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***| ***(Optional)*** Deprecated|
| `storageAccountName` | ***string***||
## StorageTableSpecAcl
##### (Appears on:[StorageTableSpec](#StorageTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessPolicy` | ***[[]StorageTableSpecAclAccessPolicy](#StorageTableSpecAclAccessPolicy)***| ***(Optional)*** |
| `ID` | ***string***||
## StorageTableSpecAclAccessPolicy
##### (Appears on:[StorageTableSpecAcl](#StorageTableSpecAcl))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `expiry` | ***string***||
| `permissions` | ***string***||
| `start` | ***string***||
## StorageTableStatus
##### (Appears on:[StorageTable](#StorageTable))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageTableSpec](#StorageTableSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
