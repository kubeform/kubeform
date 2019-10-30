---
title: DataFactoryPipeline
menu:
  docs_{{ .version }}:
    identifier: datafactorypipeline-azurerm.kubeform.com
    name: DataFactoryPipeline
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataFactoryPipeline
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryPipeline` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryPipelineSpec](#datafactorypipelinespec)***||
| `status` | ***[DataFactoryPipelineStatus](#datafactorypipelinestatus)***||
## DataFactoryPipelineSpec

Appears on:[DataFactoryPipeline](#datafactorypipeline), [DataFactoryPipelineStatus](#datafactorypipelinestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `annotations` | ***[]string***| ***(Optional)*** |
| `dataFactoryName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `variables` | ***map[string]string***| ***(Optional)*** |
## DataFactoryPipelineStatus

Appears on:[DataFactoryPipeline](#datafactorypipeline)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryPipelineSpec](#datafactorypipelinespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DataFactoryPipelineStatus](#datafactorypipelinestatus)

---
