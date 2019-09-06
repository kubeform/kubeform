## VpnGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpnGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpnGatewaySpec](#VpnGatewaySpec)***||
| `status` | ***[VpnGatewayStatus](#VpnGatewayStatus)***||
## VpnGatewaySpec
##### (Appears on:[VpnGateway](#VpnGateway), [VpnGatewayStatus](#VpnGatewayStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `amazonSideAsn` | ***string***| ***(Optional)*** |
| `availabilityZone` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## VpnGatewayStatus
##### (Appears on:[VpnGateway](#VpnGateway))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpnGatewaySpec](#VpnGatewaySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
