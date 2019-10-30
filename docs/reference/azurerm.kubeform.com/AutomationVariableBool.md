---
title: AutomationVariableBool
menu:
  docs_{{ .version }}:
    identifier: automationvariablebool-azurerm.kubeform.com
    name: AutomationVariableBool
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationVariableBool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationVariableBool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationVariableBoolSpec](#automationvariableboolspec)***||
| `status` | ***[AutomationVariableBoolStatus](#automationvariableboolstatus)***||
## AutomationVariableBoolSpec

Appears on:[AutomationVariableBool](#automationvariablebool), [AutomationVariableBoolStatus](#automationvariableboolstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `automationAccountName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `value` | ***bool***| ***(Optional)*** |
## AutomationVariableBoolStatus

Appears on:[AutomationVariableBool](#automationvariablebool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationVariableBoolSpec](#automationvariableboolspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutomationVariableBoolStatus](#automationvariableboolstatus)

---
