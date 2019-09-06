## SagemakerNotebookInstanceLifecycleConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SagemakerNotebookInstanceLifecycleConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SagemakerNotebookInstanceLifecycleConfigurationSpec](#SagemakerNotebookInstanceLifecycleConfigurationSpec)***||
| `status` | ***[SagemakerNotebookInstanceLifecycleConfigurationStatus](#SagemakerNotebookInstanceLifecycleConfigurationStatus)***||
## SagemakerNotebookInstanceLifecycleConfigurationSpec
##### (Appears on:[SagemakerNotebookInstanceLifecycleConfiguration](#SagemakerNotebookInstanceLifecycleConfiguration), [SagemakerNotebookInstanceLifecycleConfigurationStatus](#SagemakerNotebookInstanceLifecycleConfigurationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `onCreate` | ***string***| ***(Optional)*** |
| `onStart` | ***string***| ***(Optional)*** |
## SagemakerNotebookInstanceLifecycleConfigurationStatus
##### (Appears on:[SagemakerNotebookInstanceLifecycleConfiguration](#SagemakerNotebookInstanceLifecycleConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SagemakerNotebookInstanceLifecycleConfigurationSpec](#SagemakerNotebookInstanceLifecycleConfigurationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
