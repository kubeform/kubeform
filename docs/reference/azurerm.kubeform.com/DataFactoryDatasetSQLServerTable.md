---
title: DataFactoryDatasetSQLServerTable
menu:
  docs_{{ .version }}:
    identifier: datafactorydatasetsqlservertable-azurerm.kubeform.com
    name: DataFactoryDatasetSQLServerTable
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataFactoryDatasetSQLServerTable
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryDatasetSQLServerTable` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryDatasetSQLServerTableSpec](#DataFactoryDatasetSQLServerTableSpec)***||
| `status` | ***[DataFactoryDatasetSQLServerTableStatus](#DataFactoryDatasetSQLServerTableStatus)***||
## DataFactoryDatasetSQLServerTableSpec
##### (Appears on:[DataFactoryDatasetSQLServerTable](#DataFactoryDatasetSQLServerTable), [DataFactoryDatasetSQLServerTableStatus](#DataFactoryDatasetSQLServerTableStatus))
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
| `schemaColumn` | ***[[]DataFactoryDatasetSQLServerTableSpecSchemaColumn](#DataFactoryDatasetSQLServerTableSpecSchemaColumn)***| ***(Optional)*** |
| `tableName` | ***string***| ***(Optional)*** |
## DataFactoryDatasetSQLServerTableSpecSchemaColumn
##### (Appears on:[DataFactoryDatasetSQLServerTableSpec](#DataFactoryDatasetSQLServerTableSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## DataFactoryDatasetSQLServerTableStatus
##### (Appears on:[DataFactoryDatasetSQLServerTable](#DataFactoryDatasetSQLServerTable))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryDatasetSQLServerTableSpec](#DataFactoryDatasetSQLServerTableSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
