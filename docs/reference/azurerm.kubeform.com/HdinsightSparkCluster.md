---
title: HdinsightSparkCluster
menu:
  docs_{{ .version }}:
    identifier: hdinsightsparkcluster-azurerm.kubeform.com
    name: HdinsightSparkCluster
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## HdinsightSparkCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HdinsightSparkCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HdinsightSparkClusterSpec](#HdinsightSparkClusterSpec)***||
| `status` | ***[HdinsightSparkClusterStatus](#HdinsightSparkClusterStatus)***||
## HdinsightSparkClusterSpec

Appears on:[HdinsightSparkCluster](#HdinsightSparkCluster), [HdinsightSparkClusterStatus](#HdinsightSparkClusterStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `componentVersion` | ***[[]HdinsightSparkClusterSpecComponentVersion](#HdinsightSparkClusterSpecComponentVersion)***||
| `gateway` | ***[[]HdinsightSparkClusterSpecGateway](#HdinsightSparkClusterSpecGateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightSparkClusterSpecRoles](#HdinsightSparkClusterSpecRoles)***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightSparkClusterSpecStorageAccount](#HdinsightSparkClusterSpecStorageAccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightSparkClusterSpecComponentVersion

Appears on:[HdinsightSparkClusterSpec](#HdinsightSparkClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `spark` | ***string***||
## HdinsightSparkClusterSpecGateway

Appears on:[HdinsightSparkClusterSpec](#HdinsightSparkClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightSparkClusterSpecRoles

Appears on:[HdinsightSparkClusterSpec](#HdinsightSparkClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `headNode` | ***[[]HdinsightSparkClusterSpecRolesHeadNode](#HdinsightSparkClusterSpecRolesHeadNode)***||
| `workerNode` | ***[[]HdinsightSparkClusterSpecRolesWorkerNode](#HdinsightSparkClusterSpecRolesWorkerNode)***||
| `zookeeperNode` | ***[[]HdinsightSparkClusterSpecRolesZookeeperNode](#HdinsightSparkClusterSpecRolesZookeeperNode)***||
## HdinsightSparkClusterSpecRolesHeadNode

Appears on:[HdinsightSparkClusterSpecRoles](#HdinsightSparkClusterSpecRoles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightSparkClusterSpecRolesWorkerNode

Appears on:[HdinsightSparkClusterSpecRoles](#HdinsightSparkClusterSpecRoles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `minInstanceCount` | ***int***| ***(Optional)*** |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `targetInstanceCount` | ***int***||
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightSparkClusterSpecRolesZookeeperNode

Appears on:[HdinsightSparkClusterSpecRoles](#HdinsightSparkClusterSpecRoles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightSparkClusterSpecStorageAccount

Appears on:[HdinsightSparkClusterSpec](#HdinsightSparkClusterSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightSparkClusterStatus

Appears on:[HdinsightSparkCluster](#HdinsightSparkCluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightSparkClusterSpec](#HdinsightSparkClusterSpec)***| ***(Optional)*** |
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
