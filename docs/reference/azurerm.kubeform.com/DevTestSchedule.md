---
title: DevTestSchedule
menu:
  docs_{{ .version }}:
    identifier: devtestschedule-azurerm.kubeform.com
    name: DevTestSchedule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DevTestSchedule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DevTestSchedule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevTestScheduleSpec](#devtestschedulespec)***||
| `status` | ***[DevTestScheduleStatus](#devtestschedulestatus)***||
## DevTestScheduleSpec

Appears on:[DevTestSchedule](#devtestschedule), [DevTestScheduleStatus](#devtestschedulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dailyRecurrence` | ***[[]DevTestScheduleSpecDailyRecurrence](#devtestschedulespecdailyrecurrence)***| ***(Optional)*** |
| `hourlyRecurrence` | ***[[]DevTestScheduleSpecHourlyRecurrence](#devtestschedulespechourlyrecurrence)***| ***(Optional)*** |
| `labName` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `notificationSettings` | ***[[]DevTestScheduleSpecNotificationSettings](#devtestschedulespecnotificationsettings)***||
| `resourceGroupName` | ***string***||
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `taskType` | ***string***||
| `timeZoneID` | ***string***||
| `weeklyRecurrence` | ***[[]DevTestScheduleSpecWeeklyRecurrence](#devtestschedulespecweeklyrecurrence)***| ***(Optional)*** |
## DevTestScheduleSpecDailyRecurrence

Appears on:[DevTestScheduleSpec](#devtestschedulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `time` | ***string***||
## DevTestScheduleSpecHourlyRecurrence

Appears on:[DevTestScheduleSpec](#devtestschedulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `minute` | ***int64***||
## DevTestScheduleSpecNotificationSettings

Appears on:[DevTestScheduleSpec](#devtestschedulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `status` | ***string***| ***(Optional)*** |
| `timeInMinutes` | ***int64***| ***(Optional)*** |
| `webhookURL` | ***string***| ***(Optional)*** |
## DevTestScheduleSpecWeeklyRecurrence

Appears on:[DevTestScheduleSpec](#devtestschedulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `time` | ***string***||
| `weekDays` | ***[]string***| ***(Optional)*** |
## DevTestScheduleStatus

Appears on:[DevTestSchedule](#devtestschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevTestScheduleSpec](#devtestschedulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DevTestScheduleStatus](#devtestschedulestatus)

---
