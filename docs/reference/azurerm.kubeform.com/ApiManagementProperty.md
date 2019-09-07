---
title: ApiManagementProperty
menu:
  docs_v0.0.1:
    identifier: apimanagementproperty-azurerm.kubeform.com
    name: ApiManagementProperty
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApiManagementProperty
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementProperty` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementPropertySpec](#ApiManagementPropertySpec)***||
| `status` | ***[ApiManagementPropertyStatus](#ApiManagementPropertyStatus)***||
## ApiManagementPropertySpec
##### (Appears on:[ApiManagementProperty](#ApiManagementProperty), [ApiManagementPropertyStatus](#ApiManagementPropertyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `displayName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `secret` | ***bool***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `value` | ***string***||
## ApiManagementPropertyStatus
##### (Appears on:[ApiManagementProperty](#ApiManagementProperty))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementPropertySpec](#ApiManagementPropertySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
