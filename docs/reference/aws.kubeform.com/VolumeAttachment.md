---
title: VolumeAttachment
menu:
  docs_{{ .version }}:
    identifier: volumeattachment-aws.kubeform.com
    name: VolumeAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VolumeAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VolumeAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VolumeAttachmentSpec](#VolumeAttachmentSpec)***||
| `status` | ***[VolumeAttachmentStatus](#VolumeAttachmentStatus)***||
## VolumeAttachmentSpec
##### (Appears on:[VolumeAttachment](#VolumeAttachment), [VolumeAttachmentStatus](#VolumeAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `deviceName` | ***string***||
| `forceDetach` | ***bool***| ***(Optional)*** |
| `instanceID` | ***string***||
| `skipDestroy` | ***bool***| ***(Optional)*** |
| `volumeID` | ***string***||
## VolumeAttachmentStatus
##### (Appears on:[VolumeAttachment](#VolumeAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VolumeAttachmentSpec](#VolumeAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
