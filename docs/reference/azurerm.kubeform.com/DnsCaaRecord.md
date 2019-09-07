---
title: DnsCaaRecord
menu:
  docs_v0.0.1:
    identifier: dnscaarecord-azurerm.kubeform.com
    name: DnsCaaRecord
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DnsCaaRecord
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DnsCaaRecord` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DnsCaaRecordSpec](#DnsCaaRecordSpec)***||
| `status` | ***[DnsCaaRecordStatus](#DnsCaaRecordStatus)***||
## DnsCaaRecordSpec
##### (Appears on:[DnsCaaRecord](#DnsCaaRecord), [DnsCaaRecordStatus](#DnsCaaRecordStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `record` | ***[[]DnsCaaRecordSpecRecord](#DnsCaaRecordSpecRecord)***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `ttl` | ***int***||
| `zoneName` | ***string***||
## DnsCaaRecordSpecRecord
##### (Appears on:[DnsCaaRecordSpec](#DnsCaaRecordSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `flags` | ***int***||
| `tag` | ***string***||
| `value` | ***string***||
## DnsCaaRecordStatus
##### (Appears on:[DnsCaaRecord](#DnsCaaRecord))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DnsCaaRecordSpec](#DnsCaaRecordSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
