---
title: ApiManagementProductGroup
menu:
  docs_{{ .version }}:
    identifier: apimanagementproductgroup-azurerm.kubeform.com
    name: ApiManagementProductGroup
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementProductGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementProductGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementProductGroupSpec](#ApiManagementProductGroupSpec)***||
| `status` | ***[ApiManagementProductGroupStatus](#ApiManagementProductGroupStatus)***||
## ApiManagementProductGroupSpec
##### (Appears on:[ApiManagementProductGroup](#ApiManagementProductGroup), [ApiManagementProductGroupStatus](#ApiManagementProductGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `groupName` | ***string***||
| `productID` | ***string***||
| `resourceGroupName` | ***string***||
## ApiManagementProductGroupStatus
##### (Appears on:[ApiManagementProductGroup](#ApiManagementProductGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementProductGroupSpec](#ApiManagementProductGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
