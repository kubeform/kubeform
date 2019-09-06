## AlbTargetGroupAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AlbTargetGroupAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AlbTargetGroupAttachmentSpec](#AlbTargetGroupAttachmentSpec)***||
| `status` | ***[AlbTargetGroupAttachmentStatus](#AlbTargetGroupAttachmentStatus)***||
## AlbTargetGroupAttachmentSpec
##### (Appears on:[AlbTargetGroupAttachment](#AlbTargetGroupAttachment), [AlbTargetGroupAttachmentStatus](#AlbTargetGroupAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availabilityZone` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `targetGroupArn` | ***string***||
| `targetID` | ***string***||
## AlbTargetGroupAttachmentStatus
##### (Appears on:[AlbTargetGroupAttachment](#AlbTargetGroupAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AlbTargetGroupAttachmentSpec](#AlbTargetGroupAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
