---
title: ApiGatewayMethodResponse
menu:
  docs_{{ .version }}:
    identifier: apigatewaymethodresponse-aws.kubeform.com
    name: ApiGatewayMethodResponse
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayMethodResponse
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayMethodResponse` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayMethodResponseSpec](#apigatewaymethodresponsespec)***||
| `status` | ***[ApiGatewayMethodResponseStatus](#apigatewaymethodresponsestatus)***||
## ApiGatewayMethodResponseSpec

Appears on:[ApiGatewayMethodResponse](#apigatewaymethodresponse), [ApiGatewayMethodResponseStatus](#apigatewaymethodresponsestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `httpMethod` | ***string***||
| `resourceID` | ***string***||
| `responseModels` | ***map[string]string***| ***(Optional)*** |
| `responseParameters` | ***map[string]bool***| ***(Optional)*** |
| `restAPIID` | ***string***||
| `statusCode` | ***string***||
## ApiGatewayMethodResponseStatus

Appears on:[ApiGatewayMethodResponse](#apigatewaymethodresponse)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayMethodResponseSpec](#apigatewaymethodresponsespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiGatewayMethodResponseStatus](#apigatewaymethodresponsestatus)

---
