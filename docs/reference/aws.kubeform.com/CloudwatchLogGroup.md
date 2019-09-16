---
title: CloudwatchLogGroup
menu:
  docs_{{ .version }}:
    identifier: cloudwatchloggroup-aws.kubeform.com
    name: CloudwatchLogGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchLogGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchLogGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchLogGroupSpec](#cloudwatchloggroupspec)***||
| `status` | ***[CloudwatchLogGroupStatus](#cloudwatchloggroupstatus)***||
## CloudwatchLogGroupSpec

Appears on:[CloudwatchLogGroup](#cloudwatchloggroup), [CloudwatchLogGroupStatus](#cloudwatchloggroupstatus)

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

Appears on:[CloudwatchLogGroup](#cloudwatchloggroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchLogGroupSpec](#cloudwatchloggroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
