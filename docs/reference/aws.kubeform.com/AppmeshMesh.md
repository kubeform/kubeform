---
title: AppmeshMesh
menu:
  docs_{{ .version }}:
    identifier: appmeshmesh-aws.kubeform.com
    name: AppmeshMesh
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppmeshMesh
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppmeshMesh` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppmeshMeshSpec](#AppmeshMeshSpec)***||
| `status` | ***[AppmeshMeshStatus](#AppmeshMeshStatus)***||
## AppmeshMeshSpec
##### (Appears on:[AppmeshMesh](#AppmeshMesh), [AppmeshMeshStatus](#AppmeshMeshStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `createdDate` | ***string***| ***(Optional)*** |
| `lastUpdatedDate` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `spec` | ***[[]AppmeshMeshSpecSpec](#AppmeshMeshSpecSpec)***| ***(Optional)*** |
## AppmeshMeshSpecSpec
##### (Appears on:[AppmeshMeshSpec](#AppmeshMeshSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `egressFilter` | ***[[]AppmeshMeshSpecSpecEgressFilter](#AppmeshMeshSpecSpecEgressFilter)***| ***(Optional)*** |
## AppmeshMeshSpecSpecEgressFilter
##### (Appears on:[AppmeshMeshSpecSpec](#AppmeshMeshSpecSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***| ***(Optional)*** |
## AppmeshMeshStatus
##### (Appears on:[AppmeshMesh](#AppmeshMesh))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppmeshMeshSpec](#AppmeshMeshSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
