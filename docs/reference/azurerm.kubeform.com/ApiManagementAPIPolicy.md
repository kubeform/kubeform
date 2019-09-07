---
title: ApiManagementAPIPolicy
menu:
  docs_v0.0.1:
    identifier: apimanagementapipolicy-azurerm.kubeform.com
    name: ApiManagementAPIPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApiManagementAPIPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementAPIPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementAPIPolicySpec](#ApiManagementAPIPolicySpec)***||
| `status` | ***[ApiManagementAPIPolicyStatus](#ApiManagementAPIPolicyStatus)***||
## ApiManagementAPIPolicySpec
##### (Appears on:[ApiManagementAPIPolicy](#ApiManagementAPIPolicy), [ApiManagementAPIPolicyStatus](#ApiManagementAPIPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `apiName` | ***string***||
| `resourceGroupName` | ***string***||
| `xmlContent` | ***string***| ***(Optional)*** |
| `xmlLink` | ***string***| ***(Optional)*** |
## ApiManagementAPIPolicyStatus
##### (Appears on:[ApiManagementAPIPolicy](#ApiManagementAPIPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementAPIPolicySpec](#ApiManagementAPIPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
