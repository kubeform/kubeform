## ApiGatewayDocumentationPart
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayDocumentationPart` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayDocumentationPartSpec](#ApiGatewayDocumentationPartSpec)***||
| `status` | ***[ApiGatewayDocumentationPartStatus](#ApiGatewayDocumentationPartStatus)***||
## ApiGatewayDocumentationPartSpec
##### (Appears on:[ApiGatewayDocumentationPart](#ApiGatewayDocumentationPart), [ApiGatewayDocumentationPartStatus](#ApiGatewayDocumentationPartStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***[[]ApiGatewayDocumentationPartSpecLocation](#ApiGatewayDocumentationPartSpecLocation)***||
| `properties` | ***string***||
| `restAPIID` | ***string***||
## ApiGatewayDocumentationPartSpecLocation
##### (Appears on:[ApiGatewayDocumentationPartSpec](#ApiGatewayDocumentationPartSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `method` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## ApiGatewayDocumentationPartStatus
##### (Appears on:[ApiGatewayDocumentationPart](#ApiGatewayDocumentationPart))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayDocumentationPartSpec](#ApiGatewayDocumentationPartSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
