---
title: StoragegatewayWorkingStorage
menu:
  docs_{{ .version }}:
    identifier: storagegatewayworkingstorage-aws.kubeform.com
    name: StoragegatewayWorkingStorage
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StoragegatewayWorkingStorage
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `StoragegatewayWorkingStorage` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayWorkingStorageSpec](#StoragegatewayWorkingStorageSpec)***||
| `status` | ***[StoragegatewayWorkingStorageStatus](#StoragegatewayWorkingStorageStatus)***||
## StoragegatewayWorkingStorageSpec

Appears on:[StoragegatewayWorkingStorage](#StoragegatewayWorkingStorage), [StoragegatewayWorkingStorageStatus](#StoragegatewayWorkingStorageStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `diskID` | ***string***||
| `gatewayArn` | ***string***||
## StoragegatewayWorkingStorageStatus

Appears on:[StoragegatewayWorkingStorage](#StoragegatewayWorkingStorage)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayWorkingStorageSpec](#StoragegatewayWorkingStorageSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
