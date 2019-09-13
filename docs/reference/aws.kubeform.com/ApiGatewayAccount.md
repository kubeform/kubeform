---
title: ApiGatewayAccount
menu:
  docs_{{ .version }}:
    identifier: apigatewayaccount-aws.kubeform.com
    name: ApiGatewayAccount
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayAccountSpec](#ApiGatewayAccountSpec)***||
| `status` | ***[ApiGatewayAccountStatus](#ApiGatewayAccountStatus)***||
## ApiGatewayAccountSpec
##### (Appears on:[ApiGatewayAccount](#ApiGatewayAccount), [ApiGatewayAccountStatus](#ApiGatewayAccountStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cloudwatchRoleArn` | ***string***| ***(Optional)*** |
| `throttleSettings` | ***[[]ApiGatewayAccountSpecThrottleSettings](#ApiGatewayAccountSpecThrottleSettings)***| ***(Optional)*** |
## ApiGatewayAccountSpecThrottleSettings
##### (Appears on:[ApiGatewayAccountSpec](#ApiGatewayAccountSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `burstLimit` | ***int***| ***(Optional)*** |
| `rateLimit` | ***encoding/json.Number***| ***(Optional)*** |
## ApiGatewayAccountStatus
##### (Appears on:[ApiGatewayAccount](#ApiGatewayAccount))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayAccountSpec](#ApiGatewayAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
