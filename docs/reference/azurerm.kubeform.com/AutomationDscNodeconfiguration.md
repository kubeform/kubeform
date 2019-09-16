---
title: AutomationDscNodeconfiguration
menu:
  docs_{{ .version }}:
    identifier: automationdscnodeconfiguration-azurerm.kubeform.com
    name: AutomationDscNodeconfiguration
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationDscNodeconfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationDscNodeconfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationDscNodeconfigurationSpec](#AutomationDscNodeconfigurationSpec)***||
| `status` | ***[AutomationDscNodeconfigurationStatus](#AutomationDscNodeconfigurationStatus)***||
## AutomationDscNodeconfigurationSpec

Appears on:[AutomationDscNodeconfiguration](#AutomationDscNodeconfiguration), [AutomationDscNodeconfigurationStatus](#AutomationDscNodeconfigurationStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `automationAccountName` | ***string***||
| `configurationName` | ***string***| ***(Optional)*** |
| `contentEmbedded` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## AutomationDscNodeconfigurationStatus

Appears on:[AutomationDscNodeconfiguration](#AutomationDscNodeconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationDscNodeconfigurationSpec](#AutomationDscNodeconfigurationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
