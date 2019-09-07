---
title: DataLakeStoreFile
menu:
  docs_v0.0.1:
    identifier: datalakestorefile-azurerm.kubeform.com
    name: DataLakeStoreFile
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DataLakeStoreFile
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataLakeStoreFile` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataLakeStoreFileSpec](#DataLakeStoreFileSpec)***||
| `status` | ***[DataLakeStoreFileStatus](#DataLakeStoreFileStatus)***||
## DataLakeStoreFileSpec
##### (Appears on:[DataLakeStoreFile](#DataLakeStoreFile), [DataLakeStoreFileStatus](#DataLakeStoreFileStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `localFilePath` | ***string***||
| `remoteFilePath` | ***string***||
## DataLakeStoreFileStatus
##### (Appears on:[DataLakeStoreFile](#DataLakeStoreFile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataLakeStoreFileSpec](#DataLakeStoreFileSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
