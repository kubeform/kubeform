---
title: ApiGatewayIntegrationResponse
menu:
  docs_{{ .version }}:
    identifier: apigatewayintegrationresponse-aws.kubeform.com
    name: ApiGatewayIntegrationResponse
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayIntegrationResponse
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayIntegrationResponse` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayIntegrationResponseSpec](#apigatewayintegrationresponsespec)***||
| `status` | ***[ApiGatewayIntegrationResponseStatus](#apigatewayintegrationresponsestatus)***||
## ApiGatewayIntegrationResponseSpec

Appears on:[ApiGatewayIntegrationResponse](#apigatewayintegrationresponse), [ApiGatewayIntegrationResponseStatus](#apigatewayintegrationresponsestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `contentHandling` | ***string***| ***(Optional)*** |
| `httpMethod` | ***string***||
| `resourceID` | ***string***||
| `responseParameters` | ***map[string]string***| ***(Optional)*** |
| `responseTemplates` | ***map[string]string***| ***(Optional)*** |
| `restAPIID` | ***string***||
| `selectionPattern` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## ApiGatewayIntegrationResponseStatus

Appears on:[ApiGatewayIntegrationResponse](#apigatewayintegrationresponse)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayIntegrationResponseSpec](#apigatewayintegrationresponsespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiGatewayIntegrationResponseStatus](#apigatewayintegrationresponsestatus)

---
