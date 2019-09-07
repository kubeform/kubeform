---
title: SchedulerJobCollection
menu:
  docs_v0.0.1:
    identifier: schedulerjobcollection-azurerm.kubeform.com
    name: SchedulerJobCollection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SchedulerJobCollection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SchedulerJobCollection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SchedulerJobCollectionSpec](#SchedulerJobCollectionSpec)***||
| `status` | ***[SchedulerJobCollectionStatus](#SchedulerJobCollectionStatus)***||
## SchedulerJobCollectionSpec
##### (Appears on:[SchedulerJobCollection](#SchedulerJobCollection), [SchedulerJobCollectionStatus](#SchedulerJobCollectionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `quota` | ***[[]SchedulerJobCollectionSpecQuota](#SchedulerJobCollectionSpecQuota)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SchedulerJobCollectionSpecQuota
##### (Appears on:[SchedulerJobCollectionSpec](#SchedulerJobCollectionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxJobCount` | ***int***| ***(Optional)*** |
| `maxRecurrenceFrequency` | ***string***||
| `maxRecurrenceInterval` | ***int***| ***(Optional)*** |
| `maxRetryInterval` | ***int***| ***(Optional)*** Deprecated|
## SchedulerJobCollectionStatus
##### (Appears on:[SchedulerJobCollection](#SchedulerJobCollection))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SchedulerJobCollectionSpec](#SchedulerJobCollectionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
