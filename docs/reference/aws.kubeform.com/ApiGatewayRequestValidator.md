---
title: ApiGatewayRequestValidator
menu:
  docs_v0.0.1:
    identifier: apigatewayrequestvalidator-aws.kubeform.com
    name: ApiGatewayRequestValidator
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApiGatewayRequestValidator
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayRequestValidator` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayRequestValidatorSpec](#ApiGatewayRequestValidatorSpec)***||
| `status` | ***[ApiGatewayRequestValidatorStatus](#ApiGatewayRequestValidatorStatus)***||
## ApiGatewayRequestValidatorSpec
##### (Appears on:[ApiGatewayRequestValidator](#ApiGatewayRequestValidator), [ApiGatewayRequestValidatorStatus](#ApiGatewayRequestValidatorStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `restAPIID` | ***string***||
| `validateRequestBody` | ***bool***| ***(Optional)*** |
| `validateRequestParameters` | ***bool***| ***(Optional)*** |
## ApiGatewayRequestValidatorStatus
##### (Appears on:[ApiGatewayRequestValidator](#ApiGatewayRequestValidator))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayRequestValidatorSpec](#ApiGatewayRequestValidatorSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
