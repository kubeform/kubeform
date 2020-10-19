---
title: DataFactoryLinkedServicePostgresql
menu:
  docs_{{ .version }}:
    identifier: datafactorylinkedservicepostgresql-azurerm.kubeform.com
    name: DataFactoryLinkedServicePostgresql
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataFactoryLinkedServicePostgresql
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryLinkedServicePostgresql` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryLinkedServicePostgresqlSpec](#datafactorylinkedservicepostgresqlspec)***||
| `status` | ***[DataFactoryLinkedServicePostgresqlStatus](#datafactorylinkedservicepostgresqlstatus)***||
## DataFactoryLinkedServicePostgresqlSpec

Appears on:[DataFactoryLinkedServicePostgresql](#datafactorylinkedservicepostgresql), [DataFactoryLinkedServicePostgresqlStatus](#datafactorylinkedservicepostgresqlstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `additionalProperties` | ***map[string]string***| ***(Optional)*** |
| `annotations` | ***[]string***| ***(Optional)*** |
| `connectionString` | ***string***||
| `dataFactoryName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `integrationRuntimeName` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
## DataFactoryLinkedServicePostgresqlStatus

Appears on:[DataFactoryLinkedServicePostgresql](#datafactorylinkedservicepostgresql)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryLinkedServicePostgresqlSpec](#datafactorylinkedservicepostgresqlspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DataFactoryLinkedServicePostgresqlStatus](#datafactorylinkedservicepostgresqlstatus)

---
