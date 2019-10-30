---
title: ComputeNetworkPeering
menu:
  docs_{{ .version }}:
    identifier: computenetworkpeering-google.kubeform.com
    name: ComputeNetworkPeering
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeNetworkPeering
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeNetworkPeering` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeNetworkPeeringSpec](#computenetworkpeeringspec)***||
| `status` | ***[ComputeNetworkPeeringStatus](#computenetworkpeeringstatus)***||
## ComputeNetworkPeeringSpec

Appears on:[ComputeNetworkPeering](#computenetworkpeering), [ComputeNetworkPeeringStatus](#computenetworkpeeringstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoCreateRoutes` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `network` | ***string***||
| `peerNetwork` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `stateDetails` | ***string***| ***(Optional)*** |
## ComputeNetworkPeeringStatus

Appears on:[ComputeNetworkPeering](#computenetworkpeering)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeNetworkPeeringSpec](#computenetworkpeeringspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeNetworkPeeringStatus](#computenetworkpeeringstatus)

---
