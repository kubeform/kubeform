---
title: BackupSelection
menu:
  docs_v0.0.1:
    identifier: backupselection-aws.kubeform.com
    name: BackupSelection
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## BackupSelection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `BackupSelection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BackupSelectionSpec](#BackupSelectionSpec)***||
| `status` | ***[BackupSelectionStatus](#BackupSelectionStatus)***||
## BackupSelectionSpec
##### (Appears on:[BackupSelection](#BackupSelection), [BackupSelectionStatus](#BackupSelectionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `iamRoleArn` | ***string***||
| `name` | ***string***||
| `planID` | ***string***||
| `resources` | ***[]string***| ***(Optional)*** |
| `selectionTag` | ***[[]BackupSelectionSpecSelectionTag](#BackupSelectionSpecSelectionTag)***| ***(Optional)*** |
## BackupSelectionSpecSelectionTag
##### (Appears on:[BackupSelectionSpec](#BackupSelectionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `type` | ***string***||
| `value` | ***string***||
## BackupSelectionStatus
##### (Appears on:[BackupSelection](#BackupSelection))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BackupSelectionSpec](#BackupSelectionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
