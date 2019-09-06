## DataFactoryPipeline
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryPipeline` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryPipelineSpec](#DataFactoryPipelineSpec)***||
| `status` | ***[DataFactoryPipelineStatus](#DataFactoryPipelineStatus)***||
## DataFactoryPipelineSpec
##### (Appears on:[DataFactoryPipeline](#DataFactoryPipeline), [DataFactoryPipelineStatus](#DataFactoryPipelineStatus))
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
##### (Appears on:[DataFactoryPipeline](#DataFactoryPipeline))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryPipelineSpec](#DataFactoryPipelineSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
