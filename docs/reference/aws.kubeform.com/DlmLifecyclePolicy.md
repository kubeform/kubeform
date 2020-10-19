---
title: DlmLifecyclePolicy
menu:
  docs_{{ .version }}:
    identifier: dlmlifecyclepolicy-aws.kubeform.com
    name: DlmLifecyclePolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DlmLifecyclePolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DlmLifecyclePolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DlmLifecyclePolicySpec](#dlmlifecyclepolicyspec)***||
| `status` | ***[DlmLifecyclePolicyStatus](#dlmlifecyclepolicystatus)***||
## DlmLifecyclePolicySpec

Appears on:[DlmLifecyclePolicy](#dlmlifecyclepolicy), [DlmLifecyclePolicyStatus](#dlmlifecyclepolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***||
| `executionRoleArn` | ***string***||
| `policyDetails` | ***[[]DlmLifecyclePolicySpecPolicyDetails](#dlmlifecyclepolicyspecpolicydetails)***||
| `state` | ***string***| ***(Optional)*** |
## DlmLifecyclePolicySpecPolicyDetails

Appears on:[DlmLifecyclePolicySpec](#dlmlifecyclepolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceTypes` | ***[]string***||
| `schedule` | ***[[]DlmLifecyclePolicySpecPolicyDetailsSchedule](#dlmlifecyclepolicyspecpolicydetailsschedule)***||
| `targetTags` | ***map[string]string***||
## DlmLifecyclePolicySpecPolicyDetailsSchedule

Appears on:[DlmLifecyclePolicySpecPolicyDetails](#dlmlifecyclepolicyspecpolicydetails)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `copyTags` | ***bool***| ***(Optional)*** |
| `createRule` | ***[[]DlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule](#dlmlifecyclepolicyspecpolicydetailsschedulecreaterule)***||
| `name` | ***string***||
| `retainRule` | ***[[]DlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule](#dlmlifecyclepolicyspecpolicydetailsscheduleretainrule)***||
| `tagsToAdd` | ***map[string]string***| ***(Optional)*** |
## DlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule

Appears on:[DlmLifecyclePolicySpecPolicyDetailsSchedule](#dlmlifecyclepolicyspecpolicydetailsschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `interval` | ***int64***||
| `intervalUnit` | ***string***| ***(Optional)*** |
| `times` | ***[]string***| ***(Optional)*** |
## DlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule

Appears on:[DlmLifecyclePolicySpecPolicyDetailsSchedule](#dlmlifecyclepolicyspecpolicydetailsschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***||
## DlmLifecyclePolicyStatus

Appears on:[DlmLifecyclePolicy](#dlmlifecyclepolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DlmLifecyclePolicySpec](#dlmlifecyclepolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DlmLifecyclePolicyStatus](#dlmlifecyclepolicystatus)

---
