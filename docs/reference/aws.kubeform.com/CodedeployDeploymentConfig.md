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
| `spec` | ***[CodedeployDeploymentConfigSpec](#codedeploydeploymentconfigspec)***||
| `status` | ***[CodedeployDeploymentConfigStatus](#codedeploydeploymentconfigstatus)***||
## CodedeployDeploymentConfigSpec

Appears on:[CodedeployDeploymentConfig](#codedeploydeploymentconfig), [CodedeployDeploymentConfigStatus](#codedeploydeploymentconfigstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `computePlatform` | ***string***| ***(Optional)*** |
| `deploymentConfigID` | ***string***| ***(Optional)*** |
| `deploymentConfigName` | ***string***||
| `minimumHealthyHosts` | ***[[]CodedeployDeploymentConfigSpecMinimumHealthyHosts](#codedeploydeploymentconfigspecminimumhealthyhosts)***| ***(Optional)*** |
| `trafficRoutingConfig` | ***[[]CodedeployDeploymentConfigSpecTrafficRoutingConfig](#codedeploydeploymentconfigspectrafficroutingconfig)***| ***(Optional)*** |
## CodedeployDeploymentConfigSpecMinimumHealthyHosts

Appears on:[CodedeployDeploymentConfigSpec](#codedeploydeploymentconfigspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***| ***(Optional)*** |
| `value` | ***int***| ***(Optional)*** |
## CodedeployDeploymentConfigSpecTrafficRoutingConfig

Appears on:[CodedeployDeploymentConfigSpec](#codedeploydeploymentconfigspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `timeBasedCanary` | ***[[]CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary](#codedeploydeploymentconfigspectrafficroutingconfigtimebasedcanary)***| ***(Optional)*** |
| `timeBasedLinear` | ***[[]CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear](#codedeploydeploymentconfigspectrafficroutingconfigtimebasedlinear)***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedCanary

Appears on:[CodedeployDeploymentConfigSpecTrafficRoutingConfig](#codedeploydeploymentconfigspectrafficroutingconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `interval` | ***int***| ***(Optional)*** |
| `percentage` | ***int***| ***(Optional)*** |
## CodedeployDeploymentConfigSpecTrafficRoutingConfigTimeBasedLinear

Appears on:[CodedeployDeploymentConfigSpecTrafficRoutingConfig](#codedeploydeploymentconfigspectrafficroutingconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `interval` | ***int***| ***(Optional)*** |
| `percentage` | ***int***| ***(Optional)*** |
## CodedeployDeploymentConfigStatus

Appears on:[CodedeployDeploymentConfig](#codedeploydeploymentconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodedeployDeploymentConfigSpec](#codedeploydeploymentconfigspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CodedeployDeploymentConfigStatus](#codedeploydeploymentconfigstatus)

---
