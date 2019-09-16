---
title: ApiManagementAPI
menu:
  docs_{{ .version }}:
    identifier: apimanagementapi-azurerm.kubeform.com
    name: ApiManagementAPI
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementAPI
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementAPI` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementAPISpec](#apimanagementapispec)***||
| `status` | ***[ApiManagementAPIStatus](#apimanagementapistatus)***||
## ApiManagementAPISpec

Appears on:[ApiManagementAPI](#apimanagementapi), [ApiManagementAPIStatus](#apimanagementapistatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `displayName` | ***string***||
| `import` | ***[[]ApiManagementAPISpecImport](#apimanagementapispecimport)***| ***(Optional)*** |
| `isCurrent` | ***bool***| ***(Optional)*** |
| `isOnline` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `path` | ***string***||
| `protocols` | ***[]string***||
| `resourceGroupName` | ***string***||
| `revision` | ***string***||
| `serviceURL` | ***string***| ***(Optional)*** |
| `soapPassThrough` | ***bool***| ***(Optional)*** |
| `subscriptionKeyParameterNames` | ***[[]ApiManagementAPISpecSubscriptionKeyParameterNames](#apimanagementapispecsubscriptionkeyparameternames)***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
| `versionSetID` | ***string***| ***(Optional)*** |
## ApiManagementAPISpecImport

Appears on:[ApiManagementAPISpec](#apimanagementapispec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentFormat` | ***string***||
| `contentValue` | ***string***||
| `wsdlSelector` | ***[[]ApiManagementAPISpecImportWsdlSelector](#apimanagementapispecimportwsdlselector)***| ***(Optional)*** |
## ApiManagementAPISpecImportWsdlSelector

Appears on:[ApiManagementAPISpecImport](#apimanagementapispecimport)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `endpointName` | ***string***||
| `serviceName` | ***string***||
## ApiManagementAPISpecSubscriptionKeyParameterNames

Appears on:[ApiManagementAPISpec](#apimanagementapispec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `header` | ***string***||
| `query` | ***string***||
## ApiManagementAPIStatus

Appears on:[ApiManagementAPI](#apimanagementapi)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementAPISpec](#apimanagementapispec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
