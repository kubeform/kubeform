---
title: LbTargetGroupAttachment
menu:
  docs_v0.0.1:
    identifier: lbtargetgroupattachment-aws.kubeform.com
    name: LbTargetGroupAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## LbTargetGroupAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbTargetGroupAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbTargetGroupAttachmentSpec](#LbTargetGroupAttachmentSpec)***||
| `status` | ***[LbTargetGroupAttachmentStatus](#LbTargetGroupAttachmentStatus)***||
## LbTargetGroupAttachmentSpec
##### (Appears on:[LbTargetGroupAttachment](#LbTargetGroupAttachment), [LbTargetGroupAttachmentStatus](#LbTargetGroupAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availabilityZone` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `targetGroupArn` | ***string***||
| `targetID` | ***string***||
## LbTargetGroupAttachmentStatus
##### (Appears on:[LbTargetGroupAttachment](#LbTargetGroupAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbTargetGroupAttachmentSpec](#LbTargetGroupAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
