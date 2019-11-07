---
title: StorageQueue
menu:
  docs_{{ .version }}:
    identifier: storagequeue-azurerm.kubeform.com
    name: StorageQueue
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageQueue
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageQueue` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageQueueSpec](#storagequeuespec)***||
| `status` | ***[StorageQueueStatus](#storagequeuestatus)***||
## Phase(`string` alias)

Appears on:[StorageQueueStatus](#storagequeuestatus)

## StorageQueueSpec

Appears on:[StorageQueue](#storagequeue), [StorageQueueStatus](#storagequeuestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***| ***(Optional)*** Deprecated|
| `storageAccountName` | ***string***||
## StorageQueueStatus

Appears on:[StorageQueue](#storagequeue)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageQueueSpec](#storagequeuespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---