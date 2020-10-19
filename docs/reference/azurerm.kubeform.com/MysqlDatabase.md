---
title: MysqlDatabase
menu:
  docs_{{ .version }}:
    identifier: mysqldatabase-azurerm.kubeform.com
    name: MysqlDatabase
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MysqlDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MysqlDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MysqlDatabaseSpec](#mysqldatabasespec)***||
| `status` | ***[MysqlDatabaseStatus](#mysqldatabasestatus)***||
## MysqlDatabaseSpec

Appears on:[MysqlDatabase](#mysqldatabase), [MysqlDatabaseStatus](#mysqldatabasestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `charset` | ***string***||
| `collation` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
## MysqlDatabaseStatus

Appears on:[MysqlDatabase](#mysqldatabase)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MysqlDatabaseSpec](#mysqldatabasespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MysqlDatabaseStatus](#mysqldatabasestatus)

---
