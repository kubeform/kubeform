---
title: DropletSnapshot
menu:
  docs_{{ .version }}:
    identifier: dropletsnapshot-digitalocean.kubeform.com
    name: DropletSnapshot
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DropletSnapshot
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `DropletSnapshot` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DropletSnapshotSpec](#DropletSnapshotSpec)***||
| `status` | ***[DropletSnapshotStatus](#DropletSnapshotStatus)***||
## DropletSnapshotSpec

Appears on:[DropletSnapshot](#DropletSnapshot), [DropletSnapshotStatus](#DropletSnapshotStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createdAt` | ***string***| ***(Optional)*** |
| `dropletID` | ***string***||
| `minDiskSize` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `regions` | ***[]string***| ***(Optional)*** |
| `size` | ***encoding/json.Number***| ***(Optional)*** |
## DropletSnapshotStatus

Appears on:[DropletSnapshot](#DropletSnapshot)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DropletSnapshotSpec](#DropletSnapshotSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
