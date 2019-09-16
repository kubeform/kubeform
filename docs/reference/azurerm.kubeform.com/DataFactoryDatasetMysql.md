---
title: DataFactoryDatasetMysql
menu:
  docs_{{ .version }}:
    identifier: datafactorydatasetmysql-azurerm.kubeform.com
    name: DataFactoryDatasetMysql
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataFactoryDatasetMysql
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryDatasetMysql` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryDatasetMysqlSpec](#DataFactoryDatasetMysqlSpec)***||
| `status` | ***[DataFactoryDatasetMysqlStatus](#DataFactoryDatasetMysqlStatus)***||
## DataFactoryDatasetMysqlSpec

Appears on:[DataFactoryDatasetMysql](#DataFactoryDatasetMysql), [DataFactoryDatasetMysqlStatus](#DataFactoryDatasetMysqlStatus)

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
| `schemaColumn` | ***[[]DataFactoryDatasetMysqlSpecSchemaColumn](#DataFactoryDatasetMysqlSpecSchemaColumn)***| ***(Optional)*** |
| `tableName` | ***string***| ***(Optional)*** |
## DataFactoryDatasetMysqlSpecSchemaColumn

Appears on:[DataFactoryDatasetMysqlSpec](#DataFactoryDatasetMysqlSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## DataFactoryDatasetMysqlStatus

Appears on:[DataFactoryDatasetMysql](#DataFactoryDatasetMysql)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryDatasetMysqlSpec](#DataFactoryDatasetMysqlSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
