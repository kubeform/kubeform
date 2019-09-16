---
title: ApiGatewayResource
menu:
  docs_{{ .version }}:
    identifier: apigatewayresource-aws.kubeform.com
    name: ApiGatewayResource
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayResource
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayResource` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayResourceSpec](#apigatewayresourcespec)***||
| `status` | ***[ApiGatewayResourceStatus](#apigatewayresourcestatus)***||
## ApiGatewayResourceSpec

Appears on:[ApiGatewayResource](#apigatewayresource), [ApiGatewayResourceStatus](#apigatewayresourcestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `parentID` | ***string***||
| `path` | ***string***| ***(Optional)*** |
| `pathPart` | ***string***||
| `restAPIID` | ***string***||
## ApiGatewayResourceStatus

Appears on:[ApiGatewayResource](#apigatewayresource)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayResourceSpec](#apigatewayresourcespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
