---
title: PolicyAssignment
menu:
  docs_{{ .version }}:
    identifier: policyassignment-azurerm.kubeform.com
    name: PolicyAssignment
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PolicyAssignment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PolicyAssignment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PolicyAssignmentSpec](#policyassignmentspec)***||
| `status` | ***[PolicyAssignmentStatus](#policyassignmentstatus)***||
## Phase(`string` alias)

Appears on:[PolicyAssignmentStatus](#policyassignmentstatus)

## PolicyAssignmentSpec

Appears on:[PolicyAssignment](#policyassignment), [PolicyAssignmentStatus](#policyassignmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `identity` | ***[[]PolicyAssignmentSpecIdentity](#policyassignmentspecidentity)***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `notScopes` | ***[]string***| ***(Optional)*** |
| `parameters` | ***string***| ***(Optional)*** |
| `policyDefinitionID` | ***string***||
| `scope` | ***string***||
## PolicyAssignmentSpecIdentity

Appears on:[PolicyAssignmentSpec](#policyassignmentspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `principalID` | ***string***| ***(Optional)*** |
| `tenantID` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## PolicyAssignmentStatus

Appears on:[PolicyAssignment](#policyassignment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PolicyAssignmentSpec](#policyassignmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
