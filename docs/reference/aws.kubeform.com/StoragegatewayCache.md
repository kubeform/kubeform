---
title: StoragegatewayCache
menu:
  docs_{{ .version }}:
    identifier: storagegatewaycache-aws.kubeform.com
    name: StoragegatewayCache
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StoragegatewayCache
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `StoragegatewayCache` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayCacheSpec](#storagegatewaycachespec)***||
| `status` | ***[StoragegatewayCacheStatus](#storagegatewaycachestatus)***||
## StoragegatewayCacheSpec

Appears on:[StoragegatewayCache](#storagegatewaycache), [StoragegatewayCacheStatus](#storagegatewaycachestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `diskID` | ***string***||
| `gatewayArn` | ***string***||
## StoragegatewayCacheStatus

Appears on:[StoragegatewayCache](#storagegatewaycache)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayCacheSpec](#storagegatewaycachespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
