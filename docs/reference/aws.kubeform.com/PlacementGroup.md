---
title: PlacementGroup
menu:
  docs_v0.0.1:
    identifier: placementgroup-aws.kubeform.com
    name: PlacementGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## PlacementGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PlacementGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PlacementGroupSpec](#PlacementGroupSpec)***||
| `status` | ***[PlacementGroupStatus](#PlacementGroupStatus)***||
## PlacementGroupSpec
##### (Appears on:[PlacementGroup](#PlacementGroup), [PlacementGroupStatus](#PlacementGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `strategy` | ***string***||
## PlacementGroupStatus
##### (Appears on:[PlacementGroup](#PlacementGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PlacementGroupSpec](#PlacementGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
