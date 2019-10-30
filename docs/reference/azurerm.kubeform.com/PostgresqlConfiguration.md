---
title: PostgresqlConfiguration
menu:
  docs_{{ .version }}:
    identifier: postgresqlconfiguration-azurerm.kubeform.com
    name: PostgresqlConfiguration
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PostgresqlConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PostgresqlConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PostgresqlConfigurationSpec](#postgresqlconfigurationspec)***||
| `status` | ***[PostgresqlConfigurationStatus](#postgresqlconfigurationstatus)***||
## Phase(`string` alias)

Appears on:[PostgresqlConfigurationStatus](#postgresqlconfigurationstatus)

## PostgresqlConfigurationSpec

Appears on:[PostgresqlConfiguration](#postgresqlconfiguration), [PostgresqlConfigurationStatus](#postgresqlconfigurationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `value` | ***string***||
## PostgresqlConfigurationStatus

Appears on:[PostgresqlConfiguration](#postgresqlconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PostgresqlConfigurationSpec](#postgresqlconfigurationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
