## SpacesBucket
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `SpacesBucket` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpacesBucketSpec](#SpacesBucketSpec)***||
| `status` | ***[SpacesBucketStatus](#SpacesBucketStatus)***||
## SpacesBucketSpec
##### (Appears on:[SpacesBucket](#SpacesBucket), [SpacesBucketStatus](#SpacesBucketStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `acl` | ***string***| ***(Optional)*** Canned ACL applied on bucket creation|
| `bucketDomainName` | ***string***| ***(Optional)*** The FQDN of the bucket|
| `forceDestroy` | ***bool***| ***(Optional)*** Unless true, the bucket will only be destroyed if empty|
| `name` | ***string***|Bucket name|
| `region` | ***string***| ***(Optional)*** Bucket region|
| `urn` | ***string***| ***(Optional)*** the uniform resource name for the bucket|
## SpacesBucketStatus
##### (Appears on:[SpacesBucket](#SpacesBucket))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpacesBucketSpec](#SpacesBucketSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
