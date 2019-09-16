---
title: MariadbServer
menu:
  docs_{{ .version }}:
    identifier: mariadbserver-azurerm.kubeform.com
    name: MariadbServer
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MariadbServer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MariadbServer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MariadbServerSpec](#MariadbServerSpec)***||
| `status` | ***[MariadbServerStatus](#MariadbServerStatus)***||
## MariadbServerSpec

Appears on:[MariadbServer](#MariadbServer), [MariadbServerStatus](#MariadbServerStatus)

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
| `sku` | ***[[]MariadbServerSpecSku](#MariadbServerSpecSku)***||
| `sslEnforcement` | ***string***||
| `storageProfile` | ***[[]MariadbServerSpecStorageProfile](#MariadbServerSpecStorageProfile)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `version` | ***string***||
## MariadbServerSpecSku

Appears on:[MariadbServerSpec](#MariadbServerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `family` | ***string***||
| `name` | ***string***||
| `tier` | ***string***||
## MariadbServerSpecStorageProfile

Appears on:[MariadbServerSpec](#MariadbServerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `backupRetentionDays` | ***int***| ***(Optional)*** |
| `geoRedundantBackup` | ***string***| ***(Optional)*** |
| `storageMb` | ***int***||
## MariadbServerStatus

Appears on:[MariadbServer](#MariadbServer)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MariadbServerSpec](#MariadbServerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `administrator_login_password` | ***string*** ||
