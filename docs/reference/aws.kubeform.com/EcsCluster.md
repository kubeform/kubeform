## EcsCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `EcsCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EcsClusterSpec](#EcsClusterSpec)***||
| `status` | ***[EcsClusterStatus](#EcsClusterStatus)***||
## EcsClusterSpec
##### (Appears on:[EcsCluster](#EcsCluster), [EcsClusterStatus](#EcsClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## EcsClusterStatus
##### (Appears on:[EcsCluster](#EcsCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EcsClusterSpec](#EcsClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
