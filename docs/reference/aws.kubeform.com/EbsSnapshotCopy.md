---
title: EbsSnapshotCopy
menu:
  docs_{{ .version }}:
    identifier: ebssnapshotcopy-aws.kubeform.com
    name: EbsSnapshotCopy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EbsSnapshotCopy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EbsSnapshotCopy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EbsSnapshotCopySpec](#ebssnapshotcopyspec)***||
| `status` | ***[EbsSnapshotCopyStatus](#ebssnapshotcopystatus)***||
## EbsSnapshotCopySpec

Appears on:[EbsSnapshotCopy](#ebssnapshotcopy), [EbsSnapshotCopyStatus](#ebssnapshotcopystatus)

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
| `sourceRegion` | ***string***||
| `sourceSnapshotID` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `volumeID` | ***string***| ***(Optional)*** |
| `volumeSize` | ***int***| ***(Optional)*** |
## EbsSnapshotCopyStatus

Appears on:[EbsSnapshotCopy](#ebssnapshotcopy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EbsSnapshotCopySpec](#ebssnapshotcopyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EbsSnapshotCopyStatus](#ebssnapshotcopystatus)

---
