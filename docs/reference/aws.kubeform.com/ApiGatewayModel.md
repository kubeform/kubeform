---
title: ApiGatewayModel
menu:
  docs_{{ .version }}:
    identifier: apigatewaymodel-aws.kubeform.com
    name: ApiGatewayModel
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayModel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayModel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayModelSpec](#ApiGatewayModelSpec)***||
| `status` | ***[ApiGatewayModelStatus](#ApiGatewayModelStatus)***||
## ApiGatewayModelSpec
##### (Appears on:[ApiGatewayModel](#ApiGatewayModel), [ApiGatewayModelStatus](#ApiGatewayModelStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `contentType` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `restAPIID` | ***string***||
| `schema` | ***string***| ***(Optional)*** |
## ApiGatewayModelStatus
##### (Appears on:[ApiGatewayModel](#ApiGatewayModel))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayModelSpec](#ApiGatewayModelSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
