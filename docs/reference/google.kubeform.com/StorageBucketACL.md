## StorageBucketACL
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageBucketACL` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageBucketACLSpec](#StorageBucketACLSpec)***||
| `status` | ***[StorageBucketACLStatus](#StorageBucketACLStatus)***||
## StorageBucketACLSpec
##### (Appears on:[StorageBucketACL](#StorageBucketACL), [StorageBucketACLStatus](#StorageBucketACLStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `defaultACL` | ***string***| ***(Optional)*** |
| `predefinedACL` | ***string***| ***(Optional)*** |
| `roleEntity` | ***[]string***| ***(Optional)*** |
## StorageBucketACLStatus
##### (Appears on:[StorageBucketACL](#StorageBucketACL))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageBucketACLSpec](#StorageBucketACLSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
