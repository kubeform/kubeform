---
title: DlmLifecyclePolicy
menu:
  docs_v0.0.1:
    identifier: dlmlifecyclepolicy-aws.kubeform.com
    name: DlmLifecyclePolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DlmLifecyclePolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DlmLifecyclePolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DlmLifecyclePolicySpec](#DlmLifecyclePolicySpec)***||
| `status` | ***[DlmLifecyclePolicyStatus](#DlmLifecyclePolicyStatus)***||
## DlmLifecyclePolicySpec
##### (Appears on:[DlmLifecyclePolicy](#DlmLifecyclePolicy), [DlmLifecyclePolicyStatus](#DlmLifecyclePolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***||
| `executionRoleArn` | ***string***||
| `policyDetails` | ***[[]DlmLifecyclePolicySpecPolicyDetails](#DlmLifecyclePolicySpecPolicyDetails)***||
| `state` | ***string***| ***(Optional)*** |
## DlmLifecyclePolicySpecPolicyDetails
##### (Appears on:[DlmLifecyclePolicySpec](#DlmLifecyclePolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceTypes` | ***[]string***||
| `schedule` | ***[[]DlmLifecyclePolicySpecPolicyDetailsSchedule](#DlmLifecyclePolicySpecPolicyDetailsSchedule)***||
| `targetTags` | ***map[string]string***||
## DlmLifecyclePolicySpecPolicyDetailsSchedule
##### (Appears on:[DlmLifecyclePolicySpecPolicyDetails](#DlmLifecyclePolicySpecPolicyDetails))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `copyTags` | ***bool***| ***(Optional)*** |
| `createRule` | ***[[]DlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule](#DlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule)***||
| `name` | ***string***||
| `retainRule` | ***[[]DlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule](#DlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule)***||
| `tagsToAdd` | ***map[string]string***| ***(Optional)*** |
## DlmLifecyclePolicySpecPolicyDetailsScheduleCreateRule
##### (Appears on:[DlmLifecyclePolicySpecPolicyDetailsSchedule](#DlmLifecyclePolicySpecPolicyDetailsSchedule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `interval` | ***int***||
| `intervalUnit` | ***string***| ***(Optional)*** |
| `times` | ***[]string***| ***(Optional)*** |
## DlmLifecyclePolicySpecPolicyDetailsScheduleRetainRule
##### (Appears on:[DlmLifecyclePolicySpecPolicyDetailsSchedule](#DlmLifecyclePolicySpecPolicyDetailsSchedule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int***||
## DlmLifecyclePolicyStatus
##### (Appears on:[DlmLifecyclePolicy](#DlmLifecyclePolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DlmLifecyclePolicySpec](#DlmLifecyclePolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
