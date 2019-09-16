---
title: AutomationSchedule
menu:
  docs_{{ .version }}:
    identifier: automationschedule-azurerm.kubeform.com
    name: AutomationSchedule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationSchedule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationSchedule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationScheduleSpec](#AutomationScheduleSpec)***||
| `status` | ***[AutomationScheduleStatus](#AutomationScheduleStatus)***||
## AutomationScheduleSpec

Appears on:[AutomationSchedule](#AutomationSchedule), [AutomationScheduleStatus](#AutomationScheduleStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***| ***(Optional)*** Deprecated|
| `automationAccountName` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `expiryTime` | ***string***| ***(Optional)*** |
| `frequency` | ***string***||
| `interval` | ***int***| ***(Optional)*** |
| `monthDays` | ***[]int64***| ***(Optional)*** |
| `monthlyOccurrence` | ***[[]AutomationScheduleSpecMonthlyOccurrence](#AutomationScheduleSpecMonthlyOccurrence)***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `startTime` | ***string***| ***(Optional)*** |
| `timezone` | ***string***| ***(Optional)*** |
| `weekDays` | ***[]string***| ***(Optional)*** |
## AutomationScheduleSpecMonthlyOccurrence

Appears on:[AutomationScheduleSpec](#AutomationScheduleSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***string***||
| `occurrence` | ***int***||
## AutomationScheduleStatus

Appears on:[AutomationSchedule](#AutomationSchedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationScheduleSpec](#AutomationScheduleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
