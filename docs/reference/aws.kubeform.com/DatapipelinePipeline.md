---
title: DatapipelinePipeline
menu:
  docs_{{ .version }}:
    identifier: datapipelinepipeline-aws.kubeform.com
    name: DatapipelinePipeline
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DatapipelinePipeline
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DatapipelinePipeline` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatapipelinePipelineSpec](#datapipelinepipelinespec)***||
| `status` | ***[DatapipelinePipelineStatus](#datapipelinepipelinestatus)***||
## DatapipelinePipelineSpec

Appears on:[DatapipelinePipeline](#datapipelinepipeline), [DatapipelinePipelineStatus](#datapipelinepipelinestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## DatapipelinePipelineStatus

Appears on:[DatapipelinePipeline](#datapipelinepipeline)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatapipelinePipelineSpec](#datapipelinepipelinespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DatapipelinePipelineStatus](#datapipelinepipelinestatus)

---
