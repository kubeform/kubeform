---
title: ApiGatewayUsagePlan
menu:
  docs_{{ .version }}:
    identifier: apigatewayusageplan-aws.kubeform.com
    name: ApiGatewayUsagePlan
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayUsagePlan
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayUsagePlan` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayUsagePlanSpec](#ApiGatewayUsagePlanSpec)***||
| `status` | ***[ApiGatewayUsagePlanStatus](#ApiGatewayUsagePlanStatus)***||
## ApiGatewayUsagePlanSpec
##### (Appears on:[ApiGatewayUsagePlan](#ApiGatewayUsagePlan), [ApiGatewayUsagePlanStatus](#ApiGatewayUsagePlanStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apiStages` | ***[[]ApiGatewayUsagePlanSpecApiStages](#ApiGatewayUsagePlanSpecApiStages)***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `productCode` | ***string***| ***(Optional)*** |
| `quotaSettings` | ***[[]ApiGatewayUsagePlanSpecQuotaSettings](#ApiGatewayUsagePlanSpecQuotaSettings)***| ***(Optional)*** |
| `throttleSettings` | ***[[]ApiGatewayUsagePlanSpecThrottleSettings](#ApiGatewayUsagePlanSpecThrottleSettings)***| ***(Optional)*** |
## ApiGatewayUsagePlanSpecApiStages
##### (Appears on:[ApiGatewayUsagePlanSpec](#ApiGatewayUsagePlanSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiID` | ***string***||
| `stage` | ***string***||
## ApiGatewayUsagePlanSpecQuotaSettings
##### (Appears on:[ApiGatewayUsagePlanSpec](#ApiGatewayUsagePlanSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `limit` | ***int***||
| `offset` | ***int***| ***(Optional)*** |
| `period` | ***string***||
## ApiGatewayUsagePlanSpecThrottleSettings
##### (Appears on:[ApiGatewayUsagePlanSpec](#ApiGatewayUsagePlanSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `burstLimit` | ***int***| ***(Optional)*** |
| `rateLimit` | ***encoding/json.Number***| ***(Optional)*** |
## ApiGatewayUsagePlanStatus
##### (Appears on:[ApiGatewayUsagePlan](#ApiGatewayUsagePlan))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayUsagePlanSpec](#ApiGatewayUsagePlanSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
