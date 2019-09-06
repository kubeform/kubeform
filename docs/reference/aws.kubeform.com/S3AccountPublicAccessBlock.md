## S3AccountPublicAccessBlock
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3AccountPublicAccessBlock` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3AccountPublicAccessBlockSpec](#S3AccountPublicAccessBlockSpec)***||
| `status` | ***[S3AccountPublicAccessBlockStatus](#S3AccountPublicAccessBlockStatus)***||
## S3AccountPublicAccessBlockSpec
##### (Appears on:[S3AccountPublicAccessBlock](#S3AccountPublicAccessBlock), [S3AccountPublicAccessBlockStatus](#S3AccountPublicAccessBlockStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***| ***(Optional)*** |
| `blockPublicAcls` | ***bool***| ***(Optional)*** |
| `blockPublicPolicy` | ***bool***| ***(Optional)*** |
| `ignorePublicAcls` | ***bool***| ***(Optional)*** |
| `restrictPublicBuckets` | ***bool***| ***(Optional)*** |
## S3AccountPublicAccessBlockStatus
##### (Appears on:[S3AccountPublicAccessBlock](#S3AccountPublicAccessBlock))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3AccountPublicAccessBlockSpec](#S3AccountPublicAccessBlockSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
