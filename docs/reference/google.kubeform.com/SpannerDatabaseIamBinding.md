## SpannerDatabaseIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SpannerDatabaseIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpannerDatabaseIamBindingSpec](#SpannerDatabaseIamBindingSpec)***||
| `status` | ***[SpannerDatabaseIamBindingStatus](#SpannerDatabaseIamBindingStatus)***||
## SpannerDatabaseIamBindingSpec
##### (Appears on:[SpannerDatabaseIamBinding](#SpannerDatabaseIamBinding), [SpannerDatabaseIamBindingStatus](#SpannerDatabaseIamBindingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `database` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## SpannerDatabaseIamBindingStatus
##### (Appears on:[SpannerDatabaseIamBinding](#SpannerDatabaseIamBinding))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpannerDatabaseIamBindingSpec](#SpannerDatabaseIamBindingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
