## S3BucketPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `S3BucketPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[S3BucketPolicySpec](#S3BucketPolicySpec)***||
| `status` | ***[S3BucketPolicyStatus](#S3BucketPolicyStatus)***||
## S3BucketPolicySpec
##### (Appears on:[S3BucketPolicy](#S3BucketPolicy), [S3BucketPolicyStatus](#S3BucketPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `policy` | ***string***||
## S3BucketPolicyStatus
##### (Appears on:[S3BucketPolicy](#S3BucketPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[S3BucketPolicySpec](#S3BucketPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
