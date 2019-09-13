---
title: HdinsightKafkaCluster
menu:
  docs_{{ .version }}:
    identifier: hdinsightkafkacluster-azurerm.kubeform.com
    name: HdinsightKafkaCluster
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## HdinsightKafkaCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HdinsightKafkaCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HdinsightKafkaClusterSpec](#HdinsightKafkaClusterSpec)***||
| `status` | ***[HdinsightKafkaClusterStatus](#HdinsightKafkaClusterStatus)***||
## HdinsightKafkaClusterSpec
##### (Appears on:[HdinsightKafkaCluster](#HdinsightKafkaCluster), [HdinsightKafkaClusterStatus](#HdinsightKafkaClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `componentVersion` | ***[[]HdinsightKafkaClusterSpecComponentVersion](#HdinsightKafkaClusterSpecComponentVersion)***||
| `gateway` | ***[[]HdinsightKafkaClusterSpecGateway](#HdinsightKafkaClusterSpecGateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightKafkaClusterSpecRoles](#HdinsightKafkaClusterSpecRoles)***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightKafkaClusterSpecStorageAccount](#HdinsightKafkaClusterSpecStorageAccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightKafkaClusterSpecComponentVersion
##### (Appears on:[HdinsightKafkaClusterSpec](#HdinsightKafkaClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `kafka` | ***string***||
## HdinsightKafkaClusterSpecGateway
##### (Appears on:[HdinsightKafkaClusterSpec](#HdinsightKafkaClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightKafkaClusterSpecRoles
##### (Appears on:[HdinsightKafkaClusterSpec](#HdinsightKafkaClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `headNode` | ***[[]HdinsightKafkaClusterSpecRolesHeadNode](#HdinsightKafkaClusterSpecRolesHeadNode)***||
| `workerNode` | ***[[]HdinsightKafkaClusterSpecRolesWorkerNode](#HdinsightKafkaClusterSpecRolesWorkerNode)***||
| `zookeeperNode` | ***[[]HdinsightKafkaClusterSpecRolesZookeeperNode](#HdinsightKafkaClusterSpecRolesZookeeperNode)***||
## HdinsightKafkaClusterSpecRolesHeadNode
##### (Appears on:[HdinsightKafkaClusterSpecRoles](#HdinsightKafkaClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightKafkaClusterSpecRolesWorkerNode
##### (Appears on:[HdinsightKafkaClusterSpecRoles](#HdinsightKafkaClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `minInstanceCount` | ***int***| ***(Optional)*** |
| `numberOfDisksPerNode` | ***int***||
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `targetInstanceCount` | ***int***||
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightKafkaClusterSpecRolesZookeeperNode
##### (Appears on:[HdinsightKafkaClusterSpecRoles](#HdinsightKafkaClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightKafkaClusterSpecStorageAccount
##### (Appears on:[HdinsightKafkaClusterSpec](#HdinsightKafkaClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightKafkaClusterStatus
##### (Appears on:[HdinsightKafkaCluster](#HdinsightKafkaCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightKafkaClusterSpec](#HdinsightKafkaClusterSpec)***| ***(Optional)*** |
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
