---
title: ComputeProjectMetadata
menu:
  docs_{{ .version }}:
    identifier: computeprojectmetadata-google.kubeform.com
    name: ComputeProjectMetadata
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeProjectMetadata
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeProjectMetadata` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeProjectMetadataSpec](#computeprojectmetadataspec)***||
| `status` | ***[ComputeProjectMetadataStatus](#computeprojectmetadatastatus)***||
## ComputeProjectMetadataSpec

Appears on:[ComputeProjectMetadata](#computeprojectmetadata), [ComputeProjectMetadataStatus](#computeprojectmetadatastatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `metadata` | ***map[string]string***||
| `project` | ***string***| ***(Optional)*** |
## ComputeProjectMetadataStatus

Appears on:[ComputeProjectMetadata](#computeprojectmetadata)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeProjectMetadataSpec](#computeprojectmetadataspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeProjectMetadataStatus](#computeprojectmetadatastatus)

---
