---
title: RoleDefinition
menu:
  docs_{{ .version }}:
    identifier: roledefinition-azurerm.kubeform.com
    name: RoleDefinition
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RoleDefinition
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RoleDefinition` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RoleDefinitionSpec](#RoleDefinitionSpec)***||
| `status` | ***[RoleDefinitionStatus](#RoleDefinitionStatus)***||
## RoleDefinitionSpec

Appears on:[RoleDefinition](#RoleDefinition), [RoleDefinitionStatus](#RoleDefinitionStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `assignableScopes` | ***[]string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `permissions` | ***[[]RoleDefinitionSpecPermissions](#RoleDefinitionSpecPermissions)***||
| `roleDefinitionID` | ***string***| ***(Optional)*** |
| `scope` | ***string***||
## RoleDefinitionSpecPermissions

Appears on:[RoleDefinitionSpec](#RoleDefinitionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `actions` | ***[]string***| ***(Optional)*** |
| `dataActions` | ***[]string***| ***(Optional)*** |
| `notActions` | ***[]string***| ***(Optional)*** |
| `notDataActions` | ***[]string***| ***(Optional)*** |
## RoleDefinitionStatus

Appears on:[RoleDefinition](#RoleDefinition)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RoleDefinitionSpec](#RoleDefinitionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
