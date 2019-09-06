## ComputeRouterPeer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRouterPeer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRouterPeerSpec](#ComputeRouterPeerSpec)***||
| `status` | ***[ComputeRouterPeerStatus](#ComputeRouterPeerStatus)***||
## ComputeRouterPeerSpec
##### (Appears on:[ComputeRouterPeer](#ComputeRouterPeer), [ComputeRouterPeerStatus](#ComputeRouterPeerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `advertisedRoutePriority` | ***int***| ***(Optional)*** |
| `interface` | ***string***||
| `ipAddress` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `peerAsn` | ***int***||
| `peerIPAddress` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `router` | ***string***||
## ComputeRouterPeerStatus
##### (Appears on:[ComputeRouterPeer](#ComputeRouterPeer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRouterPeerSpec](#ComputeRouterPeerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
