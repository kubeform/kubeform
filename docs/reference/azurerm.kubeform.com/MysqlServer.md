---
title: MysqlServer
menu:
  docs_{{ .version }}:
    identifier: mysqlserver-azurerm.kubeform.com
    name: MysqlServer
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MysqlServer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MysqlServer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MysqlServerSpec](#mysqlserverspec)***||
| `status` | ***[MysqlServerStatus](#mysqlserverstatus)***||
## MysqlServerSpec

Appears on:[MysqlServer](#mysqlserver), [MysqlServerStatus](#mysqlserverstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `administratorLogin` | ***string***||
| `fqdn` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***[[]MysqlServerSpecSku](#mysqlserverspecsku)***||
| `sslEnforcement` | ***string***||
| `storageProfile` | ***[[]MysqlServerSpecStorageProfile](#mysqlserverspecstorageprofile)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `version` | ***string***||
## MysqlServerSpecSku

Appears on:[MysqlServerSpec](#mysqlserverspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `family` | ***string***||
| `name` | ***string***||
| `tier` | ***string***||
## MysqlServerSpecStorageProfile

Appears on:[MysqlServerSpec](#mysqlserverspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `backupRetentionDays` | ***int***| ***(Optional)*** |
| `geoRedundantBackup` | ***string***| ***(Optional)*** |
| `storageMb` | ***int***||
## MysqlServerStatus

Appears on:[MysqlServer](#mysqlserver)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MysqlServerSpec](#mysqlserverspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `administrator_login_password` | ***string*** ||
