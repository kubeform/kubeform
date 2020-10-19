---
title: StorageDataLakeGen2Filesystem
menu:
  docs_{{ .version }}:
    identifier: storagedatalakegen2filesystem-azurerm.kubeform.com
    name: StorageDataLakeGen2Filesystem
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageDataLakeGen2Filesystem
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageDataLakeGen2Filesystem` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageDataLakeGen2FilesystemSpec](#storagedatalakegen2filesystemspec)***||
| `status` | ***[StorageDataLakeGen2FilesystemStatus](#storagedatalakegen2filesystemstatus)***||
## Phase(`string` alias)

Appears on:[StorageDataLakeGen2FilesystemStatus](#storagedatalakegen2filesystemstatus)

## StorageDataLakeGen2FilesystemSpec

Appears on:[StorageDataLakeGen2Filesystem](#storagedatalakegen2filesystem), [StorageDataLakeGen2FilesystemStatus](#storagedatalakegen2filesystemstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `properties` | ***map[string]string***| ***(Optional)*** |
| `storageAccountID` | ***string***||
## StorageDataLakeGen2FilesystemStatus

Appears on:[StorageDataLakeGen2Filesystem](#storagedatalakegen2filesystem)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageDataLakeGen2FilesystemSpec](#storagedatalakegen2filesystemspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
