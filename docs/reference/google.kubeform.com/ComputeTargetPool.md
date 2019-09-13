---
title: ComputeTargetPool
menu:
  docs_{{ .version }}:
    identifier: computetargetpool-google.kubeform.com
    name: ComputeTargetPool
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeTargetPool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeTargetPool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeTargetPoolSpec](#ComputeTargetPoolSpec)***||
| `status` | ***[ComputeTargetPoolStatus](#ComputeTargetPoolStatus)***||
## ComputeTargetPoolSpec
##### (Appears on:[ComputeTargetPool](#ComputeTargetPool), [ComputeTargetPoolStatus](#ComputeTargetPoolStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backupPool` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `failoverRatio` | ***encoding/json.Number***| ***(Optional)*** |
| `healthChecks` | ***[]string***| ***(Optional)*** |
| `instances` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sessionAffinity` | ***string***| ***(Optional)*** |
## ComputeTargetPoolStatus
##### (Appears on:[ComputeTargetPool](#ComputeTargetPool))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeTargetPoolSpec](#ComputeTargetPoolSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
