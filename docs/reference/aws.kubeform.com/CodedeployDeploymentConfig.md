---
title: CodedeployDeploymentConfig
menu:
  docs_{{ .version }}:
    identifier: codedeploydeploymentconfig-aws.kubeform.com
    name: CodedeployDeploymentConfig
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CodedeployDeploymentConfig
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CodedeployDeploymentConfig` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodedeployDeploymentConfigSpec](#CodedeployDeploymentConfigSpec)***||
| `status` | ***[CodedeployDeploymentConfigStatus](#CodedeployDeploymentConfigStatus)***||
## CodedeployDeploymentConfigSpec

Appears on:[CodedeployDeploymentConfig](#CodedeployDeploymentConfig), [CodedeployDeploymentConfigStatus](#CodedeployDeploymentConfigStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `computePlatform` | ***string***| ***(Optional)*** |
| `deploymentConfigID` | ***string***| ***(Optional)*** |
| `deploymentConfigName` | ***string***||
| `minimumHealthyHosts` | ***[[]CodedeployDeploymentConfigSpecMinimumHealthyHosts](#CodedeployDeploymentConfigSpecMinimumHealthyHosts)***| ***(Optional)*** |
| `trafficRoutingConfig` | ***[[]CodedeployDeploymentConfigSpecTrafficRoutingConfig](#CodedeployDeploymentConfigSpecTrafficRoutingConfig)***| ***(Optional)*** |
## CodedeployDeploymentConfigSpecMinimumHealthyHosts

Appears on:[CodedeployDeploymentConfigSpec](#CodedeployDeploymentConfigSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***| ***(Optional)*** |
| `value` | ***int***| ***(Optional)*** |
## CodedeployDeploymentConfigSpecTrafficRoutingConfig

Appears on:[CodedeployDeploymentConfigSpec](#CodedeployDeploymentConfigSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `timeBasedCanary` | ***[[]CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary](#CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary)***| ***(Optional)*** |
| `timeBasedLinear` | ***[[]CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear](#CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear)***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary

Appears on:[CodedeployDeploymentConfigSpecTrafficRoutingConfig](#CodedeployDeploymentConfigSpecTrafficRoutingConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `interval` | ***int***| ***(Optional)*** |
| `percentage` | ***int***| ***(Optional)*** |
## CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear

Appears on:[CodedeployDeploymentConfigSpecTrafficRoutingConfig](#CodedeployDeploymentConfigSpecTrafficRoutingConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `interval` | ***int***| ***(Optional)*** |
| `percentage` | ***int***| ***(Optional)*** |
## CodedeployDeploymentConfigStatus

Appears on:[CodedeployDeploymentConfig](#CodedeployDeploymentConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodedeployDeploymentConfigSpec](#CodedeployDeploymentConfigSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
