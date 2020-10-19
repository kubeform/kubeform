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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EcsTaskDefinitionSpec](#ecstaskdefinitionspec)***||
| `status` | ***[EcsTaskDefinitionStatus](#ecstaskdefinitionstatus)***||
## EcsTaskDefinitionSpec

Appears on:[EcsTaskDefinition](#ecstaskdefinition), [EcsTaskDefinitionStatus](#ecstaskdefinitionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
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
| `placementConstraints` | ***[[]EcsTaskDefinitionSpecPlacementConstraints](#ecstaskdefinitionspecplacementconstraints)***| ***(Optional)*** |
| `proxyConfiguration` | ***[[]EcsTaskDefinitionSpecProxyConfiguration](#ecstaskdefinitionspecproxyconfiguration)***| ***(Optional)*** |
| `requiresCompatibilities` | ***[]string***| ***(Optional)*** |
| `revision` | ***int64***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `taskRoleArn` | ***string***| ***(Optional)*** |
| `volume` | ***[[]EcsTaskDefinitionSpecVolume](#ecstaskdefinitionspecvolume)***| ***(Optional)*** |
## EcsTaskDefinitionSpecPlacementConstraints

Appears on:[EcsTaskDefinitionSpec](#ecstaskdefinitionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `expression` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## EcsTaskDefinitionSpecProxyConfiguration

Appears on:[EcsTaskDefinitionSpec](#ecstaskdefinitionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `containerName` | ***string***||
| `properties` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## EcsTaskDefinitionSpecVolume

Appears on:[EcsTaskDefinitionSpec](#ecstaskdefinitionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dockerVolumeConfiguration` | ***[[]EcsTaskDefinitionSpecVolumeDockerVolumeConfiguration](#ecstaskdefinitionspecvolumedockervolumeconfiguration)***| ***(Optional)*** |
| `hostPath` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## EcsTaskDefinitionSpecVolumeDockerVolumeConfiguration

Appears on:[EcsTaskDefinitionSpecVolume](#ecstaskdefinitionspecvolume)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `autoprovision` | ***bool***| ***(Optional)*** |
| `driver` | ***string***| ***(Optional)*** |
| `driverOpts` | ***map[string]string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `scope` | ***string***| ***(Optional)*** |
## EcsTaskDefinitionStatus

Appears on:[EcsTaskDefinition](#ecstaskdefinition)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EcsTaskDefinitionSpec](#ecstaskdefinitionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[EcsTaskDefinitionStatus](#ecstaskdefinitionstatus)

---
