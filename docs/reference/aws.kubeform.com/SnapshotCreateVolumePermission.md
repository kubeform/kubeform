---
title: SnapshotCreateVolumePermission
menu:
  docs_{{ .version }}:
    identifier: snapshotcreatevolumepermission-aws.kubeform.com
    name: SnapshotCreateVolumePermission
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SnapshotCreateVolumePermission
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SnapshotCreateVolumePermission` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SnapshotCreateVolumePermissionSpec](#snapshotcreatevolumepermissionspec)***||
| `status` | ***[SnapshotCreateVolumePermissionStatus](#snapshotcreatevolumepermissionstatus)***||
## Phase(`string` alias)

Appears on:[SnapshotCreateVolumePermissionStatus](#snapshotcreatevolumepermissionstatus)

## SnapshotCreateVolumePermissionSpec

Appears on:[SnapshotCreateVolumePermission](#snapshotcreatevolumepermission), [SnapshotCreateVolumePermissionStatus](#snapshotcreatevolumepermissionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***||
| `snapshotID` | ***string***||
## SnapshotCreateVolumePermissionStatus

Appears on:[SnapshotCreateVolumePermission](#snapshotcreatevolumepermission)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SnapshotCreateVolumePermissionSpec](#snapshotcreatevolumepermissionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
