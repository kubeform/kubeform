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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationScheduleSpec](#automationschedulespec)***||
| `status` | ***[AutomationScheduleStatus](#automationschedulestatus)***||
## AutomationScheduleSpec

Appears on:[AutomationSchedule](#automationschedule), [AutomationScheduleStatus](#automationschedulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***| ***(Optional)*** Deprecated|
| `automationAccountName` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `expiryTime` | ***string***| ***(Optional)*** |
| `frequency` | ***string***||
| `interval` | ***int64***| ***(Optional)*** |
| `monthDays` | ***[]int64***| ***(Optional)*** |
| `monthlyOccurrence` | ***[[]AutomationScheduleSpecMonthlyOccurrence](#automationschedulespecmonthlyoccurrence)***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `startTime` | ***string***| ***(Optional)*** |
| `timezone` | ***string***| ***(Optional)*** |
| `weekDays` | ***[]string***| ***(Optional)*** |
## AutomationScheduleSpecMonthlyOccurrence

Appears on:[AutomationScheduleSpec](#automationschedulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***string***||
| `occurrence` | ***int64***||
## AutomationScheduleStatus

Appears on:[AutomationSchedule](#automationschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationScheduleSpec](#automationschedulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutomationScheduleStatus](#automationschedulestatus)

---
