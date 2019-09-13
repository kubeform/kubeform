---
title: HdinsightHbaseCluster
menu:
  docs_{{ .version }}:
    identifier: hdinsighthbasecluster-azurerm.kubeform.com
    name: HdinsightHbaseCluster
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## HdinsightHbaseCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HdinsightHbaseCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HdinsightHbaseClusterSpec](#HdinsightHbaseClusterSpec)***||
| `status` | ***[HdinsightHbaseClusterStatus](#HdinsightHbaseClusterStatus)***||
## HdinsightHbaseClusterSpec
##### (Appears on:[HdinsightHbaseCluster](#HdinsightHbaseCluster), [HdinsightHbaseClusterStatus](#HdinsightHbaseClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `componentVersion` | ***[[]HdinsightHbaseClusterSpecComponentVersion](#HdinsightHbaseClusterSpecComponentVersion)***||
| `gateway` | ***[[]HdinsightHbaseClusterSpecGateway](#HdinsightHbaseClusterSpecGateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightHbaseClusterSpecRoles](#HdinsightHbaseClusterSpecRoles)***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightHbaseClusterSpecStorageAccount](#HdinsightHbaseClusterSpecStorageAccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightHbaseClusterSpecComponentVersion
##### (Appears on:[HdinsightHbaseClusterSpec](#HdinsightHbaseClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `hbase` | ***string***||
## HdinsightHbaseClusterSpecGateway
##### (Appears on:[HdinsightHbaseClusterSpec](#HdinsightHbaseClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightHbaseClusterSpecRoles
##### (Appears on:[HdinsightHbaseClusterSpec](#HdinsightHbaseClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `headNode` | ***[[]HdinsightHbaseClusterSpecRolesHeadNode](#HdinsightHbaseClusterSpecRolesHeadNode)***||
| `workerNode` | ***[[]HdinsightHbaseClusterSpecRolesWorkerNode](#HdinsightHbaseClusterSpecRolesWorkerNode)***||
| `zookeeperNode` | ***[[]HdinsightHbaseClusterSpecRolesZookeeperNode](#HdinsightHbaseClusterSpecRolesZookeeperNode)***||
## HdinsightHbaseClusterSpecRolesHeadNode
##### (Appears on:[HdinsightHbaseClusterSpecRoles](#HdinsightHbaseClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightHbaseClusterSpecRolesWorkerNode
##### (Appears on:[HdinsightHbaseClusterSpecRoles](#HdinsightHbaseClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `minInstanceCount` | ***int***| ***(Optional)*** |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `targetInstanceCount` | ***int***||
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightHbaseClusterSpecRolesZookeeperNode
##### (Appears on:[HdinsightHbaseClusterSpecRoles](#HdinsightHbaseClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightHbaseClusterSpecStorageAccount
##### (Appears on:[HdinsightHbaseClusterSpec](#HdinsightHbaseClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightHbaseClusterStatus
##### (Appears on:[HdinsightHbaseCluster](#HdinsightHbaseCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightHbaseClusterSpec](#HdinsightHbaseClusterSpec)***| ***(Optional)*** |
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
