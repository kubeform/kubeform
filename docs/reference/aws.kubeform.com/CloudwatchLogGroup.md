---
title: CloudwatchLogGroup
menu:
  docs_v0.0.1:
    identifier: cloudwatchloggroup-aws.kubeform.com
    name: CloudwatchLogGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## CloudwatchLogGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchLogGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchLogGroupSpec](#CloudwatchLogGroupSpec)***||
| `status` | ***[CloudwatchLogGroupStatus](#CloudwatchLogGroupStatus)***||
## CloudwatchLogGroupSpec
##### (Appears on:[CloudwatchLogGroup](#CloudwatchLogGroup), [CloudwatchLogGroupStatus](#CloudwatchLogGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `retentionInDays` | ***int***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## CloudwatchLogGroupStatus
##### (Appears on:[CloudwatchLogGroup](#CloudwatchLogGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchLogGroupSpec](#CloudwatchLogGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
