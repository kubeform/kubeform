## LoggingProjectSink
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `LoggingProjectSink` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoggingProjectSinkSpec](#LoggingProjectSinkSpec)***||
| `status` | ***[LoggingProjectSinkStatus](#LoggingProjectSinkStatus)***||
## LoggingProjectSinkSpec
##### (Appears on:[LoggingProjectSink](#LoggingProjectSink), [LoggingProjectSinkStatus](#LoggingProjectSinkStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `destination` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `uniqueWriterIdentity` | ***bool***| ***(Optional)*** |
| `writerIdentity` | ***string***| ***(Optional)*** |
## LoggingProjectSinkStatus
##### (Appears on:[LoggingProjectSink](#LoggingProjectSink))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoggingProjectSinkSpec](#LoggingProjectSinkSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
