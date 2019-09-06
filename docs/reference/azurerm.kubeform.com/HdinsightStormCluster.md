## HdinsightStormCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HdinsightStormCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HdinsightStormClusterSpec](#HdinsightStormClusterSpec)***||
| `status` | ***[HdinsightStormClusterStatus](#HdinsightStormClusterStatus)***||
## HdinsightStormClusterSpec
##### (Appears on:[HdinsightStormCluster](#HdinsightStormCluster), [HdinsightStormClusterStatus](#HdinsightStormClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `componentVersion` | ***[[]HdinsightStormClusterSpecComponentVersion](#HdinsightStormClusterSpecComponentVersion)***||
| `gateway` | ***[[]HdinsightStormClusterSpecGateway](#HdinsightStormClusterSpecGateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightStormClusterSpecRoles](#HdinsightStormClusterSpecRoles)***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightStormClusterSpecStorageAccount](#HdinsightStormClusterSpecStorageAccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightStormClusterSpecComponentVersion
##### (Appears on:[HdinsightStormClusterSpec](#HdinsightStormClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `storm` | ***string***||
## HdinsightStormClusterSpecGateway
##### (Appears on:[HdinsightStormClusterSpec](#HdinsightStormClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightStormClusterSpecRoles
##### (Appears on:[HdinsightStormClusterSpec](#HdinsightStormClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `headNode` | ***[[]HdinsightStormClusterSpecRolesHeadNode](#HdinsightStormClusterSpecRolesHeadNode)***||
| `workerNode` | ***[[]HdinsightStormClusterSpecRolesWorkerNode](#HdinsightStormClusterSpecRolesWorkerNode)***||
| `zookeeperNode` | ***[[]HdinsightStormClusterSpecRolesZookeeperNode](#HdinsightStormClusterSpecRolesZookeeperNode)***||
## HdinsightStormClusterSpecRolesHeadNode
##### (Appears on:[HdinsightStormClusterSpecRoles](#HdinsightStormClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightStormClusterSpecRolesWorkerNode
##### (Appears on:[HdinsightStormClusterSpecRoles](#HdinsightStormClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `minInstanceCount` | ***int***| ***(Optional)*** |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `targetInstanceCount` | ***int***||
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightStormClusterSpecRolesZookeeperNode
##### (Appears on:[HdinsightStormClusterSpecRoles](#HdinsightStormClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightStormClusterSpecStorageAccount
##### (Appears on:[HdinsightStormClusterSpec](#HdinsightStormClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightStormClusterStatus
##### (Appears on:[HdinsightStormCluster](#HdinsightStormCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightStormClusterSpec](#HdinsightStormClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `gateway.<index>.password` | ***string*** ||
| `roles.<index>.head_node.<index>.password` | ***string*** ||
| `roles.<index>.worker_node.<index>.password` | ***string*** ||
| `roles.<index>.zookeeper_node.<index>.password` | ***string*** ||
| `storage_account.<index>.storage_account_key` | ***string*** ||
