---
title: BackupPlan
menu:
  docs_v0.0.1:
    identifier: backupplan-aws.kubeform.com
    name: BackupPlan
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## BackupPlan
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `BackupPlan` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BackupPlanSpec](#BackupPlanSpec)***||
| `status` | ***[BackupPlanStatus](#BackupPlanStatus)***||
## BackupPlanSpec
##### (Appears on:[BackupPlan](#BackupPlan), [BackupPlanStatus](#BackupPlanStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `rule` | ***[[]BackupPlanSpecRule](#BackupPlanSpecRule)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## BackupPlanSpecRule
##### (Appears on:[BackupPlanSpec](#BackupPlanSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `completionWindow` | ***int***| ***(Optional)*** |
| `lifecycle` | ***[[]BackupPlanSpecRuleLifecycle](#BackupPlanSpecRuleLifecycle)***| ***(Optional)*** |
| `recoveryPointTags` | ***map[string]string***| ***(Optional)*** |
| `ruleName` | ***string***||
| `schedule` | ***string***| ***(Optional)*** |
| `startWindow` | ***int***| ***(Optional)*** |
| `targetVaultName` | ***string***||
## BackupPlanSpecRuleLifecycle
##### (Appears on:[BackupPlanSpecRule](#BackupPlanSpecRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `coldStorageAfter` | ***int***| ***(Optional)*** |
| `deleteAfter` | ***int***| ***(Optional)*** |
## BackupPlanStatus
##### (Appears on:[BackupPlan](#BackupPlan))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BackupPlanSpec](#BackupPlanSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
