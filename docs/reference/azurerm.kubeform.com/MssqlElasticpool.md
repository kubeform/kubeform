---
title: MssqlElasticpool
menu:
  docs_v0.0.1:
    identifier: mssqlelasticpool-azurerm.kubeform.com
    name: MssqlElasticpool
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## MssqlElasticpool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MssqlElasticpool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MssqlElasticpoolSpec](#MssqlElasticpoolSpec)***||
| `status` | ***[MssqlElasticpoolStatus](#MssqlElasticpoolStatus)***||
## MssqlElasticpoolSpec
##### (Appears on:[MssqlElasticpool](#MssqlElasticpool), [MssqlElasticpoolStatus](#MssqlElasticpoolStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `elasticPoolProperties` | ***[[]MssqlElasticpoolSpecElasticPoolProperties](#MssqlElasticpoolSpecElasticPoolProperties)***| ***(Optional)*** Deprecated|
| `location` | ***string***||
| `maxSizeBytes` | ***int***| ***(Optional)*** |
| `maxSizeGb` | ***encoding/json.Number***| ***(Optional)*** |
| `name` | ***string***||
| `perDatabaseSettings` | ***[[]MssqlElasticpoolSpecPerDatabaseSettings](#MssqlElasticpoolSpecPerDatabaseSettings)***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `sku` | ***[[]MssqlElasticpoolSpecSku](#MssqlElasticpoolSpecSku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zoneRedundant` | ***bool***| ***(Optional)*** |
## MssqlElasticpoolSpecElasticPoolProperties
##### (Appears on:[MssqlElasticpoolSpec](#MssqlElasticpoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `creationDate` | ***string***| ***(Optional)*** Deprecated|
| `licenseType` | ***string***| ***(Optional)*** Deprecated|
| `maxSizeBytes` | ***int***| ***(Optional)*** Deprecated|
| `state` | ***string***| ***(Optional)*** Deprecated|
| `zoneRedundant` | ***bool***| ***(Optional)*** Deprecated|
## MssqlElasticpoolSpecPerDatabaseSettings
##### (Appears on:[MssqlElasticpoolSpec](#MssqlElasticpoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `maxCapacity` | ***encoding/json.Number***||
| `minCapacity` | ***encoding/json.Number***||
## MssqlElasticpoolSpecSku
##### (Appears on:[MssqlElasticpoolSpec](#MssqlElasticpoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `family` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `tier` | ***string***||
## MssqlElasticpoolStatus
##### (Appears on:[MssqlElasticpool](#MssqlElasticpool))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MssqlElasticpoolSpec](#MssqlElasticpoolSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
