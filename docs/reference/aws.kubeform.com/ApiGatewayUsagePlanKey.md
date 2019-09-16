---
title: ApiGatewayUsagePlanKey
menu:
  docs_{{ .version }}:
    identifier: apigatewayusageplankey-aws.kubeform.com
    name: ApiGatewayUsagePlanKey
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayUsagePlanKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayUsagePlanKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayUsagePlanKeySpec](#ApiGatewayUsagePlanKeySpec)***||
| `status` | ***[ApiGatewayUsagePlanKeyStatus](#ApiGatewayUsagePlanKeyStatus)***||
## ApiGatewayUsagePlanKeySpec

Appears on:[ApiGatewayUsagePlanKey](#ApiGatewayUsagePlanKey), [ApiGatewayUsagePlanKeyStatus](#ApiGatewayUsagePlanKeyStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `keyID` | ***string***||
| `keyType` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `usagePlanID` | ***string***||
| `value` | ***string***| ***(Optional)*** |
## ApiGatewayUsagePlanKeyStatus

Appears on:[ApiGatewayUsagePlanKey](#ApiGatewayUsagePlanKey)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayUsagePlanKeySpec](#ApiGatewayUsagePlanKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
