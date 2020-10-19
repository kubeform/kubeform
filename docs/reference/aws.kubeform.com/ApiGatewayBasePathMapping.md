---
title: ApiGatewayBasePathMapping
menu:
  docs_{{ .version }}:
    identifier: apigatewaybasepathmapping-aws.kubeform.com
    name: ApiGatewayBasePathMapping
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayBasePathMapping
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayBasePathMapping` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayBasePathMappingSpec](#apigatewaybasepathmappingspec)***||
| `status` | ***[ApiGatewayBasePathMappingStatus](#apigatewaybasepathmappingstatus)***||
## ApiGatewayBasePathMappingSpec

Appears on:[ApiGatewayBasePathMapping](#apigatewaybasepathmapping), [ApiGatewayBasePathMappingStatus](#apigatewaybasepathmappingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiID` | ***string***||
| `basePath` | ***string***| ***(Optional)*** |
| `domainName` | ***string***||
| `stageName` | ***string***| ***(Optional)*** |
## ApiGatewayBasePathMappingStatus

Appears on:[ApiGatewayBasePathMapping](#apigatewaybasepathmapping)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayBasePathMappingSpec](#apigatewaybasepathmappingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiGatewayBasePathMappingStatus](#apigatewaybasepathmappingstatus)

---
