## ApiGatewayBasePathMapping
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayBasePathMapping` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayBasePathMappingSpec](#ApiGatewayBasePathMappingSpec)***||
| `status` | ***[ApiGatewayBasePathMappingStatus](#ApiGatewayBasePathMappingStatus)***||
## ApiGatewayBasePathMappingSpec
##### (Appears on:[ApiGatewayBasePathMapping](#ApiGatewayBasePathMapping), [ApiGatewayBasePathMappingStatus](#ApiGatewayBasePathMappingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiID` | ***string***||
| `basePath` | ***string***| ***(Optional)*** |
| `domainName` | ***string***||
| `stageName` | ***string***| ***(Optional)*** |
## ApiGatewayBasePathMappingStatus
##### (Appears on:[ApiGatewayBasePathMapping](#ApiGatewayBasePathMapping))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayBasePathMappingSpec](#ApiGatewayBasePathMappingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
