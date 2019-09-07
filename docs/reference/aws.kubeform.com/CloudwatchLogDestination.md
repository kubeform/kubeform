---
title: CloudwatchLogDestination
menu:
  docs_v0.0.1:
    identifier: cloudwatchlogdestination-aws.kubeform.com
    name: CloudwatchLogDestination
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## CloudwatchLogDestination
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchLogDestination` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchLogDestinationSpec](#CloudwatchLogDestinationSpec)***||
| `status` | ***[CloudwatchLogDestinationStatus](#CloudwatchLogDestinationStatus)***||
## CloudwatchLogDestinationSpec
##### (Appears on:[CloudwatchLogDestination](#CloudwatchLogDestination), [CloudwatchLogDestinationStatus](#CloudwatchLogDestinationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `roleArn` | ***string***||
| `targetArn` | ***string***||
## CloudwatchLogDestinationStatus
##### (Appears on:[CloudwatchLogDestination](#CloudwatchLogDestination))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchLogDestinationSpec](#CloudwatchLogDestinationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
