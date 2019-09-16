---
title: AppServiceActiveSlot
menu:
  docs_{{ .version }}:
    identifier: appserviceactiveslot-azurerm.kubeform.com
    name: AppServiceActiveSlot
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppServiceActiveSlot
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AppServiceActiveSlot` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppServiceActiveSlotSpec](#AppServiceActiveSlotSpec)***||
| `status` | ***[AppServiceActiveSlotStatus](#AppServiceActiveSlotStatus)***||
## AppServiceActiveSlotSpec

Appears on:[AppServiceActiveSlot](#AppServiceActiveSlot), [AppServiceActiveSlotStatus](#AppServiceActiveSlotStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appServiceName` | ***string***||
| `appServiceSlotName` | ***string***||
| `resourceGroupName` | ***string***||
## AppServiceActiveSlotStatus

Appears on:[AppServiceActiveSlot](#AppServiceActiveSlot)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppServiceActiveSlotSpec](#AppServiceActiveSlotSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
