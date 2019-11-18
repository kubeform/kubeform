---
title: AutomationVariableInt
menu:
  docs_{{ .version }}:
    identifier: automationvariableint-azurerm.kubeform.com
    name: AutomationVariableInt
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationVariableInt
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationVariableInt` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationVariableIntSpec](#automationvariableintspec)***||
| `status` | ***[AutomationVariableIntStatus](#automationvariableintstatus)***||
## AutomationVariableIntSpec

Appears on:[AutomationVariableInt](#automationvariableint), [AutomationVariableIntStatus](#automationvariableintstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `automationAccountName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `value` | ***int64***| ***(Optional)*** |
## AutomationVariableIntStatus

Appears on:[AutomationVariableInt](#automationvariableint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationVariableIntSpec](#automationvariableintspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutomationVariableIntStatus](#automationvariableintstatus)

---
