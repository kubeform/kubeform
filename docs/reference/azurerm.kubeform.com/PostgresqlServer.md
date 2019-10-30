---
title: PostgresqlServer
menu:
  docs_{{ .version }}:
    identifier: postgresqlserver-azurerm.kubeform.com
    name: PostgresqlServer
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PostgresqlServer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PostgresqlServer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PostgresqlServerSpec](#postgresqlserverspec)***||
| `status` | ***[PostgresqlServerStatus](#postgresqlserverstatus)***||
## Phase(`string` alias)

Appears on:[PostgresqlServerStatus](#postgresqlserverstatus)

## PostgresqlServerSpec

Appears on:[PostgresqlServer](#postgresqlserver), [PostgresqlServerStatus](#postgresqlserverstatus)

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
| `sku` | ***[[]PostgresqlServerSpecSku](#postgresqlserverspecsku)***||
| `sslEnforcement` | ***string***||
| `storageProfile` | ***[[]PostgresqlServerSpecStorageProfile](#postgresqlserverspecstorageprofile)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `version` | ***string***||
## PostgresqlServerSpecSku

Appears on:[PostgresqlServerSpec](#postgresqlserverspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `family` | ***string***||
| `name` | ***string***||
| `tier` | ***string***||
## PostgresqlServerSpecStorageProfile

Appears on:[PostgresqlServerSpec](#postgresqlserverspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `backupRetentionDays` | ***int***| ***(Optional)*** |
| `geoRedundantBackup` | ***string***| ***(Optional)*** |
| `storageMb` | ***int***||
## PostgresqlServerStatus

Appears on:[PostgresqlServer](#postgresqlserver)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PostgresqlServerSpec](#postgresqlserverspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `administrator_login_password` | ***string*** ||
