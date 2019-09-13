---
title: MariadbDatabase
menu:
  docs_{{ .version }}:
    identifier: mariadbdatabase-azurerm.kubeform.com
    name: MariadbDatabase
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MariadbDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MariadbDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MariadbDatabaseSpec](#MariadbDatabaseSpec)***||
| `status` | ***[MariadbDatabaseStatus](#MariadbDatabaseStatus)***||
## MariadbDatabaseSpec
##### (Appears on:[MariadbDatabase](#MariadbDatabase), [MariadbDatabaseStatus](#MariadbDatabaseStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `charset` | ***string***||
| `collation` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
## MariadbDatabaseStatus
##### (Appears on:[MariadbDatabase](#MariadbDatabase))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MariadbDatabaseSpec](#MariadbDatabaseSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
