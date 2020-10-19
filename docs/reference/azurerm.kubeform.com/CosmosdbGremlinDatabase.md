---
title: CosmosdbGremlinDatabase
menu:
  docs_{{ .version }}:
    identifier: cosmosdbgremlindatabase-azurerm.kubeform.com
    name: CosmosdbGremlinDatabase
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CosmosdbGremlinDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CosmosdbGremlinDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CosmosdbGremlinDatabaseSpec](#cosmosdbgremlindatabasespec)***||
| `status` | ***[CosmosdbGremlinDatabaseStatus](#cosmosdbgremlindatabasestatus)***||
## CosmosdbGremlinDatabaseSpec

Appears on:[CosmosdbGremlinDatabase](#cosmosdbgremlindatabase), [CosmosdbGremlinDatabaseStatus](#cosmosdbgremlindatabasestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `throughput` | ***int64***| ***(Optional)*** |
## CosmosdbGremlinDatabaseStatus

Appears on:[CosmosdbGremlinDatabase](#cosmosdbgremlindatabase)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CosmosdbGremlinDatabaseSpec](#cosmosdbgremlindatabasespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CosmosdbGremlinDatabaseStatus](#cosmosdbgremlindatabasestatus)

---
