---
title: MariadbConfiguration
menu:
  docs_{{ .version }}:
    identifier: mariadbconfiguration-azurerm.kubeform.com
    name: MariadbConfiguration
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MariadbConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MariadbConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MariadbConfigurationSpec](#mariadbconfigurationspec)***||
| `status` | ***[MariadbConfigurationStatus](#mariadbconfigurationstatus)***||
## MariadbConfigurationSpec

Appears on:[MariadbConfiguration](#mariadbconfiguration), [MariadbConfigurationStatus](#mariadbconfigurationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `value` | ***string***||
## MariadbConfigurationStatus

Appears on:[MariadbConfiguration](#mariadbconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MariadbConfigurationSpec](#mariadbconfigurationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MariadbConfigurationStatus](#mariadbconfigurationstatus)

---
