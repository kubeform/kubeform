---
title: ComputeNetworkEndpoint
menu:
  docs_{{ .version }}:
    identifier: computenetworkendpoint-google.kubeform.com
    name: ComputeNetworkEndpoint
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeNetworkEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeNetworkEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeNetworkEndpointSpec](#computenetworkendpointspec)***||
| `status` | ***[ComputeNetworkEndpointStatus](#computenetworkendpointstatus)***||
## ComputeNetworkEndpointSpec

Appears on:[ComputeNetworkEndpoint](#computenetworkendpoint), [ComputeNetworkEndpointStatus](#computenetworkendpointstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `instance` | ***string***||
| `ipAddress` | ***string***||
| `networkEndpointGroup` | ***string***||
| `port` | ***int64***||
| `project` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComputeNetworkEndpointStatus

Appears on:[ComputeNetworkEndpoint](#computenetworkendpoint)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeNetworkEndpointSpec](#computenetworkendpointspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeNetworkEndpointStatus](#computenetworkendpointstatus)

---
