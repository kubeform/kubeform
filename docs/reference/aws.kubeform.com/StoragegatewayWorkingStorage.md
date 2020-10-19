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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayWorkingStorageSpec](#storagegatewayworkingstoragespec)***||
| `status` | ***[StoragegatewayWorkingStorageStatus](#storagegatewayworkingstoragestatus)***||
## Phase(`string` alias)

Appears on:[StoragegatewayWorkingStorageStatus](#storagegatewayworkingstoragestatus)

## StoragegatewayWorkingStorageSpec

Appears on:[StoragegatewayWorkingStorage](#storagegatewayworkingstorage), [StoragegatewayWorkingStorageStatus](#storagegatewayworkingstoragestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `diskID` | ***string***||
| `gatewayArn` | ***string***||
## StoragegatewayWorkingStorageStatus

Appears on:[StoragegatewayWorkingStorage](#storagegatewayworkingstorage)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayWorkingStorageSpec](#storagegatewayworkingstoragespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
