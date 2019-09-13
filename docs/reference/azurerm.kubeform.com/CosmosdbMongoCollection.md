---
title: CosmosdbMongoCollection
menu:
  docs_{{ .version }}:
    identifier: cosmosdbmongocollection-azurerm.kubeform.com
    name: CosmosdbMongoCollection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CosmosdbMongoCollection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CosmosdbMongoCollection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CosmosdbMongoCollectionSpec](#CosmosdbMongoCollectionSpec)***||
| `status` | ***[CosmosdbMongoCollectionStatus](#CosmosdbMongoCollectionStatus)***||
## CosmosdbMongoCollectionSpec
##### (Appears on:[CosmosdbMongoCollection](#CosmosdbMongoCollection), [CosmosdbMongoCollectionStatus](#CosmosdbMongoCollectionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `databaseName` | ***string***||
| `defaultTtlSeconds` | ***int***| ***(Optional)*** |
| `indexes` | ***[[]CosmosdbMongoCollectionSpecIndexes](#CosmosdbMongoCollectionSpecIndexes)***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `shardKey` | ***string***| ***(Optional)*** |
## CosmosdbMongoCollectionSpecIndexes
##### (Appears on:[CosmosdbMongoCollectionSpec](#CosmosdbMongoCollectionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `unique` | ***bool***| ***(Optional)*** |
## CosmosdbMongoCollectionStatus
##### (Appears on:[CosmosdbMongoCollection](#CosmosdbMongoCollection))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CosmosdbMongoCollectionSpec](#CosmosdbMongoCollectionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
