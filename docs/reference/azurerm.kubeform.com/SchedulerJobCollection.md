---
title: SchedulerJobCollection
menu:
  docs_{{ .version }}:
    identifier: schedulerjobcollection-azurerm.kubeform.com
    name: SchedulerJobCollection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SchedulerJobCollection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SchedulerJobCollection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SchedulerJobCollectionSpec](#schedulerjobcollectionspec)***||
| `status` | ***[SchedulerJobCollectionStatus](#schedulerjobcollectionstatus)***||
## Phase(`string` alias)

Appears on:[SchedulerJobCollectionStatus](#schedulerjobcollectionstatus)

## SchedulerJobCollectionSpec

Appears on:[SchedulerJobCollection](#schedulerjobcollection), [SchedulerJobCollectionStatus](#schedulerjobcollectionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `quota` | ***[[]SchedulerJobCollectionSpecQuota](#schedulerjobcollectionspecquota)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SchedulerJobCollectionSpecQuota

Appears on:[SchedulerJobCollectionSpec](#schedulerjobcollectionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxJobCount` | ***int64***| ***(Optional)*** |
| `maxRecurrenceFrequency` | ***string***||
| `maxRecurrenceInterval` | ***int64***| ***(Optional)*** |
| `maxRetryInterval` | ***int64***| ***(Optional)*** Deprecated|
## SchedulerJobCollectionStatus

Appears on:[SchedulerJobCollection](#schedulerjobcollection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SchedulerJobCollectionSpec](#schedulerjobcollectionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
