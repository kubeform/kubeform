## DynamodbGlobalTable
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DynamodbGlobalTable` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DynamodbGlobalTableSpec](#DynamodbGlobalTableSpec)***||
| `status` | ***[DynamodbGlobalTableStatus](#DynamodbGlobalTableStatus)***||
## DynamodbGlobalTableSpec
##### (Appears on:[DynamodbGlobalTable](#DynamodbGlobalTable), [DynamodbGlobalTableStatus](#DynamodbGlobalTableStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `replica` | ***[[]DynamodbGlobalTableSpecReplica](#DynamodbGlobalTableSpecReplica)***||
## DynamodbGlobalTableSpecReplica
##### (Appears on:[DynamodbGlobalTableSpec](#DynamodbGlobalTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `regionName` | ***string***||
## DynamodbGlobalTableStatus
##### (Appears on:[DynamodbGlobalTable](#DynamodbGlobalTable))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DynamodbGlobalTableSpec](#DynamodbGlobalTableSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
