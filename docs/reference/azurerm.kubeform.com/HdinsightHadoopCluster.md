---
title: HdinsightHadoopCluster
menu:
  docs_{{ .version }}:
    identifier: hdinsighthadoopcluster-azurerm.kubeform.com
    name: HdinsightHadoopCluster
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## HdinsightHadoopCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HdinsightHadoopCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HdinsightHadoopClusterSpec](#hdinsighthadoopclusterspec)***||
| `status` | ***[HdinsightHadoopClusterStatus](#hdinsighthadoopclusterstatus)***||
## HdinsightHadoopClusterSpec

Appears on:[HdinsightHadoopCluster](#hdinsighthadoopcluster), [HdinsightHadoopClusterStatus](#hdinsighthadoopclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `componentVersion` | ***[[]HdinsightHadoopClusterSpecComponentVersion](#hdinsighthadoopclusterspeccomponentversion)***||
| `gateway` | ***[[]HdinsightHadoopClusterSpecGateway](#hdinsighthadoopclusterspecgateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightHadoopClusterSpecRoles](#hdinsighthadoopclusterspecroles)***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightHadoopClusterSpecStorageAccount](#hdinsighthadoopclusterspecstorageaccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightHadoopClusterSpecComponentVersion

Appears on:[HdinsightHadoopClusterSpec](#hdinsighthadoopclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hadoop` | ***string***||
## HdinsightHadoopClusterSpecGateway

Appears on:[HdinsightHadoopClusterSpec](#hdinsighthadoopclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightHadoopClusterSpecRoles

Appears on:[HdinsightHadoopClusterSpec](#hdinsighthadoopclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `headNode` | ***[[]HdinsightHadoopClusterSpecRolesHeadNode](#hdinsighthadoopclusterspecrolesheadnode)***||
| `workerNode` | ***[[]HdinsightHadoopClusterSpecRolesWorkerNode](#hdinsighthadoopclusterspecrolesworkernode)***||
| `zookeeperNode` | ***[[]HdinsightHadoopClusterSpecRolesZookeeperNode](#hdinsighthadoopclusterspecroleszookeepernode)***||
## HdinsightHadoopClusterSpecRolesHeadNode

Appears on:[HdinsightHadoopClusterSpecRoles](#hdinsighthadoopclusterspecroles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightHadoopClusterSpecRolesWorkerNode

Appears on:[HdinsightHadoopClusterSpecRoles](#hdinsighthadoopclusterspecroles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `minInstanceCount` | ***int***| ***(Optional)*** |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `targetInstanceCount` | ***int***||
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightHadoopClusterSpecRolesZookeeperNode

Appears on:[HdinsightHadoopClusterSpecRoles](#hdinsighthadoopclusterspecroles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightHadoopClusterSpecStorageAccount

Appears on:[HdinsightHadoopClusterSpec](#hdinsighthadoopclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightHadoopClusterStatus

Appears on:[HdinsightHadoopCluster](#hdinsighthadoopcluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightHadoopClusterSpec](#hdinsighthadoopclusterspec)***| ***(Optional)*** |
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
