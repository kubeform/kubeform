---
title: ComputeNetworkEndpointGroup
menu:
  docs_{{ .version }}:
    identifier: computenetworkendpointgroup-google.kubeform.com
    name: ComputeNetworkEndpointGroup
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeNetworkEndpointGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeNetworkEndpointGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeNetworkEndpointGroupSpec](#computenetworkendpointgroupspec)***||
| `status` | ***[ComputeNetworkEndpointGroupStatus](#computenetworkendpointgroupstatus)***||
## ComputeNetworkEndpointGroupSpec

Appears on:[ComputeNetworkEndpointGroup](#computenetworkendpointgroup), [ComputeNetworkEndpointGroupStatus](#computenetworkendpointgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `defaultPort` | ***int64***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `network` | ***string***||
| `networkEndpointType` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `size` | ***int64***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeNetworkEndpointGroupStatus

Appears on:[ComputeNetworkEndpointGroup](#computenetworkendpointgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeNetworkEndpointGroupSpec](#computenetworkendpointgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeNetworkEndpointGroupStatus](#computenetworkendpointgroupstatus)

---
