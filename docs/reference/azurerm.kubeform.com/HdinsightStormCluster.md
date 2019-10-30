---
title: HdinsightStormCluster
menu:
  docs_{{ .version }}:
    identifier: hdinsightstormcluster-azurerm.kubeform.com
    name: HdinsightStormCluster
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## HdinsightStormCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `HdinsightStormCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[HdinsightStormClusterSpec](#hdinsightstormclusterspec)***||
| `status` | ***[HdinsightStormClusterStatus](#hdinsightstormclusterstatus)***||
## HdinsightStormClusterSpec

Appears on:[HdinsightStormCluster](#hdinsightstormcluster), [HdinsightStormClusterStatus](#hdinsightstormclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `clusterVersion` | ***string***||
| `componentVersion` | ***[[]HdinsightStormClusterSpecComponentVersion](#hdinsightstormclusterspeccomponentversion)***||
| `gateway` | ***[[]HdinsightStormClusterSpecGateway](#hdinsightstormclusterspecgateway)***||
| `httpsEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `roles` | ***[[]HdinsightStormClusterSpecRoles](#hdinsightstormclusterspecroles)***||
| `sshEndpoint` | ***string***| ***(Optional)*** |
| `storageAccount` | ***[[]HdinsightStormClusterSpecStorageAccount](#hdinsightstormclusterspecstorageaccount)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***||
## HdinsightStormClusterSpecComponentVersion

Appears on:[HdinsightStormClusterSpec](#hdinsightstormclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `storm` | ***string***||
## HdinsightStormClusterSpecGateway

Appears on:[HdinsightStormClusterSpec](#hdinsightstormclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***||
| `username` | ***string***||
## HdinsightStormClusterSpecRoles

Appears on:[HdinsightStormClusterSpec](#hdinsightstormclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `headNode` | ***[[]HdinsightStormClusterSpecRolesHeadNode](#hdinsightstormclusterspecrolesheadnode)***||
| `workerNode` | ***[[]HdinsightStormClusterSpecRolesWorkerNode](#hdinsightstormclusterspecrolesworkernode)***||
| `zookeeperNode` | ***[[]HdinsightStormClusterSpecRolesZookeeperNode](#hdinsightstormclusterspecroleszookeepernode)***||
## HdinsightStormClusterSpecRolesHeadNode

Appears on:[HdinsightStormClusterSpecRoles](#hdinsightstormclusterspecroles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightStormClusterSpecRolesWorkerNode

Appears on:[HdinsightStormClusterSpecRoles](#hdinsightstormclusterspecroles)

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

Appears on:[HdinsightStormClusterSpecRoles](#hdinsightstormclusterspecroles)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `sshKeys` | ***[]string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `username` | ***string***||
| `virtualNetworkID` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
## HdinsightStormClusterSpecStorageAccount

Appears on:[HdinsightStormClusterSpec](#hdinsightstormclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `isDefault` | ***bool***||
| `storageContainerID` | ***string***||
## HdinsightStormClusterStatus

Appears on:[HdinsightStormCluster](#hdinsightstormcluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[HdinsightStormClusterSpec](#hdinsightstormclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[HdinsightStormClusterStatus](#hdinsightstormclusterstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `gateway.<index>.password` | ***string*** ||
| `roles.<index>.head_node.<index>.password` | ***string*** ||
| `roles.<index>.worker_node.<index>.password` | ***string*** ||
| `roles.<index>.zookeeper_node.<index>.password` | ***string*** ||
| `storage_account.<index>.storage_account_key` | ***string*** ||
