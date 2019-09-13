---
title: ApiGatewayStage
menu:
  docs_{{ .version }}:
    identifier: apigatewaystage-aws.kubeform.com
    name: ApiGatewayStage
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayStage
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayStage` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayStageSpec](#ApiGatewayStageSpec)***||
| `status` | ***[ApiGatewayStageStatus](#ApiGatewayStageStatus)***||
## ApiGatewayStageSpec
##### (Appears on:[ApiGatewayStage](#ApiGatewayStage), [ApiGatewayStageStatus](#ApiGatewayStageStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessLogSettings` | ***[[]ApiGatewayStageSpecAccessLogSettings](#ApiGatewayStageSpecAccessLogSettings)***| ***(Optional)*** |
| `cacheClusterEnabled` | ***bool***| ***(Optional)*** |
| `cacheClusterSize` | ***string***| ***(Optional)*** |
| `clientCertificateID` | ***string***| ***(Optional)*** |
| `deploymentID` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `documentationVersion` | ***string***| ***(Optional)*** |
| `executionArn` | ***string***| ***(Optional)*** |
| `invokeURL` | ***string***| ***(Optional)*** |
| `restAPIID` | ***string***||
| `stageName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `variables` | ***map[string]string***| ***(Optional)*** |
| `xrayTracingEnabled` | ***bool***| ***(Optional)*** |
## ApiGatewayStageSpecAccessLogSettings
##### (Appears on:[ApiGatewayStageSpec](#ApiGatewayStageSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `destinationArn` | ***string***||
| `format` | ***string***||
## ApiGatewayStageStatus
##### (Appears on:[ApiGatewayStage](#ApiGatewayStage))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayStageSpec](#ApiGatewayStageSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
