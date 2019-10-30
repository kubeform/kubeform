---
title: VolumeSnapshot
menu:
  docs_{{ .version }}:
    identifier: volumesnapshot-digitalocean.kubeform.com
    name: VolumeSnapshot
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VolumeSnapshot
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `VolumeSnapshot` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VolumeSnapshotSpec](#volumesnapshotspec)***||
| `status` | ***[VolumeSnapshotStatus](#volumesnapshotstatus)***||
## Phase(`string` alias)

Appears on:[VolumeSnapshotStatus](#volumesnapshotstatus)

## VolumeSnapshotSpec

Appears on:[VolumeSnapshot](#volumesnapshot), [VolumeSnapshotStatus](#volumesnapshotstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createdAt` | ***string***| ***(Optional)*** |
| `minDiskSize` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `regions` | ***[]string***| ***(Optional)*** |
| `size` | ***encoding/json.Number***| ***(Optional)*** |
| `volumeID` | ***string***||
## VolumeSnapshotStatus

Appears on:[VolumeSnapshot](#volumesnapshot)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VolumeSnapshotSpec](#volumesnapshotspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
