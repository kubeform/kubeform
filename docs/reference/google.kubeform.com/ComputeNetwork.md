## ComputeNetwork
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeNetwork` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeNetworkSpec](#ComputeNetworkSpec)***||
| `status` | ***[ComputeNetworkStatus](#ComputeNetworkStatus)***||
## ComputeNetworkSpec
##### (Appears on:[ComputeNetwork](#ComputeNetwork), [ComputeNetworkStatus](#ComputeNetworkStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoCreateSubnetworks` | ***bool***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `gatewayIpv4` | ***string***| ***(Optional)*** |
| `ipv4Range` | ***string***| ***(Optional)*** Deprecated|
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `routingMode` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeNetworkStatus
##### (Appears on:[ComputeNetwork](#ComputeNetwork))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeNetworkSpec](#ComputeNetworkSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
