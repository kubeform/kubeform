## BatchApplication
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BatchApplication` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BatchApplicationSpec](#BatchApplicationSpec)***||
| `status` | ***[BatchApplicationStatus](#BatchApplicationStatus)***||
## BatchApplicationSpec
##### (Appears on:[BatchApplication](#BatchApplication), [BatchApplicationStatus](#BatchApplicationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `allowUpdates` | ***bool***| ***(Optional)*** |
| `defaultVersion` | ***string***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## BatchApplicationStatus
##### (Appears on:[BatchApplication](#BatchApplication))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BatchApplicationSpec](#BatchApplicationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
