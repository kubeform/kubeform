---
title: ApiGatewayMethod
menu:
  docs_{{ .version }}:
    identifier: apigatewaymethod-aws.kubeform.com
    name: ApiGatewayMethod
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayMethod
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayMethod` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayMethodSpec](#ApiGatewayMethodSpec)***||
| `status` | ***[ApiGatewayMethodStatus](#ApiGatewayMethodStatus)***||
## ApiGatewayMethodSpec
##### (Appears on:[ApiGatewayMethod](#ApiGatewayMethod), [ApiGatewayMethodStatus](#ApiGatewayMethodStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiKeyRequired` | ***bool***| ***(Optional)*** |
| `authorization` | ***string***||
| `authorizationScopes` | ***[]string***| ***(Optional)*** |
| `authorizerID` | ***string***| ***(Optional)*** |
| `httpMethod` | ***string***||
| `requestModels` | ***map[string]string***| ***(Optional)*** |
| `requestParameters` | ***map[string]bool***| ***(Optional)*** |
| `requestValidatorID` | ***string***| ***(Optional)*** |
| `resourceID` | ***string***||
| `restAPIID` | ***string***||
## ApiGatewayMethodStatus
##### (Appears on:[ApiGatewayMethod](#ApiGatewayMethod))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayMethodSpec](#ApiGatewayMethodSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
