---
title: ComposerEnvironment
menu:
  docs_{{ .version }}:
    identifier: composerenvironment-google.kubeform.com
    name: ComposerEnvironment
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComposerEnvironment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComposerEnvironment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComposerEnvironmentSpec](#ComposerEnvironmentSpec)***||
| `status` | ***[ComposerEnvironmentStatus](#ComposerEnvironmentStatus)***||
## ComposerEnvironmentSpec

Appears on:[ComposerEnvironment](#ComposerEnvironment), [ComposerEnvironmentStatus](#ComposerEnvironmentStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `config` | ***[[]ComposerEnvironmentSpecConfig](#ComposerEnvironmentSpecConfig)***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
## ComposerEnvironmentSpecConfig

Appears on:[ComposerEnvironmentSpec](#ComposerEnvironmentSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `airflowURI` | ***string***| ***(Optional)*** |
| `dagGcsPrefix` | ***string***| ***(Optional)*** |
| `gkeCluster` | ***string***| ***(Optional)*** |
| `nodeConfig` | ***[[]ComposerEnvironmentSpecConfigNodeConfig](#ComposerEnvironmentSpecConfigNodeConfig)***| ***(Optional)*** |
| `nodeCount` | ***int***| ***(Optional)*** |
| `softwareConfig` | ***[[]ComposerEnvironmentSpecConfigSoftwareConfig](#ComposerEnvironmentSpecConfigSoftwareConfig)***| ***(Optional)*** |
## ComposerEnvironmentSpecConfigNodeConfig

Appears on:[ComposerEnvironmentSpecConfig](#ComposerEnvironmentSpecConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `diskSizeGb` | ***int***| ***(Optional)*** |
| `machineType` | ***string***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `oauthScopes` | ***[]string***| ***(Optional)*** |
| `serviceAccount` | ***string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## ComposerEnvironmentSpecConfigSoftwareConfig

Appears on:[ComposerEnvironmentSpecConfig](#ComposerEnvironmentSpecConfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `airflowConfigOverrides` | ***map[string]string***| ***(Optional)*** |
| `envVariables` | ***map[string]string***| ***(Optional)*** |
| `imageVersion` | ***string***| ***(Optional)*** |
| `pypiPackages` | ***map[string]string***| ***(Optional)*** |
## ComposerEnvironmentStatus

Appears on:[ComposerEnvironment](#ComposerEnvironment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComposerEnvironmentSpec](#ComposerEnvironmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
