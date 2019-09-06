## StoragegatewayCache
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `StoragegatewayCache` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StoragegatewayCacheSpec](#StoragegatewayCacheSpec)***||
| `status` | ***[StoragegatewayCacheStatus](#StoragegatewayCacheStatus)***||
## StoragegatewayCacheSpec
##### (Appears on:[StoragegatewayCache](#StoragegatewayCache), [StoragegatewayCacheStatus](#StoragegatewayCacheStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `diskID` | ***string***||
| `gatewayArn` | ***string***||
## StoragegatewayCacheStatus
##### (Appears on:[StoragegatewayCache](#StoragegatewayCache))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StoragegatewayCacheSpec](#StoragegatewayCacheSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
