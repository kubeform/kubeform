---
title: AutomationVariableDatetime
menu:
  docs_{{ .version }}:
    identifier: automationvariabledatetime-azurerm.kubeform.com
    name: AutomationVariableDatetime
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationVariableDatetime
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationVariableDatetime` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationVariableDatetimeSpec](#automationvariabledatetimespec)***||
| `status` | ***[AutomationVariableDatetimeStatus](#automationvariabledatetimestatus)***||
## AutomationVariableDatetimeSpec

Appears on:[AutomationVariableDatetime](#automationvariabledatetime), [AutomationVariableDatetimeStatus](#automationvariabledatetimestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `automationAccountName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `value` | ***string***| ***(Optional)*** |
## AutomationVariableDatetimeStatus

Appears on:[AutomationVariableDatetime](#automationvariabledatetime)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationVariableDatetimeSpec](#automationvariabledatetimespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutomationVariableDatetimeStatus](#automationvariabledatetimestatus)

---
