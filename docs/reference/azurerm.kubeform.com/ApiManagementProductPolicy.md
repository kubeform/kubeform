---
title: ApiManagementProductPolicy
menu:
  docs_{{ .version }}:
    identifier: apimanagementproductpolicy-azurerm.kubeform.com
    name: ApiManagementProductPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementProductPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementProductPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementProductPolicySpec](#apimanagementproductpolicyspec)***||
| `status` | ***[ApiManagementProductPolicyStatus](#apimanagementproductpolicystatus)***||
## ApiManagementProductPolicySpec

Appears on:[ApiManagementProductPolicy](#apimanagementproductpolicy), [ApiManagementProductPolicyStatus](#apimanagementproductpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `productID` | ***string***||
| `resourceGroupName` | ***string***||
| `xmlContent` | ***string***| ***(Optional)*** |
| `xmlLink` | ***string***| ***(Optional)*** |
## ApiManagementProductPolicyStatus

Appears on:[ApiManagementProductPolicy](#apimanagementproductpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementProductPolicySpec](#apimanagementproductpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiManagementProductPolicyStatus](#apimanagementproductpolicystatus)

---
