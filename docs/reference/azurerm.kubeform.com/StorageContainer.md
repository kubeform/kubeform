---
title: StorageContainer
menu:
  docs_{{ .version }}:
    identifier: storagecontainer-azurerm.kubeform.com
    name: StorageContainer
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageContainer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageContainer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageContainerSpec](#StorageContainerSpec)***||
| `status` | ***[StorageContainerStatus](#StorageContainerStatus)***||
## StorageContainerSpec
##### (Appears on:[StorageContainer](#StorageContainer), [StorageContainerStatus](#StorageContainerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `containerAccessType` | ***string***| ***(Optional)*** |
| `hasImmutabilityPolicy` | ***bool***| ***(Optional)*** |
| `hasLegalHold` | ***bool***| ***(Optional)*** |
| `metadata` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `properties` | ***map[string]string***| ***(Optional)*** Deprecated|
| `resourceGroupName` | ***string***| ***(Optional)*** Deprecated|
| `storageAccountName` | ***string***||
## StorageContainerStatus
##### (Appears on:[StorageContainer](#StorageContainer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageContainerSpec](#StorageContainerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
