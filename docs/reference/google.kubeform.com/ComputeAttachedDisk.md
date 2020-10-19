---
title: ComputeAttachedDisk
menu:
  docs_{{ .version }}:
    identifier: computeattacheddisk-google.kubeform.com
    name: ComputeAttachedDisk
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeAttachedDisk
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeAttachedDisk` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeAttachedDiskSpec](#computeattacheddiskspec)***||
| `status` | ***[ComputeAttachedDiskStatus](#computeattacheddiskstatus)***||
## ComputeAttachedDiskSpec

Appears on:[ComputeAttachedDisk](#computeattacheddisk), [ComputeAttachedDiskStatus](#computeattacheddiskstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `deviceName` | ***string***| ***(Optional)*** |
| `disk` | ***string***||
| `instance` | ***string***||
| `mode` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeAttachedDiskStatus

Appears on:[ComputeAttachedDisk](#computeattacheddisk)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeAttachedDiskSpec](#computeattacheddiskspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeAttachedDiskStatus](#computeattacheddiskstatus)

---
