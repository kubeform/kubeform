---
title: DataFactoryDatasetPostgresql
menu:
  docs_{{ .version }}:
    identifier: datafactorydatasetpostgresql-azurerm.kubeform.com
    name: DataFactoryDatasetPostgresql
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataFactoryDatasetPostgresql
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryDatasetPostgresql` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryDatasetPostgresqlSpec](#DataFactoryDatasetPostgresqlSpec)***||
| `status` | ***[DataFactoryDatasetPostgresqlStatus](#DataFactoryDatasetPostgresqlStatus)***||
## DataFactoryDatasetPostgresqlSpec
##### (Appears on:[DataFactoryDatasetPostgresql](#DataFactoryDatasetPostgresql), [DataFactoryDatasetPostgresqlStatus](#DataFactoryDatasetPostgresqlStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `additionalProperties` | ***map[string]string***| ***(Optional)*** |
| `annotations` | ***[]string***| ***(Optional)*** |
| `dataFactoryName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `folder` | ***string***| ***(Optional)*** |
| `linkedServiceName` | ***string***||
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `schemaColumn` | ***[[]DataFactoryDatasetPostgresqlSpecSchemaColumn](#DataFactoryDatasetPostgresqlSpecSchemaColumn)***| ***(Optional)*** |
| `tableName` | ***string***| ***(Optional)*** |
## DataFactoryDatasetPostgresqlSpecSchemaColumn
##### (Appears on:[DataFactoryDatasetPostgresqlSpec](#DataFactoryDatasetPostgresqlSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## DataFactoryDatasetPostgresqlStatus
##### (Appears on:[DataFactoryDatasetPostgresql](#DataFactoryDatasetPostgresql))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryDatasetPostgresqlSpec](#DataFactoryDatasetPostgresqlSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
