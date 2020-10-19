---
title: ApiManagementAPIOperationPolicy
menu:
  docs_{{ .version }}:
    identifier: apimanagementapioperationpolicy-azurerm.kubeform.com
    name: ApiManagementAPIOperationPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementAPIOperationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementAPIOperationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementAPIOperationPolicySpec](#apimanagementapioperationpolicyspec)***||
| `status` | ***[ApiManagementAPIOperationPolicyStatus](#apimanagementapioperationpolicystatus)***||
## ApiManagementAPIOperationPolicySpec

Appears on:[ApiManagementAPIOperationPolicy](#apimanagementapioperationpolicy), [ApiManagementAPIOperationPolicyStatus](#apimanagementapioperationpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `apiName` | ***string***||
| `operationID` | ***string***||
| `resourceGroupName` | ***string***||
| `xmlContent` | ***string***| ***(Optional)*** |
| `xmlLink` | ***string***| ***(Optional)*** |
## ApiManagementAPIOperationPolicyStatus

Appears on:[ApiManagementAPIOperationPolicy](#apimanagementapioperationpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementAPIOperationPolicySpec](#apimanagementapioperationpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiManagementAPIOperationPolicyStatus](#apimanagementapioperationpolicystatus)

---
