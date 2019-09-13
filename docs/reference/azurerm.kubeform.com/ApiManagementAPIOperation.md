---
title: ApiManagementAPIOperation
menu:
  docs_{{ .version }}:
    identifier: apimanagementapioperation-azurerm.kubeform.com
    name: ApiManagementAPIOperation
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementAPIOperation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementAPIOperation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementAPIOperationSpec](#ApiManagementAPIOperationSpec)***||
| `status` | ***[ApiManagementAPIOperationStatus](#ApiManagementAPIOperationStatus)***||
## ApiManagementAPIOperationSpec
##### (Appears on:[ApiManagementAPIOperation](#ApiManagementAPIOperation), [ApiManagementAPIOperationStatus](#ApiManagementAPIOperationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `apiName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `displayName` | ***string***||
| `method` | ***string***||
| `operationID` | ***string***||
| `request` | ***[[]ApiManagementAPIOperationSpecRequest](#ApiManagementAPIOperationSpecRequest)***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `response` | ***[[]ApiManagementAPIOperationSpecResponse](#ApiManagementAPIOperationSpecResponse)***| ***(Optional)*** |
| `templateParameter` | ***[[]ApiManagementAPIOperationSpecTemplateParameter](#ApiManagementAPIOperationSpecTemplateParameter)***| ***(Optional)*** |
| `urlTemplate` | ***string***||
## ApiManagementAPIOperationSpecRequest
##### (Appears on:[ApiManagementAPIOperationSpec](#ApiManagementAPIOperationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `header` | ***[[]ApiManagementAPIOperationSpecRequestHeader](#ApiManagementAPIOperationSpecRequestHeader)***| ***(Optional)*** |
| `queryParameter` | ***[[]ApiManagementAPIOperationSpecRequestQueryParameter](#ApiManagementAPIOperationSpecRequestQueryParameter)***| ***(Optional)*** |
| `representation` | ***[[]ApiManagementAPIOperationSpecRequestRepresentation](#ApiManagementAPIOperationSpecRequestRepresentation)***| ***(Optional)*** |
## ApiManagementAPIOperationSpecRequestHeader
##### (Appears on:[ApiManagementAPIOperationSpecRequest](#ApiManagementAPIOperationSpecRequest))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `required` | ***bool***||
| `type` | ***string***||
| `values` | ***[]string***| ***(Optional)*** |
## ApiManagementAPIOperationSpecRequestQueryParameter
##### (Appears on:[ApiManagementAPIOperationSpecRequest](#ApiManagementAPIOperationSpecRequest))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `required` | ***bool***||
| `type` | ***string***||
| `values` | ***[]string***| ***(Optional)*** |
## ApiManagementAPIOperationSpecRequestRepresentation
##### (Appears on:[ApiManagementAPIOperationSpecRequest](#ApiManagementAPIOperationSpecRequest))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***||
| `formParameter` | ***[[]ApiManagementAPIOperationSpecRequestRepresentationFormParameter](#ApiManagementAPIOperationSpecRequestRepresentationFormParameter)***| ***(Optional)*** |
| `sample` | ***string***| ***(Optional)*** |
| `schemaID` | ***string***| ***(Optional)*** |
| `typeName` | ***string***| ***(Optional)*** |
## ApiManagementAPIOperationSpecRequestRepresentationFormParameter
##### (Appears on:[ApiManagementAPIOperationSpecRequestRepresentation](#ApiManagementAPIOperationSpecRequestRepresentation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `required` | ***bool***||
| `type` | ***string***||
| `values` | ***[]string***| ***(Optional)*** |
## ApiManagementAPIOperationSpecResponse
##### (Appears on:[ApiManagementAPIOperationSpec](#ApiManagementAPIOperationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `header` | ***[[]ApiManagementAPIOperationSpecResponseHeader](#ApiManagementAPIOperationSpecResponseHeader)***| ***(Optional)*** |
| `representation` | ***[[]ApiManagementAPIOperationSpecResponseRepresentation](#ApiManagementAPIOperationSpecResponseRepresentation)***| ***(Optional)*** |
| `statusCode` | ***int***||
## ApiManagementAPIOperationSpecResponseHeader
##### (Appears on:[ApiManagementAPIOperationSpecResponse](#ApiManagementAPIOperationSpecResponse))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `required` | ***bool***||
| `type` | ***string***||
| `values` | ***[]string***| ***(Optional)*** |
## ApiManagementAPIOperationSpecResponseRepresentation
##### (Appears on:[ApiManagementAPIOperationSpecResponse](#ApiManagementAPIOperationSpecResponse))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***||
| `formParameter` | ***[[]ApiManagementAPIOperationSpecResponseRepresentationFormParameter](#ApiManagementAPIOperationSpecResponseRepresentationFormParameter)***| ***(Optional)*** |
| `sample` | ***string***| ***(Optional)*** |
| `schemaID` | ***string***| ***(Optional)*** |
| `typeName` | ***string***| ***(Optional)*** |
## ApiManagementAPIOperationSpecResponseRepresentationFormParameter
##### (Appears on:[ApiManagementAPIOperationSpecResponseRepresentation](#ApiManagementAPIOperationSpecResponseRepresentation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `required` | ***bool***||
| `type` | ***string***||
| `values` | ***[]string***| ***(Optional)*** |
## ApiManagementAPIOperationSpecTemplateParameter
##### (Appears on:[ApiManagementAPIOperationSpec](#ApiManagementAPIOperationSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `required` | ***bool***||
| `type` | ***string***||
| `values` | ***[]string***| ***(Optional)*** |
## ApiManagementAPIOperationStatus
##### (Appears on:[ApiManagementAPIOperation](#ApiManagementAPIOperation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementAPIOperationSpec](#ApiManagementAPIOperationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
