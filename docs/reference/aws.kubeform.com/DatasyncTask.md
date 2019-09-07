---
title: DatasyncTask
menu:
  docs_v0.0.1:
    identifier: datasynctask-aws.kubeform.com
    name: DatasyncTask
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DatasyncTask
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DatasyncTask` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatasyncTaskSpec](#DatasyncTaskSpec)***||
| `status` | ***[DatasyncTaskStatus](#DatasyncTaskStatus)***||
## DatasyncTaskSpec
##### (Appears on:[DatasyncTask](#DatasyncTask), [DatasyncTaskStatus](#DatasyncTaskStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `cloudwatchLogGroupArn` | ***string***| ***(Optional)*** |
| `destinationLocationArn` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `options` | ***[[]DatasyncTaskSpecOptions](#DatasyncTaskSpecOptions)***| ***(Optional)*** |
| `sourceLocationArn` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## DatasyncTaskSpecOptions
##### (Appears on:[DatasyncTaskSpec](#DatasyncTaskSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `atime` | ***string***| ***(Optional)*** |
| `bytesPerSecond` | ***int***| ***(Optional)*** |
| `gid` | ***string***| ***(Optional)*** |
| `mtime` | ***string***| ***(Optional)*** |
| `posixPermissions` | ***string***| ***(Optional)*** |
| `preserveDeletedFiles` | ***string***| ***(Optional)*** |
| `preserveDevices` | ***string***| ***(Optional)*** |
| `uid` | ***string***| ***(Optional)*** |
| `verifyMode` | ***string***| ***(Optional)*** |
## DatasyncTaskStatus
##### (Appears on:[DatasyncTask](#DatasyncTask))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatasyncTaskSpec](#DatasyncTaskSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
