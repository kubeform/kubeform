---
title: EbsSnapshot
menu:
  docs_{{ .version }}:
    identifier: ebssnapshot-aws.kubeform.com
    name: EbsSnapshot
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EbsSnapshot
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EbsSnapshot` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EbsSnapshotSpec](#EbsSnapshotSpec)***||
| `status` | ***[EbsSnapshotStatus](#EbsSnapshotStatus)***||
## EbsSnapshotSpec

Appears on:[EbsSnapshot](#EbsSnapshot), [EbsSnapshotStatus](#EbsSnapshotStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dataEncryptionKeyID` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `ownerAlias` | ***string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `volumeID` | ***string***||
| `volumeSize` | ***int***| ***(Optional)*** |
## EbsSnapshotStatus

Appears on:[EbsSnapshot](#EbsSnapshot)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EbsSnapshotSpec](#EbsSnapshotSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
