## BigtableInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BigtableInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BigtableInstanceSpec](#BigtableInstanceSpec)***||
| `status` | ***[BigtableInstanceStatus](#BigtableInstanceStatus)***||
## BigtableInstanceSpec
##### (Appears on:[BigtableInstance](#BigtableInstance), [BigtableInstanceStatus](#BigtableInstanceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cluster` | ***[[]BigtableInstanceSpecCluster](#BigtableInstanceSpecCluster)***| ***(Optional)*** |
| `clusterID` | ***string***| ***(Optional)*** Deprecated|
| `displayName` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `numNodes` | ***int***| ***(Optional)*** Deprecated|
| `project` | ***string***| ***(Optional)*** |
| `storageType` | ***string***| ***(Optional)*** Deprecated|
| `zone` | ***string***| ***(Optional)*** Deprecated|
## BigtableInstanceSpecCluster
##### (Appears on:[BigtableInstanceSpec](#BigtableInstanceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clusterID` | ***string***| ***(Optional)*** |
| `numNodes` | ***int***| ***(Optional)*** |
| `storageType` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## BigtableInstanceStatus
##### (Appears on:[BigtableInstance](#BigtableInstance))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BigtableInstanceSpec](#BigtableInstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
