## Folder
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `Folder` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FolderSpec](#FolderSpec)***||
| `status` | ***[FolderStatus](#FolderStatus)***||
## FolderSpec
##### (Appears on:[Folder](#Folder), [FolderStatus](#FolderStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createTime` | ***string***| ***(Optional)*** |
| `displayName` | ***string***||
| `lifecycleState` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `parent` | ***string***||
## FolderStatus
##### (Appears on:[Folder](#Folder))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FolderSpec](#FolderSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
