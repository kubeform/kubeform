## NetworkInterfaceAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkInterfaceAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkInterfaceAttachmentSpec](#NetworkInterfaceAttachmentSpec)***||
| `status` | ***[NetworkInterfaceAttachmentStatus](#NetworkInterfaceAttachmentStatus)***||
## NetworkInterfaceAttachmentSpec
##### (Appears on:[NetworkInterfaceAttachment](#NetworkInterfaceAttachment), [NetworkInterfaceAttachmentStatus](#NetworkInterfaceAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attachmentID` | ***string***| ***(Optional)*** |
| `deviceIndex` | ***int***||
| `instanceID` | ***string***||
| `networkInterfaceID` | ***string***||
| `status` | ***string***| ***(Optional)*** |
## NetworkInterfaceAttachmentStatus
##### (Appears on:[NetworkInterfaceAttachment](#NetworkInterfaceAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkInterfaceAttachmentSpec](#NetworkInterfaceAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
