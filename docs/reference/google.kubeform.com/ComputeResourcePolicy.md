---
title: ComputeResourcePolicy
menu:
  docs_{{ .version }}:
    identifier: computeresourcepolicy-google.kubeform.com
    name: ComputeResourcePolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeResourcePolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeResourcePolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeResourcePolicySpec](#computeresourcepolicyspec)***||
| `status` | ***[ComputeResourcePolicyStatus](#computeresourcepolicystatus)***||
## ComputeResourcePolicySpec

Appears on:[ComputeResourcePolicy](#computeresourcepolicy), [ComputeResourcePolicyStatus](#computeresourcepolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `snapshotSchedulePolicy` | ***[[]ComputeResourcePolicySpecSnapshotSchedulePolicy](#computeresourcepolicyspecsnapshotschedulepolicy)***| ***(Optional)*** |
## ComputeResourcePolicySpecSnapshotSchedulePolicy

Appears on:[ComputeResourcePolicySpec](#computeresourcepolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `retentionPolicy` | ***[[]ComputeResourcePolicySpecSnapshotSchedulePolicyRetentionPolicy](#computeresourcepolicyspecsnapshotschedulepolicyretentionpolicy)***| ***(Optional)*** |
| `schedule` | ***[[]ComputeResourcePolicySpecSnapshotSchedulePolicySchedule](#computeresourcepolicyspecsnapshotschedulepolicyschedule)***||
| `snapshotProperties` | ***[[]ComputeResourcePolicySpecSnapshotSchedulePolicySnapshotProperties](#computeresourcepolicyspecsnapshotschedulepolicysnapshotproperties)***| ***(Optional)*** |
## ComputeResourcePolicySpecSnapshotSchedulePolicyRetentionPolicy

Appears on:[ComputeResourcePolicySpecSnapshotSchedulePolicy](#computeresourcepolicyspecsnapshotschedulepolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxRetentionDays` | ***int64***||
| `onSourceDiskDelete` | ***string***| ***(Optional)*** |
## ComputeResourcePolicySpecSnapshotSchedulePolicySchedule

Appears on:[ComputeResourcePolicySpecSnapshotSchedulePolicy](#computeresourcepolicyspecsnapshotschedulepolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dailySchedule` | ***[[]ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleDailySchedule](#computeresourcepolicyspecsnapshotschedulepolicyscheduledailyschedule)***| ***(Optional)*** |
| `hourlySchedule` | ***[[]ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleHourlySchedule](#computeresourcepolicyspecsnapshotschedulepolicyschedulehourlyschedule)***| ***(Optional)*** |
| `weeklySchedule` | ***[[]ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleWeeklySchedule](#computeresourcepolicyspecsnapshotschedulepolicyscheduleweeklyschedule)***| ***(Optional)*** |
## ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleDailySchedule

Appears on:[ComputeResourcePolicySpecSnapshotSchedulePolicySchedule](#computeresourcepolicyspecsnapshotschedulepolicyschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `daysInCycle` | ***int64***||
| `startTime` | ***string***||
## ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleHourlySchedule

Appears on:[ComputeResourcePolicySpecSnapshotSchedulePolicySchedule](#computeresourcepolicyspecsnapshotschedulepolicyschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hoursInCycle` | ***int64***||
| `startTime` | ***string***||
## ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleWeeklySchedule

Appears on:[ComputeResourcePolicySpecSnapshotSchedulePolicySchedule](#computeresourcepolicyspecsnapshotschedulepolicyschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dayOfWeeks` | ***[[]ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks](#computeresourcepolicyspecsnapshotschedulepolicyscheduleweeklyscheduledayofweeks)***||
## ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleWeeklyScheduleDayOfWeeks

Appears on:[ComputeResourcePolicySpecSnapshotSchedulePolicyScheduleWeeklySchedule](#computeresourcepolicyspecsnapshotschedulepolicyscheduleweeklyschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***string***||
| `startTime` | ***string***||
## ComputeResourcePolicySpecSnapshotSchedulePolicySnapshotProperties

Appears on:[ComputeResourcePolicySpecSnapshotSchedulePolicy](#computeresourcepolicyspecsnapshotschedulepolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `guestFlush` | ***bool***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `storageLocations` | ***[]string***| ***(Optional)*** |
## ComputeResourcePolicyStatus

Appears on:[ComputeResourcePolicy](#computeresourcepolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeResourcePolicySpec](#computeresourcepolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeResourcePolicyStatus](#computeresourcepolicystatus)

---
