---
title: CosmosdbMongoDatabase
menu:
  docs_v0.0.1:
    identifier: cosmosdbmongodatabase-azurerm.kubeform.com
    name: CosmosdbMongoDatabase
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## CosmosdbMongoDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CosmosdbMongoDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CosmosdbMongoDatabaseSpec](#CosmosdbMongoDatabaseSpec)***||
| `status` | ***[CosmosdbMongoDatabaseStatus](#CosmosdbMongoDatabaseStatus)***||
## CosmosdbMongoDatabaseSpec
##### (Appears on:[CosmosdbMongoDatabase](#CosmosdbMongoDatabase), [CosmosdbMongoDatabaseStatus](#CosmosdbMongoDatabaseStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## CosmosdbMongoDatabaseStatus
##### (Appears on:[CosmosdbMongoDatabase](#CosmosdbMongoDatabase))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CosmosdbMongoDatabaseSpec](#CosmosdbMongoDatabaseSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
