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
| `spec` | ***[ApiGatewayAccountSpec](#apigatewayaccountspec)***||
| `status` | ***[ApiGatewayAccountStatus](#apigatewayaccountstatus)***||
## ApiGatewayAccountSpec

Appears on:[ApiGatewayAccount](#apigatewayaccount), [ApiGatewayAccountStatus](#apigatewayaccountstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cloudwatchRoleArn` | ***string***| ***(Optional)*** |
| `throttleSettings` | ***[[]ApiGatewayAccountSpecThrottleSettings](#apigatewayaccountspecthrottlesettings)***| ***(Optional)*** |
## ApiGatewayAccountSpecThrottleSettings

Appears on:[ApiGatewayAccountSpec](#apigatewayaccountspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `burstLimit` | ***int***| ***(Optional)*** |
| `rateLimit` | ***encoding/json.Number***| ***(Optional)*** |
## ApiGatewayAccountStatus

Appears on:[ApiGatewayAccount](#apigatewayaccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayAccountSpec](#apigatewayaccountspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiGatewayAccountStatus](#apigatewayaccountstatus)

---
