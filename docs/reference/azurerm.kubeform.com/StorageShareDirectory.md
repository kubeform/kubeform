---
title: StorageShareDirectory
menu:
  docs_{{ .version }}:
    identifier: storagesharedirectory-azurerm.kubeform.com
    name: StorageShareDirectory
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageShareDirectory
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageShareDirectory` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageShareDirectorySpec](#storagesharedirectoryspec)***||
| `status` | ***[StorageShareDirectoryStatus](#storagesharedirectorystatus)***||
## StorageShareDirectorySpec

Appears on:[StorageShareDirectory](#storagesharedirectory), [StorageShareDirectoryStatus](#storagesharedirectorystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `shareName` | ***string***||
| `storageAccountName` | ***string***||
## StorageShareDirectoryStatus

Appears on:[StorageShareDirectory](#storagesharedirectory)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageShareDirectorySpec](#storagesharedirectoryspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
