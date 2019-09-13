---
title: AutoscalingAttachment
menu:
  docs_{{ .version }}:
    identifier: autoscalingattachment-aws.kubeform.com
    name: AutoscalingAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutoscalingAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AutoscalingAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutoscalingAttachmentSpec](#AutoscalingAttachmentSpec)***||
| `status` | ***[AutoscalingAttachmentStatus](#AutoscalingAttachmentStatus)***||
## AutoscalingAttachmentSpec
##### (Appears on:[AutoscalingAttachment](#AutoscalingAttachment), [AutoscalingAttachmentStatus](#AutoscalingAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `albTargetGroupArn` | ***string***| ***(Optional)*** |
| `autoscalingGroupName` | ***string***||
| `elb` | ***string***| ***(Optional)*** |
## AutoscalingAttachmentStatus
##### (Appears on:[AutoscalingAttachment](#AutoscalingAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutoscalingAttachmentSpec](#AutoscalingAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
