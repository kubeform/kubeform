---
title: MssqlElasticpool
menu:
  docs_{{ .version }}:
    identifier: mssqlelasticpool-azurerm.kubeform.com
    name: MssqlElasticpool
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MssqlElasticpool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MssqlElasticpool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MssqlElasticpoolSpec](#mssqlelasticpoolspec)***||
| `status` | ***[MssqlElasticpoolStatus](#mssqlelasticpoolstatus)***||
## MssqlElasticpoolSpec

Appears on:[MssqlElasticpool](#mssqlelasticpool), [MssqlElasticpoolStatus](#mssqlelasticpoolstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `elasticPoolProperties` | ***[[]MssqlElasticpoolSpecElasticPoolProperties](#mssqlelasticpoolspecelasticpoolproperties)***| ***(Optional)*** Deprecated|
| `location` | ***string***||
| `maxSizeBytes` | ***int***| ***(Optional)*** |
| `maxSizeGb` | ***encoding/json.Number***| ***(Optional)*** |
| `name` | ***string***||
| `perDatabaseSettings` | ***[[]MssqlElasticpoolSpecPerDatabaseSettings](#mssqlelasticpoolspecperdatabasesettings)***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `sku` | ***[[]MssqlElasticpoolSpecSku](#mssqlelasticpoolspecsku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zoneRedundant` | ***bool***| ***(Optional)*** |
## MssqlElasticpoolSpecElasticPoolProperties

Appears on:[MssqlElasticpoolSpec](#mssqlelasticpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `creationDate` | ***string***| ***(Optional)*** Deprecated|
| `licenseType` | ***string***| ***(Optional)*** Deprecated|
| `maxSizeBytes` | ***int***| ***(Optional)*** Deprecated|
| `state` | ***string***| ***(Optional)*** Deprecated|
| `zoneRedundant` | ***bool***| ***(Optional)*** Deprecated|
## MssqlElasticpoolSpecPerDatabaseSettings

Appears on:[MssqlElasticpoolSpec](#mssqlelasticpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxCapacity` | ***encoding/json.Number***||
| `minCapacity` | ***encoding/json.Number***||
## MssqlElasticpoolSpecSku

Appears on:[MssqlElasticpoolSpec](#mssqlelasticpoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `family` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `tier` | ***string***||
## MssqlElasticpoolStatus

Appears on:[MssqlElasticpool](#mssqlelasticpool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MssqlElasticpoolSpec](#mssqlelasticpoolspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
