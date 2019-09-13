---
title: ComputeRegionDisk
menu:
  docs_{{ .version }}:
    identifier: computeregiondisk-google.kubeform.com
    name: ComputeRegionDisk
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeRegionDisk
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRegionDisk` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRegionDiskSpec](#ComputeRegionDiskSpec)***||
| `status` | ***[ComputeRegionDiskStatus](#ComputeRegionDiskStatus)***||
## ComputeRegionDiskSpec
##### (Appears on:[ComputeRegionDisk](#ComputeRegionDisk), [ComputeRegionDiskStatus](#ComputeRegionDiskStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `diskEncryptionKey` | ***[[]ComputeRegionDiskSpecDiskEncryptionKey](#ComputeRegionDiskSpecDiskEncryptionKey)***| ***(Optional)*** |
| `labelFingerprint` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `lastAttachTimestamp` | ***string***| ***(Optional)*** |
| `lastDetachTimestamp` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `replicaZones` | ***[]string***||
| `selfLink` | ***string***| ***(Optional)*** |
| `size` | ***int***| ***(Optional)*** |
| `snapshot` | ***string***| ***(Optional)*** |
| `sourceSnapshotEncryptionKey` | ***[[]ComputeRegionDiskSpecSourceSnapshotEncryptionKey](#ComputeRegionDiskSpecSourceSnapshotEncryptionKey)***| ***(Optional)*** |
| `sourceSnapshotID` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `users` | ***[]string***| ***(Optional)*** |
## ComputeRegionDiskSpecDiskEncryptionKey
##### (Appears on:[ComputeRegionDiskSpec](#ComputeRegionDiskSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `rawKey` | ***string***| ***(Optional)*** |
| `sha256` | ***string***| ***(Optional)*** |
## ComputeRegionDiskSpecSourceSnapshotEncryptionKey
##### (Appears on:[ComputeRegionDiskSpec](#ComputeRegionDiskSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `rawKey` | ***string***| ***(Optional)*** |
| `sha256` | ***string***| ***(Optional)*** |
## ComputeRegionDiskStatus
##### (Appears on:[ComputeRegionDisk](#ComputeRegionDisk))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRegionDiskSpec](#ComputeRegionDiskSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
