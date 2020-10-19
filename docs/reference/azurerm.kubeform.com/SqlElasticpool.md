---
title: SqlElasticpool
menu:
  docs_{{ .version }}:
    identifier: sqlelasticpool-azurerm.kubeform.com
    name: SqlElasticpool
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SqlElasticpool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlElasticpool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlElasticpoolSpec](#sqlelasticpoolspec)***||
| `status` | ***[SqlElasticpoolStatus](#sqlelasticpoolstatus)***||
## Phase(`string` alias)

Appears on:[SqlElasticpoolStatus](#sqlelasticpoolstatus)

## SqlElasticpoolSpec

Appears on:[SqlElasticpool](#sqlelasticpool), [SqlElasticpoolStatus](#sqlelasticpoolstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `creationDate` | ***string***| ***(Optional)*** |
| `dbDtuMax` | ***int64***| ***(Optional)*** |
| `dbDtuMin` | ***int64***| ***(Optional)*** |
| `dtu` | ***int64***||
| `edition` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `poolSize` | ***int64***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## SqlElasticpoolStatus

Appears on:[SqlElasticpool](#sqlelasticpool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlElasticpoolSpec](#sqlelasticpoolspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
