---
title: ResourceGroup
menu:
  docs_v0.0.1:
    identifier: resourcegroup-azurerm.kubeform.com
    name: ResourceGroup
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ResourceGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ResourceGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ResourceGroupSpec](#ResourceGroupSpec)***||
| `status` | ***[ResourceGroupStatus](#ResourceGroupStatus)***||
## ResourceGroupSpec
##### (Appears on:[ResourceGroup](#ResourceGroup), [ResourceGroupStatus](#ResourceGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## ResourceGroupStatus
##### (Appears on:[ResourceGroup](#ResourceGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ResourceGroupSpec](#ResourceGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
