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
| `spec` | ***[HdinsightHbaseClusterSpec](#hdinsighthbaseclusterspec)***||
| `status` | ***[HdinsightHbaseClusterStatus](#hdinsighthbaseclusterstatus)***||
## HdinsightHbaseClusterSpec

Appears on:[HdinsightHbaseCluster](#hdinsighthbasecluster), [HdinsightHbaseClusterStatus](#hdinsighthbaseclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `componentVersion` | ***[[]HdinsightHbaseClusterSpecComponentVersion](#hdinsighthbaseclusterspeccomponentversion)***||
| `gateway` | ***[[]HdinsightHbaseClusterSpecGateway](#hdinsighthbaseclusterspecgateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightHbaseClusterSpecRoles](#hdinsighthbaseclusterspecroles)***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightHbaseClusterSpecStorageAccount](#hdinsighthbaseclusterspecstorageaccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightHbaseClusterSpecComponentVersion

Appears on:[HdinsightHbaseClusterSpec](#hdinsighthbaseclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hbase` | ***string***||
## HdinsightHbaseClusterSpecGateway

Appears on:[HdinsightHbaseClusterSpec](#hdinsighthbaseclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightHbaseClusterSpecRoles

Appears on:[HdinsightHbaseClusterSpec](#hdinsighthbaseclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `headNode` | ***[[]HdinsightHbaseClusterSpecRolesHeadNode](#hdinsighthbaseclusterspecrolesheadnode)***||
| `workerNode` | ***[[]HdinsightHbaseClusterSpecRolesWorkerNode](#hdinsighthbaseclusterspecrolesworkernode)***||
| `zookeeperNode` | ***[[]HdinsightHbaseClusterSpecRolesZookeeperNode](#hdinsighthbaseclusterspecroleszookeepernode)***||
## HdinsightHbaseClusterSpecRolesHeadNode

Appears on:[HdinsightHbaseClusterSpecRoles](#hdinsighthbaseclusterspecroles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightHbaseClusterSpecRolesWorkerNode

Appears on:[HdinsightHbaseClusterSpecRoles](#hdinsighthbaseclusterspecroles)

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

Appears on:[HdinsightHbaseClusterSpecRoles](#hdinsighthbaseclusterspecroles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightHbaseClusterSpecStorageAccount

Appears on:[HdinsightHbaseClusterSpec](#hdinsighthbaseclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightHbaseClusterStatus

Appears on:[HdinsightHbaseCluster](#hdinsighthbasecluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightHbaseClusterSpec](#hdinsighthbaseclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[HdinsightHbaseClusterStatus](#hdinsighthbaseclusterstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `gateway.<index>.password` | ***string*** ||
| `roles.<index>.head_node.<index>.password` | ***string*** ||
| `roles.<index>.worker_node.<index>.password` | ***string*** ||
| `roles.<index>.zookeeper_node.<index>.password` | ***string*** ||
| `storage_account.<index>.storage_account_key` | ***string*** ||
