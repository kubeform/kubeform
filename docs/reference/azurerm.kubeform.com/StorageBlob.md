---
title: StorageBlob
menu:
  docs_{{ .version }}:
    identifier: storageblob-azurerm.kubeform.com
    name: StorageBlob
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageBlob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageBlob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageBlobSpec](#StorageBlobSpec)***||
| `status` | ***[StorageBlobStatus](#StorageBlobStatus)***||
## StorageBlobSpec
##### (Appears on:[StorageBlob](#StorageBlob), [StorageBlobStatus](#StorageBlobStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attempts` | ***int***| ***(Optional)*** |
| `contentType` | ***string***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `parallelism` | ***int***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `size` | ***int***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
| `sourceURI` | ***string***| ***(Optional)*** |
| `storageAccountName` | ***string***||
| `storageContainerName` | ***string***||
| `type` | ***string***| ***(Optional)*** |
| `url` | ***string***| ***(Optional)*** |
## StorageBlobStatus
##### (Appears on:[StorageBlob](#StorageBlob))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageBlobSpec](#StorageBlobSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
