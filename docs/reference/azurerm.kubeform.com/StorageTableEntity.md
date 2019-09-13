---
title: StorageTableEntity
menu:
  docs_{{ .version }}:
    identifier: storagetableentity-azurerm.kubeform.com
    name: StorageTableEntity
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageTableEntity
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageTableEntity` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageTableEntitySpec](#StorageTableEntitySpec)***||
| `status` | ***[StorageTableEntityStatus](#StorageTableEntityStatus)***||
## StorageTableEntitySpec
##### (Appears on:[StorageTableEntity](#StorageTableEntity), [StorageTableEntityStatus](#StorageTableEntityStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `entity` | ***map[string]string***||
| `partitionKey` | ***string***||
| `rowKey` | ***string***||
| `storageAccountName` | ***string***||
| `tableName` | ***string***||
## StorageTableEntityStatus
##### (Appears on:[StorageTableEntity](#StorageTableEntity))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageTableEntitySpec](#StorageTableEntitySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
