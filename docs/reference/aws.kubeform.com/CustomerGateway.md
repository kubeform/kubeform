## CustomerGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CustomerGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CustomerGatewaySpec](#CustomerGatewaySpec)***||
| `status` | ***[CustomerGatewayStatus](#CustomerGatewayStatus)***||
## CustomerGatewaySpec
##### (Appears on:[CustomerGateway](#CustomerGateway), [CustomerGatewayStatus](#CustomerGatewayStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bgpAsn` | ***int***||
| `ipAddress` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***||
## CustomerGatewayStatus
##### (Appears on:[CustomerGateway](#CustomerGateway))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CustomerGatewaySpec](#CustomerGatewaySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
