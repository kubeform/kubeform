---
title: SagemakerModel
menu:
  docs_v0.0.1:
    identifier: sagemakermodel-aws.kubeform.com
    name: SagemakerModel
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SagemakerModel
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SagemakerModel` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SagemakerModelSpec](#SagemakerModelSpec)***||
| `status` | ***[SagemakerModelStatus](#SagemakerModelStatus)***||
## SagemakerModelSpec
##### (Appears on:[SagemakerModel](#SagemakerModel), [SagemakerModelStatus](#SagemakerModelStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `container` | ***[[]SagemakerModelSpecContainer](#SagemakerModelSpecContainer)***| ***(Optional)*** |
| `enableNetworkIsolation` | ***bool***| ***(Optional)*** |
| `executionRoleArn` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `primaryContainer` | ***[[]SagemakerModelSpecPrimaryContainer](#SagemakerModelSpecPrimaryContainer)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcConfig` | ***[[]SagemakerModelSpecVpcConfig](#SagemakerModelSpecVpcConfig)***| ***(Optional)*** |
## SagemakerModelSpecContainer
##### (Appears on:[SagemakerModelSpec](#SagemakerModelSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `containerHostname` | ***string***| ***(Optional)*** |
| `environment` | ***map[string]string***| ***(Optional)*** |
| `image` | ***string***||
| `modelDataURL` | ***string***| ***(Optional)*** |
## SagemakerModelSpecPrimaryContainer
##### (Appears on:[SagemakerModelSpec](#SagemakerModelSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `containerHostname` | ***string***| ***(Optional)*** |
| `environment` | ***map[string]string***| ***(Optional)*** |
| `image` | ***string***||
| `modelDataURL` | ***string***| ***(Optional)*** |
## SagemakerModelSpecVpcConfig
##### (Appears on:[SagemakerModelSpec](#SagemakerModelSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `securityGroupIDS` | ***[]string***||
| `subnets` | ***[]string***||
## SagemakerModelStatus
##### (Appears on:[SagemakerModel](#SagemakerModel))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SagemakerModelSpec](#SagemakerModelSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
