---
title: CosmosdbSQLDatabase
menu:
  docs_{{ .version }}:
    identifier: cosmosdbsqldatabase-azurerm.kubeform.com
    name: CosmosdbSQLDatabase
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CosmosdbSQLDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CosmosdbSQLDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CosmosdbSQLDatabaseSpec](#cosmosdbsqldatabasespec)***||
| `status` | ***[CosmosdbSQLDatabaseStatus](#cosmosdbsqldatabasestatus)***||
## CosmosdbSQLDatabaseSpec

Appears on:[CosmosdbSQLDatabase](#cosmosdbsqldatabase), [CosmosdbSQLDatabaseStatus](#cosmosdbsqldatabasestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## CosmosdbSQLDatabaseStatus

Appears on:[CosmosdbSQLDatabase](#cosmosdbsqldatabase)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CosmosdbSQLDatabaseSpec](#cosmosdbsqldatabasespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
