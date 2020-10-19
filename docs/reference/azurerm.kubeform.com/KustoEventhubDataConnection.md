---
title: KustoEventhubDataConnection
menu:
  docs_{{ .version }}:
    identifier: kustoeventhubdataconnection-azurerm.kubeform.com
    name: KustoEventhubDataConnection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KustoEventhubDataConnection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KustoEventhubDataConnection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KustoEventhubDataConnectionSpec](#kustoeventhubdataconnectionspec)***||
| `status` | ***[KustoEventhubDataConnectionStatus](#kustoeventhubdataconnectionstatus)***||
## KustoEventhubDataConnectionSpec

Appears on:[KustoEventhubDataConnection](#kustoeventhubdataconnection), [KustoEventhubDataConnectionStatus](#kustoeventhubdataconnectionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `clusterName` | ***string***||
| `consumerGroup` | ***string***||
| `dataFormat` | ***string***| ***(Optional)*** |
| `databaseName` | ***string***||
| `eventhubID` | ***string***||
| `location` | ***string***||
| `mappingRuleName` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tableName` | ***string***| ***(Optional)*** |
## KustoEventhubDataConnectionStatus

Appears on:[KustoEventhubDataConnection](#kustoeventhubdataconnection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KustoEventhubDataConnectionSpec](#kustoeventhubdataconnectionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KustoEventhubDataConnectionStatus](#kustoeventhubdataconnectionstatus)

---
