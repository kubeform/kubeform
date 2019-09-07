---
title: DataFactoryLinkedServiceSQLServer
menu:
  docs_v0.0.1:
    identifier: datafactorylinkedservicesqlserver-azurerm.kubeform.com
    name: DataFactoryLinkedServiceSQLServer
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DataFactoryLinkedServiceSQLServer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryLinkedServiceSQLServer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryLinkedServiceSQLServerSpec](#DataFactoryLinkedServiceSQLServerSpec)***||
| `status` | ***[DataFactoryLinkedServiceSQLServerStatus](#DataFactoryLinkedServiceSQLServerStatus)***||
## DataFactoryLinkedServiceSQLServerSpec
##### (Appears on:[DataFactoryLinkedServiceSQLServer](#DataFactoryLinkedServiceSQLServer), [DataFactoryLinkedServiceSQLServerStatus](#DataFactoryLinkedServiceSQLServerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
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
## DataFactoryLinkedServiceSQLServerStatus
##### (Appears on:[DataFactoryLinkedServiceSQLServer](#DataFactoryLinkedServiceSQLServer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryLinkedServiceSQLServerSpec](#DataFactoryLinkedServiceSQLServerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
