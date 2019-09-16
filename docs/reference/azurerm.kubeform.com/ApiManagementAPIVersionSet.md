---
title: ApiManagementAPIVersionSet
menu:
  docs_{{ .version }}:
    identifier: apimanagementapiversionset-azurerm.kubeform.com
    name: ApiManagementAPIVersionSet
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementAPIVersionSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementAPIVersionSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementAPIVersionSetSpec](#ApiManagementAPIVersionSetSpec)***||
| `status` | ***[ApiManagementAPIVersionSetStatus](#ApiManagementAPIVersionSetStatus)***||
## ApiManagementAPIVersionSetSpec

Appears on:[ApiManagementAPIVersionSet](#ApiManagementAPIVersionSet), [ApiManagementAPIVersionSetStatus](#ApiManagementAPIVersionSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiManagementName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `displayName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `versionHeaderName` | ***string***| ***(Optional)*** |
| `versionQueryName` | ***string***| ***(Optional)*** |
| `versioningScheme` | ***string***||
## ApiManagementAPIVersionSetStatus

Appears on:[ApiManagementAPIVersionSet](#ApiManagementAPIVersionSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementAPIVersionSetSpec](#ApiManagementAPIVersionSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
