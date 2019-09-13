---
title: MediaServicesAccount
menu:
  docs_{{ .version }}:
    identifier: mediaservicesaccount-azurerm.kubeform.com
    name: MediaServicesAccount
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MediaServicesAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MediaServicesAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MediaServicesAccountSpec](#MediaServicesAccountSpec)***||
| `status` | ***[MediaServicesAccountStatus](#MediaServicesAccountStatus)***||
## MediaServicesAccountSpec
##### (Appears on:[MediaServicesAccount](#MediaServicesAccount), [MediaServicesAccountStatus](#MediaServicesAccountStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `storageAccount` | ***[[]MediaServicesAccountSpecStorageAccount](#MediaServicesAccountSpecStorageAccount)***||
## MediaServicesAccountSpecStorageAccount
##### (Appears on:[MediaServicesAccountSpec](#MediaServicesAccountSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***||
| `isPrimary` | ***bool***| ***(Optional)*** |
## MediaServicesAccountStatus
##### (Appears on:[MediaServicesAccount](#MediaServicesAccount))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MediaServicesAccountSpec](#MediaServicesAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
