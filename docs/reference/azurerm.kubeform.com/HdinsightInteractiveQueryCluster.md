---
title: HdinsightInteractiveQueryCluster
menu:
  docs_{{ .version }}:
    identifier: hdinsightinteractivequerycluster-azurerm.kubeform.com
    name: HdinsightInteractiveQueryCluster
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## HdinsightInteractiveQueryCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HdinsightInteractiveQueryCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HdinsightInteractiveQueryClusterSpec](#HdinsightInteractiveQueryClusterSpec)***||
| `status` | ***[HdinsightInteractiveQueryClusterStatus](#HdinsightInteractiveQueryClusterStatus)***||
## HdinsightInteractiveQueryClusterSpec
##### (Appears on:[HdinsightInteractiveQueryCluster](#HdinsightInteractiveQueryCluster), [HdinsightInteractiveQueryClusterStatus](#HdinsightInteractiveQueryClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `componentVersion` | ***[[]HdinsightInteractiveQueryClusterSpecComponentVersion](#HdinsightInteractiveQueryClusterSpecComponentVersion)***||
| `gateway` | ***[[]HdinsightInteractiveQueryClusterSpecGateway](#HdinsightInteractiveQueryClusterSpecGateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightInteractiveQueryClusterSpecRoles](#HdinsightInteractiveQueryClusterSpecRoles)***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightInteractiveQueryClusterSpecStorageAccount](#HdinsightInteractiveQueryClusterSpecStorageAccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightInteractiveQueryClusterSpecComponentVersion
##### (Appears on:[HdinsightInteractiveQueryClusterSpec](#HdinsightInteractiveQueryClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `interactiveHive` | ***string***||
## HdinsightInteractiveQueryClusterSpecGateway
##### (Appears on:[HdinsightInteractiveQueryClusterSpec](#HdinsightInteractiveQueryClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightInteractiveQueryClusterSpecRoles
##### (Appears on:[HdinsightInteractiveQueryClusterSpec](#HdinsightInteractiveQueryClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `headNode` | ***[[]HdinsightInteractiveQueryClusterSpecRolesHeadNode](#HdinsightInteractiveQueryClusterSpecRolesHeadNode)***||
| `workerNode` | ***[[]HdinsightInteractiveQueryClusterSpecRolesWorkerNode](#HdinsightInteractiveQueryClusterSpecRolesWorkerNode)***||
| `zookeeperNode` | ***[[]HdinsightInteractiveQueryClusterSpecRolesZookeeperNode](#HdinsightInteractiveQueryClusterSpecRolesZookeeperNode)***||
## HdinsightInteractiveQueryClusterSpecRolesHeadNode
##### (Appears on:[HdinsightInteractiveQueryClusterSpecRoles](#HdinsightInteractiveQueryClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightInteractiveQueryClusterSpecRolesWorkerNode
##### (Appears on:[HdinsightInteractiveQueryClusterSpecRoles](#HdinsightInteractiveQueryClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `minInstanceCount` | ***int***| ***(Optional)*** |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `targetInstanceCount` | ***int***||
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightInteractiveQueryClusterSpecRolesZookeeperNode
##### (Appears on:[HdinsightInteractiveQueryClusterSpecRoles](#HdinsightInteractiveQueryClusterSpecRoles))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightInteractiveQueryClusterSpecStorageAccount
##### (Appears on:[HdinsightInteractiveQueryClusterSpec](#HdinsightInteractiveQueryClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightInteractiveQueryClusterStatus
##### (Appears on:[HdinsightInteractiveQueryCluster](#HdinsightInteractiveQueryCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightInteractiveQueryClusterSpec](#HdinsightInteractiveQueryClusterSpec)***| ***(Optional)*** |
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
