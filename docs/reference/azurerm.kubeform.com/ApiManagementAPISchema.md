---
title: ApiManagementAPISchema
menu:
  docs_v0.0.1:
    identifier: apimanagementapischema-azurerm.kubeform.com
    name: ApiManagementAPISchema
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApiManagementAPISchema
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementAPISchema` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementAPISchemaSpec](#ApiManagementAPISchemaSpec)***||
| `status` | ***[ApiManagementAPISchemaStatus](#ApiManagementAPISchemaStatus)***||
## ApiManagementAPISchemaSpec
##### (Appears on:[ApiManagementAPISchema](#ApiManagementAPISchema), [ApiManagementAPISchemaStatus](#ApiManagementAPISchemaStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `apiName` | ***string***||
| `contentType` | ***string***||
| `resourceGroupName` | ***string***||
| `schemaID` | ***string***||
| `value` | ***string***||
## ApiManagementAPISchemaStatus
##### (Appears on:[ApiManagementAPISchema](#ApiManagementAPISchema))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementAPISchemaSpec](#ApiManagementAPISchemaSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
