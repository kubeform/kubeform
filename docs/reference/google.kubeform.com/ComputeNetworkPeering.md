## ComputeNetworkPeering
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeNetworkPeering` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeNetworkPeeringSpec](#ComputeNetworkPeeringSpec)***||
| `status` | ***[ComputeNetworkPeeringStatus](#ComputeNetworkPeeringStatus)***||
## ComputeNetworkPeeringSpec
##### (Appears on:[ComputeNetworkPeering](#ComputeNetworkPeering), [ComputeNetworkPeeringStatus](#ComputeNetworkPeeringStatus))
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
##### (Appears on:[ComputeNetworkPeering](#ComputeNetworkPeering))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeNetworkPeeringSpec](#ComputeNetworkPeeringSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
