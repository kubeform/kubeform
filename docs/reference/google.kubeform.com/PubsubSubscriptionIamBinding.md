## PubsubSubscriptionIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `PubsubSubscriptionIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PubsubSubscriptionIamBindingSpec](#PubsubSubscriptionIamBindingSpec)***||
| `status` | ***[PubsubSubscriptionIamBindingStatus](#PubsubSubscriptionIamBindingStatus)***||
## PubsubSubscriptionIamBindingSpec
##### (Appears on:[PubsubSubscriptionIamBinding](#PubsubSubscriptionIamBinding), [PubsubSubscriptionIamBindingStatus](#PubsubSubscriptionIamBindingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `subscription` | ***string***||
## PubsubSubscriptionIamBindingStatus
##### (Appears on:[PubsubSubscriptionIamBinding](#PubsubSubscriptionIamBinding))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PubsubSubscriptionIamBindingSpec](#PubsubSubscriptionIamBindingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
