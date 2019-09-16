---
title: Codepipeline
menu:
  docs_{{ .version }}:
    identifier: codepipeline-aws.kubeform.com
    name: Codepipeline
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Codepipeline
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Codepipeline` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CodepipelineSpec](#CodepipelineSpec)***||
| `status` | ***[CodepipelineStatus](#CodepipelineStatus)***||
## CodepipelineSpec

Appears on:[Codepipeline](#Codepipeline), [CodepipelineStatus](#CodepipelineStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `artifactStore` | ***[[]CodepipelineSpecArtifactStore](#CodepipelineSpecArtifactStore)***||
| `name` | ***string***||
| `roleArn` | ***string***||
| `stage` | ***[[]CodepipelineSpecStage](#CodepipelineSpecStage)***||
## CodepipelineSpecArtifactStore

Appears on:[CodepipelineSpec](#CodepipelineSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `encryptionKey` | ***[[]CodepipelineSpecArtifactStoreEncryptionKey](#CodepipelineSpecArtifactStoreEncryptionKey)***| ***(Optional)*** |
| `location` | ***string***||
| `type` | ***string***||
## CodepipelineSpecArtifactStoreEncryptionKey

Appears on:[CodepipelineSpecArtifactStore](#CodepipelineSpecArtifactStore)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***||
| `type` | ***string***||
## CodepipelineSpecStage

Appears on:[CodepipelineSpec](#CodepipelineSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]CodepipelineSpecStageAction](#CodepipelineSpecStageAction)***||
| `name` | ***string***||
## CodepipelineSpecStageAction

Appears on:[CodepipelineSpecStage](#CodepipelineSpecStage)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `category` | ***string***||
| `configuration` | ***map[string]string***| ***(Optional)*** |
| `inputArtifacts` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `outputArtifacts` | ***[]string***| ***(Optional)*** |
| `owner` | ***string***||
| `provider` | ***string***||
| `roleArn` | ***string***| ***(Optional)*** |
| `runOrder` | ***int***| ***(Optional)*** |
| `version` | ***string***||
## CodepipelineStatus

Appears on:[Codepipeline](#Codepipeline)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodepipelineSpec](#CodepipelineSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
