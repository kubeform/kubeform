---
title: NetworkWatcher
menu:
  docs_v0.0.1:
    identifier: networkwatcher-azurerm.kubeform.com
    name: NetworkWatcher
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## NetworkWatcher
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkWatcher` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkWatcherSpec](#NetworkWatcherSpec)***||
| `status` | ***[NetworkWatcherStatus](#NetworkWatcherStatus)***||
## NetworkWatcherSpec
##### (Appears on:[NetworkWatcher](#NetworkWatcher), [NetworkWatcherStatus](#NetworkWatcherStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## NetworkWatcherStatus
##### (Appears on:[NetworkWatcher](#NetworkWatcher))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkWatcherSpec](#NetworkWatcherSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
