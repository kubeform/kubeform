## EfsFileSystem
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EfsFileSystem` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EfsFileSystemSpec](#EfsFileSystemSpec)***||
| `status` | ***[EfsFileSystemStatus](#EfsFileSystemStatus)***||
## EfsFileSystemSpec
##### (Appears on:[EfsFileSystem](#EfsFileSystem), [EfsFileSystemStatus](#EfsFileSystemStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `creationToken` | ***string***| ***(Optional)*** |
| `dnsName` | ***string***| ***(Optional)*** |
| `encrypted` | ***bool***| ***(Optional)*** |
| `kmsKeyID` | ***string***| ***(Optional)*** |
| `performanceMode` | ***string***| ***(Optional)*** |
| `provisionedThroughputInMibps` | ***encoding/json.Number***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `throughputMode` | ***string***| ***(Optional)*** |
## EfsFileSystemStatus
##### (Appears on:[EfsFileSystem](#EfsFileSystem))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EfsFileSystemSpec](#EfsFileSystemSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
