---
title: RoleAssignment
menu:
  docs_{{ .version }}:
    identifier: roleassignment-azurerm.kubeform.com
    name: RoleAssignment
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RoleAssignment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RoleAssignment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RoleAssignmentSpec](#roleassignmentspec)***||
| `status` | ***[RoleAssignmentStatus](#roleassignmentstatus)***||
## Phase(`string` alias)

Appears on:[RoleAssignmentStatus](#roleassignmentstatus)

## RoleAssignmentSpec

Appears on:[RoleAssignment](#roleassignment), [RoleAssignmentStatus](#roleassignmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `principalID` | ***string***||
| `principalType` | ***string***| ***(Optional)*** |
| `roleDefinitionID` | ***string***| ***(Optional)*** |
| `roleDefinitionName` | ***string***| ***(Optional)*** |
| `scope` | ***string***||
| `skipServicePrincipalAadCheck` | ***bool***| ***(Optional)*** |
## RoleAssignmentStatus

Appears on:[RoleAssignment](#roleassignment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RoleAssignmentSpec](#roleassignmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
