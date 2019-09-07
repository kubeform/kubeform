---
title: ComputeProjectMetadataItem
menu:
  docs_v0.0.1:
    identifier: computeprojectmetadataitem-google.kubeform.com
    name: ComputeProjectMetadataItem
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ComputeProjectMetadataItem
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeProjectMetadataItem` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeProjectMetadataItemSpec](#ComputeProjectMetadataItemSpec)***||
| `status` | ***[ComputeProjectMetadataItemStatus](#ComputeProjectMetadataItemStatus)***||
## ComputeProjectMetadataItemSpec
##### (Appears on:[ComputeProjectMetadataItem](#ComputeProjectMetadataItem), [ComputeProjectMetadataItemStatus](#ComputeProjectMetadataItemStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `key` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `value` | ***string***||
## ComputeProjectMetadataItemStatus
##### (Appears on:[ComputeProjectMetadataItem](#ComputeProjectMetadataItem))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeProjectMetadataItemSpec](#ComputeProjectMetadataItemSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
