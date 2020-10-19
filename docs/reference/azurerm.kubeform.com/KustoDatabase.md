---
title: KustoDatabase
menu:
  docs_{{ .version }}:
    identifier: kustodatabase-azurerm.kubeform.com
    name: KustoDatabase
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KustoDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KustoDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KustoDatabaseSpec](#kustodatabasespec)***||
| `status` | ***[KustoDatabaseStatus](#kustodatabasestatus)***||
## KustoDatabaseSpec

Appears on:[KustoDatabase](#kustodatabase), [KustoDatabaseStatus](#kustodatabasestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `clusterName` | ***string***||
| `hotCachePeriod` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `size` | ***float64***| ***(Optional)*** |
| `softDeletePeriod` | ***string***| ***(Optional)*** |
## KustoDatabaseStatus

Appears on:[KustoDatabase](#kustodatabase)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KustoDatabaseSpec](#kustodatabasespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KustoDatabaseStatus](#kustodatabasestatus)

---
