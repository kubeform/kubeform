## HdinsightMlServicesCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HdinsightMlServicesCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HdinsightMlServicesClusterSpec](#HdinsightMlServicesClusterSpec)***||
| `status` | ***[HdinsightMlServicesClusterStatus](#HdinsightMlServicesClusterStatus)***||
## HdinsightMlServicesClusterSpec
##### (Appears on:[HdinsightMlServicesCluster](#HdinsightMlServicesCluster), [HdinsightMlServicesClusterStatus](#HdinsightMlServicesClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `edgeSSHEndpoint` | ***string***| ***(Optional)*** |
| `gateway` | ***[[]HdinsightMlServicesClusterSpecGateway](#HdinsightMlServicesClusterSpecGateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightMlServicesClusterSpecRoles](#HdinsightMlServicesClusterSpecRoles)***||
| `rstudio` | ***bool***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightMlServicesClusterSpecStorageAccount](#HdinsightMlServicesClusterSpecStorageAccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightMlServicesClusterSpecGateway
##### (Appears on:[HdinsightMlServicesClusterSpec](#HdinsightMlServicesClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightMlServicesClusterSpecRoles
##### (Appears on:[HdinsightMlServicesClusterSpec](#HdinsightMlServicesClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `edgeNode` | ***[[]HdinsightMlServicesClusterSpecRolesEdgeNode](#HdinsightMlServicesClusterSpecRolesEdgeNode)***||
| `headNode` | ***[[]HdinsightMlServicesClusterSpecRolesHeadNode](#HdinsightMlServicesClusterSpecRolesHeadNode)***||
| `workerNode` | ***[[]HdinsightMlServicesClusterSpecRolesWorkerNode](#HdinsightMlServicesClusterSpecRolesWorkerNode)***||
| `zookeeperNode` | ***[[]HdinsightMlServicesClusterSpecRolesZookeeperNode](#HdinsightMlServicesClusterSpecRolesZookeeperNode)***||
## HdinsightMlServicesClusterSpecRolesEdgeNode
##### (Appears on:[HdinsightMlServicesClusterSpecRoles](#HdinsightMlServicesClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightMlServicesClusterSpecRolesHeadNode
##### (Appears on:[HdinsightMlServicesClusterSpecRoles](#HdinsightMlServicesClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightMlServicesClusterSpecRolesWorkerNode
##### (Appears on:[HdinsightMlServicesClusterSpecRoles](#HdinsightMlServicesClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `minInstanceCount` | ***int***| ***(Optional)*** |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `targetInstanceCount` | ***int***||
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightMlServicesClusterSpecRolesZookeeperNode
##### (Appears on:[HdinsightMlServicesClusterSpecRoles](#HdinsightMlServicesClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightMlServicesClusterSpecStorageAccount
##### (Appears on:[HdinsightMlServicesClusterSpec](#HdinsightMlServicesClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightMlServicesClusterStatus
##### (Appears on:[HdinsightMlServicesCluster](#HdinsightMlServicesCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightMlServicesClusterSpec](#HdinsightMlServicesClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `gateway.<index>.password` | ***string*** ||
| `roles.<index>.edge_node.<index>.password` | ***string*** ||
| `roles.<index>.head_node.<index>.password` | ***string*** ||
| `roles.<index>.worker_node.<index>.password` | ***string*** ||
| `roles.<index>.zookeeper_node.<index>.password` | ***string*** ||
| `storage_account.<index>.storage_account_key` | ***string*** ||
