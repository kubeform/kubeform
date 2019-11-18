---
title: VolumeAttachment
menu:
  docs_{{ .version }}:
    identifier: volumeattachment-digitalocean.kubeform.com
    name: VolumeAttachment
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VolumeAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `VolumeAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VolumeAttachmentSpec](#volumeattachmentspec)***||
| `status` | ***[VolumeAttachmentStatus](#volumeattachmentstatus)***||
## Phase(`string` alias)

Appears on:[VolumeAttachmentStatus](#volumeattachmentstatus)

## VolumeAttachmentSpec

Appears on:[VolumeAttachment](#volumeattachment), [VolumeAttachmentStatus](#volumeattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dropletID` | ***int64***||
| `volumeID` | ***string***||
## VolumeAttachmentStatus

Appears on:[VolumeAttachment](#volumeattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VolumeAttachmentSpec](#volumeattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
