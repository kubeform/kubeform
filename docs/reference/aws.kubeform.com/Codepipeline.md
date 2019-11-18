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
| `spec` | ***[CodepipelineSpec](#codepipelinespec)***||
| `status` | ***[CodepipelineStatus](#codepipelinestatus)***||
## CodepipelineSpec

Appears on:[Codepipeline](#codepipeline), [CodepipelineStatus](#codepipelinestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `artifactStore` | ***[[]CodepipelineSpecArtifactStore](#codepipelinespecartifactstore)***||
| `name` | ***string***||
| `roleArn` | ***string***||
| `stage` | ***[[]CodepipelineSpecStage](#codepipelinespecstage)***||
## CodepipelineSpecArtifactStore

Appears on:[CodepipelineSpec](#codepipelinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `encryptionKey` | ***[[]CodepipelineSpecArtifactStoreEncryptionKey](#codepipelinespecartifactstoreencryptionkey)***| ***(Optional)*** |
| `location` | ***string***||
| `type` | ***string***||
## CodepipelineSpecArtifactStoreEncryptionKey

Appears on:[CodepipelineSpecArtifactStore](#codepipelinespecartifactstore)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***||
| `type` | ***string***||
## CodepipelineSpecStage

Appears on:[CodepipelineSpec](#codepipelinespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***[[]CodepipelineSpecStageAction](#codepipelinespecstageaction)***||
| `name` | ***string***||
## CodepipelineSpecStageAction

Appears on:[CodepipelineSpecStage](#codepipelinespecstage)

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
| `runOrder` | ***int64***| ***(Optional)*** |
| `version` | ***string***||
## CodepipelineStatus

Appears on:[Codepipeline](#codepipeline)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CodepipelineSpec](#codepipelinespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CodepipelineStatus](#codepipelinestatus)

---
