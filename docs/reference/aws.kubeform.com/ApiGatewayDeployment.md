---
title: ApiGatewayDeployment
menu:
  docs_{{ .version }}:
    identifier: apigatewaydeployment-aws.kubeform.com
    name: ApiGatewayDeployment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiGatewayDeployment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiGatewayDeployment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiGatewayDeploymentSpec](#ApiGatewayDeploymentSpec)***||
| `status` | ***[ApiGatewayDeploymentStatus](#ApiGatewayDeploymentStatus)***||
## ApiGatewayDeploymentSpec
##### (Appears on:[ApiGatewayDeployment](#ApiGatewayDeployment), [ApiGatewayDeploymentStatus](#ApiGatewayDeploymentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createdDate` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `executionArn` | ***string***| ***(Optional)*** |
| `invokeURL` | ***string***| ***(Optional)*** |
| `restAPIID` | ***string***||
| `stageDescription` | ***string***| ***(Optional)*** |
| `stageName` | ***string***| ***(Optional)*** |
| `variables` | ***map[string]string***| ***(Optional)*** |
## ApiGatewayDeploymentStatus
##### (Appears on:[ApiGatewayDeployment](#ApiGatewayDeployment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiGatewayDeploymentSpec](#ApiGatewayDeploymentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
