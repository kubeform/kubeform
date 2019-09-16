---
title: ComputeInstanceGroup
menu:
  docs_{{ .version }}:
    identifier: computeinstancegroup-google.kubeform.com
    name: ComputeInstanceGroup
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeInstanceGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeInstanceGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeInstanceGroupSpec](#ComputeInstanceGroupSpec)***||
| `status` | ***[ComputeInstanceGroupStatus](#ComputeInstanceGroupStatus)***||
## ComputeInstanceGroupSpec

Appears on:[ComputeInstanceGroup](#ComputeInstanceGroup), [ComputeInstanceGroupStatus](#ComputeInstanceGroupStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `instances` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `namedPort` | ***[[]ComputeInstanceGroupSpecNamedPort](#ComputeInstanceGroupSpecNamedPort)***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `size` | ***int***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeInstanceGroupSpecNamedPort

Appears on:[ComputeInstanceGroupSpec](#ComputeInstanceGroupSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `port` | ***int***||
## ComputeInstanceGroupStatus

Appears on:[ComputeInstanceGroup](#ComputeInstanceGroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeInstanceGroupSpec](#ComputeInstanceGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
