## RdsClusterEndpoint
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RdsClusterEndpoint` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RdsClusterEndpointSpec](#RdsClusterEndpointSpec)***||
| `status` | ***[RdsClusterEndpointStatus](#RdsClusterEndpointStatus)***||
## RdsClusterEndpointSpec
##### (Appears on:[RdsClusterEndpoint](#RdsClusterEndpoint), [RdsClusterEndpointStatus](#RdsClusterEndpointStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `clusterEndpointIdentifier` | ***string***||
| `clusterIdentifier` | ***string***||
| `customEndpointType` | ***string***||
| `endpoint` | ***string***| ***(Optional)*** |
| `excludedMembers` | ***[]string***| ***(Optional)*** |
| `staticMembers` | ***[]string***| ***(Optional)*** |
## RdsClusterEndpointStatus
##### (Appears on:[RdsClusterEndpoint](#RdsClusterEndpoint))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RdsClusterEndpointSpec](#RdsClusterEndpointSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
