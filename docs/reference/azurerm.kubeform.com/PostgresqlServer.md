---
title: PostgresqlServer
menu:
  docs_v0.0.1:
    identifier: postgresqlserver-azurerm.kubeform.com
    name: PostgresqlServer
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## PostgresqlServer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PostgresqlServer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PostgresqlServerSpec](#PostgresqlServerSpec)***||
| `status` | ***[PostgresqlServerStatus](#PostgresqlServerStatus)***||
## PostgresqlServerSpec
##### (Appears on:[PostgresqlServer](#PostgresqlServer), [PostgresqlServerStatus](#PostgresqlServerStatus))
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
| `sku` | ***[[]PostgresqlServerSpecSku](#PostgresqlServerSpecSku)***||
| `sslEnforcement` | ***string***||
| `storageProfile` | ***[[]PostgresqlServerSpecStorageProfile](#PostgresqlServerSpecStorageProfile)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `version` | ***string***||
## PostgresqlServerSpecSku
##### (Appears on:[PostgresqlServerSpec](#PostgresqlServerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `family` | ***string***||
| `name` | ***string***||
| `tier` | ***string***||
## PostgresqlServerSpecStorageProfile
##### (Appears on:[PostgresqlServerSpec](#PostgresqlServerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `backupRetentionDays` | ***int***| ***(Optional)*** |
| `geoRedundantBackup` | ***string***| ***(Optional)*** |
| `storageMb` | ***int***||
## PostgresqlServerStatus
##### (Appears on:[PostgresqlServer](#PostgresqlServer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PostgresqlServerSpec](#PostgresqlServerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `administrator_login_password` | ***string*** ||
