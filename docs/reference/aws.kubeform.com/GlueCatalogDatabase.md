---
title: GlueCatalogDatabase
menu:
  docs_{{ .version }}:
    identifier: gluecatalogdatabase-aws.kubeform.com
    name: GlueCatalogDatabase
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlueCatalogDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueCatalogDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueCatalogDatabaseSpec](#GlueCatalogDatabaseSpec)***||
| `status` | ***[GlueCatalogDatabaseStatus](#GlueCatalogDatabaseStatus)***||
## GlueCatalogDatabaseSpec
##### (Appears on:[GlueCatalogDatabase](#GlueCatalogDatabase), [GlueCatalogDatabaseStatus](#GlueCatalogDatabaseStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `catalogID` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `locationURI` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
## GlueCatalogDatabaseStatus
##### (Appears on:[GlueCatalogDatabase](#GlueCatalogDatabase))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueCatalogDatabaseSpec](#GlueCatalogDatabaseSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
