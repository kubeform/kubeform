## IotPolicyAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IotPolicyAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IotPolicyAttachmentSpec](#IotPolicyAttachmentSpec)***||
| `status` | ***[IotPolicyAttachmentStatus](#IotPolicyAttachmentStatus)***||
## IotPolicyAttachmentSpec
##### (Appears on:[IotPolicyAttachment](#IotPolicyAttachment), [IotPolicyAttachmentStatus](#IotPolicyAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `policy` | ***string***||
| `target` | ***string***||
## IotPolicyAttachmentStatus
##### (Appears on:[IotPolicyAttachment](#IotPolicyAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IotPolicyAttachmentSpec](#IotPolicyAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
