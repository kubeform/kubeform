---
title: VolumeAttachment
menu:
  docs_v0.0.1:
    identifier: volumeattachment-digitalocean.kubeform.com
    name: VolumeAttachment
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## VolumeAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
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
| `dropletID` | ***int***||
| `volumeID` | ***string***||
## VolumeAttachmentStatus
##### (Appears on:[VolumeAttachment](#VolumeAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VolumeAttachmentSpec](#VolumeAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
