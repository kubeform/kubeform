---
title: LbTargetGroupAttachment
menu:
  docs_{{ .version }}:
    identifier: lbtargetgroupattachment-aws.kubeform.com
    name: LbTargetGroupAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LbTargetGroupAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbTargetGroupAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbTargetGroupAttachmentSpec](#lbtargetgroupattachmentspec)***||
| `status` | ***[LbTargetGroupAttachmentStatus](#lbtargetgroupattachmentstatus)***||
## LbTargetGroupAttachmentSpec

Appears on:[LbTargetGroupAttachment](#lbtargetgroupattachment), [LbTargetGroupAttachmentStatus](#lbtargetgroupattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availabilityZone` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `targetGroupArn` | ***string***||
| `targetID` | ***string***||
## LbTargetGroupAttachmentStatus

Appears on:[LbTargetGroupAttachment](#lbtargetgroupattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbTargetGroupAttachmentSpec](#lbtargetgroupattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LbTargetGroupAttachmentStatus](#lbtargetgroupattachmentstatus)

---
