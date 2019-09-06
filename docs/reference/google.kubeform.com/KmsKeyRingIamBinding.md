## KmsKeyRingIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsKeyRingIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsKeyRingIamBindingSpec](#KmsKeyRingIamBindingSpec)***||
| `status` | ***[KmsKeyRingIamBindingStatus](#KmsKeyRingIamBindingStatus)***||
## KmsKeyRingIamBindingSpec
##### (Appears on:[KmsKeyRingIamBinding](#KmsKeyRingIamBinding), [KmsKeyRingIamBindingStatus](#KmsKeyRingIamBindingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `keyRingID` | ***string***||
| `members` | ***[]string***||
| `role` | ***string***||
## KmsKeyRingIamBindingStatus
##### (Appears on:[KmsKeyRingIamBinding](#KmsKeyRingIamBinding))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsKeyRingIamBindingSpec](#KmsKeyRingIamBindingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
