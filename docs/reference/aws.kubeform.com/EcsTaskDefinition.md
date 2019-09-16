---
title: EcsTaskDefinition
menu:
  docs_{{ .version }}:
    identifier: ecstaskdefinition-aws.kubeform.com
    name: EcsTaskDefinition
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EcsTaskDefinition
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EcsTaskDefinition` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EcsTaskDefinitionSpec](#EcsTaskDefinitionSpec)***||
| `status` | ***[EcsTaskDefinitionStatus](#EcsTaskDefinitionStatus)***||
## EcsTaskDefinitionSpec

Appears on:[EcsTaskDefinition](#EcsTaskDefinition), [EcsTaskDefinitionStatus](#EcsTaskDefinitionStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `containerDefinitions` | ***string***||
| `cpu` | ***string***| ***(Optional)*** |
| `executionRoleArn` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `ipcMode` | ***string***| ***(Optional)*** |
| `memory` | ***string***| ***(Optional)*** |
| `networkMode` | ***string***| ***(Optional)*** |
| `pidMode` | ***string***| ***(Optional)*** |
| `placementConstraints` | ***[[]EcsTaskDefinitionSpecPlacementConstraints](#EcsTaskDefinitionSpecPlacementConstraints)***| ***(Optional)*** |
| `requiresCompatibilities` | ***[]string***| ***(Optional)*** |
| `revision` | ***int***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `taskRoleArn` | ***string***| ***(Optional)*** |
| `volume` | ***[[]EcsTaskDefinitionSpecVolume](#EcsTaskDefinitionSpecVolume)***| ***(Optional)*** |
## EcsTaskDefinitionSpecPlacementConstraints

Appears on:[EcsTaskDefinitionSpec](#EcsTaskDefinitionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `expression` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## EcsTaskDefinitionSpecVolume

Appears on:[EcsTaskDefinitionSpec](#EcsTaskDefinitionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dockerVolumeConfiguration` | ***[[]EcsTaskDefinitionSpecVolumeDockerVolumeConfiguration](#EcsTaskDefinitionSpecVolumeDockerVolumeConfiguration)***| ***(Optional)*** |
| `hostPath` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## EcsTaskDefinitionSpecVolumeDockerVolumeConfiguration

Appears on:[EcsTaskDefinitionSpecVolume](#EcsTaskDefinitionSpecVolume)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoprovision` | ***bool***| ***(Optional)*** |
| `driver` | ***string***| ***(Optional)*** |
| `driverOpts` | ***map[string]string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `scope` | ***string***| ***(Optional)*** |
## EcsTaskDefinitionStatus

Appears on:[EcsTaskDefinition](#EcsTaskDefinition)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EcsTaskDefinitionSpec](#EcsTaskDefinitionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
