---
title: StorageShareDirectory
menu:
  docs_v0.0.1:
    identifier: storagesharedirectory-azurerm.kubeform.com
    name: StorageShareDirectory
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## StorageShareDirectory
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageShareDirectory` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageShareDirectorySpec](#StorageShareDirectorySpec)***||
| `status` | ***[StorageShareDirectoryStatus](#StorageShareDirectoryStatus)***||
## StorageShareDirectorySpec
##### (Appears on:[StorageShareDirectory](#StorageShareDirectory), [StorageShareDirectoryStatus](#StorageShareDirectoryStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `shareName` | ***string***||
| `storageAccountName` | ***string***||
## StorageShareDirectoryStatus
##### (Appears on:[StorageShareDirectory](#StorageShareDirectory))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageShareDirectorySpec](#StorageShareDirectorySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
