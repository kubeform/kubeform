---
title: CosmosdbCassandraKeyspace
menu:
  docs_v0.0.1:
    identifier: cosmosdbcassandrakeyspace-azurerm.kubeform.com
    name: CosmosdbCassandraKeyspace
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## CosmosdbCassandraKeyspace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CosmosdbCassandraKeyspace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CosmosdbCassandraKeyspaceSpec](#CosmosdbCassandraKeyspaceSpec)***||
| `status` | ***[CosmosdbCassandraKeyspaceStatus](#CosmosdbCassandraKeyspaceStatus)***||
## CosmosdbCassandraKeyspaceSpec
##### (Appears on:[CosmosdbCassandraKeyspace](#CosmosdbCassandraKeyspace), [CosmosdbCassandraKeyspaceStatus](#CosmosdbCassandraKeyspaceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## CosmosdbCassandraKeyspaceStatus
##### (Appears on:[CosmosdbCassandraKeyspace](#CosmosdbCassandraKeyspace))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CosmosdbCassandraKeyspaceSpec](#CosmosdbCassandraKeyspaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
