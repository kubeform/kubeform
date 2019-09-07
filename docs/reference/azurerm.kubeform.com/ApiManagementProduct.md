---
title: ApiManagementProduct
menu:
  docs_v0.0.1:
    identifier: apimanagementproduct-azurerm.kubeform.com
    name: ApiManagementProduct
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApiManagementProduct
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementProduct` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementProductSpec](#ApiManagementProductSpec)***||
| `status` | ***[ApiManagementProductStatus](#ApiManagementProductStatus)***||
## ApiManagementProductSpec
##### (Appears on:[ApiManagementProduct](#ApiManagementProduct), [ApiManagementProductStatus](#ApiManagementProductStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `approvalRequired` | ***bool***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `displayName` | ***string***||
| `productID` | ***string***||
| `published` | ***bool***||
| `resourceGroupName` | ***string***||
| `subscriptionRequired` | ***bool***||
| `subscriptionsLimit` | ***int***| ***(Optional)*** |
| `terms` | ***string***| ***(Optional)*** |
## ApiManagementProductStatus
##### (Appears on:[ApiManagementProduct](#ApiManagementProduct))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementProductSpec](#ApiManagementProductSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
