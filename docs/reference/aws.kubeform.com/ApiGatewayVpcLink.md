## ApiGatewayVpcLink
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayVpcLink` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayVpcLinkSpec](#ApiGatewayVpcLinkSpec)***||
| `status` | ***[ApiGatewayVpcLinkStatus](#ApiGatewayVpcLinkStatus)***||
## ApiGatewayVpcLinkSpec
##### (Appears on:[ApiGatewayVpcLink](#ApiGatewayVpcLink), [ApiGatewayVpcLinkStatus](#ApiGatewayVpcLinkStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `targetArns` | ***[]string***||
## ApiGatewayVpcLinkStatus
##### (Appears on:[ApiGatewayVpcLink](#ApiGatewayVpcLink))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayVpcLinkSpec](#ApiGatewayVpcLinkSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
