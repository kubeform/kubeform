## SpannerInstanceIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SpannerInstanceIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpannerInstanceIamBindingSpec](#SpannerInstanceIamBindingSpec)***||
| `status` | ***[SpannerInstanceIamBindingStatus](#SpannerInstanceIamBindingStatus)***||
## SpannerInstanceIamBindingSpec
##### (Appears on:[SpannerInstanceIamBinding](#SpannerInstanceIamBinding), [SpannerInstanceIamBindingStatus](#SpannerInstanceIamBindingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## SpannerInstanceIamBindingStatus
##### (Appears on:[SpannerInstanceIamBinding](#SpannerInstanceIamBinding))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpannerInstanceIamBindingSpec](#SpannerInstanceIamBindingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
