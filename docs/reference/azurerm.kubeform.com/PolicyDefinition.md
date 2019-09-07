---
title: PolicyDefinition
menu:
  docs_v0.0.1:
    identifier: policydefinition-azurerm.kubeform.com
    name: PolicyDefinition
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## PolicyDefinition
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PolicyDefinition` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PolicyDefinitionSpec](#PolicyDefinitionSpec)***||
| `status` | ***[PolicyDefinitionStatus](#PolicyDefinitionStatus)***||
## PolicyDefinitionSpec
##### (Appears on:[PolicyDefinition](#PolicyDefinition), [PolicyDefinitionStatus](#PolicyDefinitionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `displayName` | ***string***||
| `managementGroupID` | ***string***| ***(Optional)*** |
| `metadata` | ***string***| ***(Optional)*** |
| `mode` | ***string***||
| `name` | ***string***||
| `parameters` | ***string***| ***(Optional)*** |
| `policyRule` | ***string***| ***(Optional)*** |
| `policyType` | ***string***||
## PolicyDefinitionStatus
##### (Appears on:[PolicyDefinition](#PolicyDefinition))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PolicyDefinitionSpec](#PolicyDefinitionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
