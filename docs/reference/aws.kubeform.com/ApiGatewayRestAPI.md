---
title: ApiGatewayRestAPI
menu:
  docs_v0.0.1:
    identifier: apigatewayrestapi-aws.kubeform.com
    name: ApiGatewayRestAPI
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApiGatewayRestAPI
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayRestAPI` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayRestAPISpec](#ApiGatewayRestAPISpec)***||
| `status` | ***[ApiGatewayRestAPIStatus](#ApiGatewayRestAPIStatus)***||
## ApiGatewayRestAPISpec
##### (Appears on:[ApiGatewayRestAPI](#ApiGatewayRestAPI), [ApiGatewayRestAPIStatus](#ApiGatewayRestAPIStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiKeySource` | ***string***| ***(Optional)*** |
| `binaryMediaTypes` | ***[]string***| ***(Optional)*** |
| `body` | ***string***| ***(Optional)*** |
| `createdDate` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `endpointConfiguration` | ***[[]ApiGatewayRestAPISpecEndpointConfiguration](#ApiGatewayRestAPISpecEndpointConfiguration)***| ***(Optional)*** |
| `executionArn` | ***string***| ***(Optional)*** |
| `minimumCompressionSize` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `policy` | ***string***| ***(Optional)*** |
| `rootResourceID` | ***string***| ***(Optional)*** |
## ApiGatewayRestAPISpecEndpointConfiguration
##### (Appears on:[ApiGatewayRestAPISpec](#ApiGatewayRestAPISpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `types` | ***[]string***||
## ApiGatewayRestAPIStatus
##### (Appears on:[ApiGatewayRestAPI](#ApiGatewayRestAPI))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayRestAPISpec](#ApiGatewayRestAPISpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
