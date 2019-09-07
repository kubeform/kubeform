---
title: DataFactoryLinkedServiceDataLakeStorageGen2
menu:
  docs_v0.0.1:
    identifier: datafactorylinkedservicedatalakestoragegen2-azurerm.kubeform.com
    name: DataFactoryLinkedServiceDataLakeStorageGen2
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DataFactoryLinkedServiceDataLakeStorageGen2
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryLinkedServiceDataLakeStorageGen2` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryLinkedServiceDataLakeStorageGen2Spec](#DataFactoryLinkedServiceDataLakeStorageGen2Spec)***||
| `status` | ***[DataFactoryLinkedServiceDataLakeStorageGen2Status](#DataFactoryLinkedServiceDataLakeStorageGen2Status)***||
## DataFactoryLinkedServiceDataLakeStorageGen2Spec
##### (Appears on:[DataFactoryLinkedServiceDataLakeStorageGen2](#DataFactoryLinkedServiceDataLakeStorageGen2), [DataFactoryLinkedServiceDataLakeStorageGen2Status](#DataFactoryLinkedServiceDataLakeStorageGen2Status))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `additionalProperties` | ***map[string]string***| ***(Optional)*** |
| `annotations` | ***[]string***| ***(Optional)*** |
| `dataFactoryName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `integrationRuntimeName` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `servicePrincipalID` | ***string***||
| `servicePrincipalKey` | ***string***||
| `tenant` | ***string***||
| `url` | ***string***||
## DataFactoryLinkedServiceDataLakeStorageGen2Status
##### (Appears on:[DataFactoryLinkedServiceDataLakeStorageGen2](#DataFactoryLinkedServiceDataLakeStorageGen2))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryLinkedServiceDataLakeStorageGen2Spec](#DataFactoryLinkedServiceDataLakeStorageGen2Spec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
