---
title: AutomationDscConfiguration
menu:
  docs_{{ .version }}:
    identifier: automationdscconfiguration-azurerm.kubeform.com
    name: AutomationDscConfiguration
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationDscConfiguration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationDscConfiguration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationDscConfigurationSpec](#automationdscconfigurationspec)***||
| `status` | ***[AutomationDscConfigurationStatus](#automationdscconfigurationstatus)***||
## AutomationDscConfigurationSpec

Appears on:[AutomationDscConfiguration](#automationdscconfiguration), [AutomationDscConfigurationStatus](#automationdscconfigurationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `automationAccountName` | ***string***||
| `contentEmbedded` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `logVerbose` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `state` | ***string***| ***(Optional)*** |
## AutomationDscConfigurationStatus

Appears on:[AutomationDscConfiguration](#automationdscconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationDscConfigurationSpec](#automationdscconfigurationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutomationDscConfigurationStatus](#automationdscconfigurationstatus)

---
