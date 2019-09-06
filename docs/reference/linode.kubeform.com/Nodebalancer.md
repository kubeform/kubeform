## Nodebalancer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `Nodebalancer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NodebalancerSpec](#NodebalancerSpec)***||
| `status` | ***[NodebalancerStatus](#NodebalancerStatus)***||
## NodebalancerSpec
##### (Appears on:[Nodebalancer](#Nodebalancer), [NodebalancerStatus](#NodebalancerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `clientConnThrottle` | ***int***| ***(Optional)*** Throttle connections per second (0-20). Set to 0 (zero) to disable throttling.|
| `created` | ***string***| ***(Optional)*** |
| `hostname` | ***string***| ***(Optional)*** This NodeBalancer's hostname, ending with .nodebalancer.linode.com|
| `ipv4` | ***string***| ***(Optional)*** The Public IPv4 Address of this NodeBalancer|
| `ipv6` | ***string***| ***(Optional)*** The Public IPv6 Address of this NodeBalancer|
| `label` | ***string***| ***(Optional)*** The label of the Linode NodeBalancer.|
| `region` | ***string***|The region where this NodeBalancer will be deployed.|
| `tags` | ***[]string***| ***(Optional)*** An array of tags applied to this object. Tags are for organizational purposes only.|
| `transfer` | ***map[string]kubeform.dev/kubeform/apis/linode/v1alpha1.NodebalancerSpecTransfer***| ***(Optional)*** |
| `updated` | ***string***| ***(Optional)*** |
## NodebalancerStatus
##### (Appears on:[Nodebalancer](#Nodebalancer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NodebalancerSpec](#NodebalancerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
