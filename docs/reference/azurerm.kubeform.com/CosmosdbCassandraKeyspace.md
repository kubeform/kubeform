---
title: CosmosdbCassandraKeyspace
menu:
  docs_{{ .version }}:
    identifier: cosmosdbcassandrakeyspace-azurerm.kubeform.com
    name: CosmosdbCassandraKeyspace
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CosmosdbCassandraKeyspace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CosmosdbCassandraKeyspace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CosmosdbCassandraKeyspaceSpec](#cosmosdbcassandrakeyspacespec)***||
| `status` | ***[CosmosdbCassandraKeyspaceStatus](#cosmosdbcassandrakeyspacestatus)***||
## CosmosdbCassandraKeyspaceSpec

Appears on:[CosmosdbCassandraKeyspace](#cosmosdbcassandrakeyspace), [CosmosdbCassandraKeyspaceStatus](#cosmosdbcassandrakeyspacestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `throughput` | ***int64***| ***(Optional)*** |
## CosmosdbCassandraKeyspaceStatus

Appears on:[CosmosdbCassandraKeyspace](#cosmosdbcassandrakeyspace)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CosmosdbCassandraKeyspaceSpec](#cosmosdbcassandrakeyspacespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CosmosdbCassandraKeyspaceStatus](#cosmosdbcassandrakeyspacestatus)

---
