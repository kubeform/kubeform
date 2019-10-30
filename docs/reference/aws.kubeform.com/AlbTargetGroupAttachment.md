---
title: AlbTargetGroupAttachment
menu:
  docs_{{ .version }}:
    identifier: albtargetgroupattachment-aws.kubeform.com
    name: AlbTargetGroupAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AlbTargetGroupAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AlbTargetGroupAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AlbTargetGroupAttachmentSpec](#albtargetgroupattachmentspec)***||
| `status` | ***[AlbTargetGroupAttachmentStatus](#albtargetgroupattachmentstatus)***||
## AlbTargetGroupAttachmentSpec

Appears on:[AlbTargetGroupAttachment](#albtargetgroupattachment), [AlbTargetGroupAttachmentStatus](#albtargetgroupattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availabilityZone` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `targetGroupArn` | ***string***||
| `targetID` | ***string***||
## AlbTargetGroupAttachmentStatus

Appears on:[AlbTargetGroupAttachment](#albtargetgroupattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AlbTargetGroupAttachmentSpec](#albtargetgroupattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AlbTargetGroupAttachmentStatus](#albtargetgroupattachmentstatus)

---
