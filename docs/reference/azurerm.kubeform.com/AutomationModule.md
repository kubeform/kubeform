---
title: AutomationModule
menu:
  docs_{{ .version }}:
    identifier: automationmodule-azurerm.kubeform.com
    name: AutomationModule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationModule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationModule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationModuleSpec](#AutomationModuleSpec)***||
| `status` | ***[AutomationModuleStatus](#AutomationModuleStatus)***||
## AutomationModuleSpec
##### (Appears on:[AutomationModule](#AutomationModule), [AutomationModuleStatus](#AutomationModuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `automationAccountName` | ***string***||
| `moduleLink` | ***[[]AutomationModuleSpecModuleLink](#AutomationModuleSpecModuleLink)***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## AutomationModuleSpecModuleLink
##### (Appears on:[AutomationModuleSpec](#AutomationModuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `hash` | ***[[]AutomationModuleSpecModuleLinkHash](#AutomationModuleSpecModuleLinkHash)***| ***(Optional)*** |
| `uri` | ***string***||
## AutomationModuleSpecModuleLinkHash
##### (Appears on:[AutomationModuleSpecModuleLink](#AutomationModuleSpecModuleLink))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `algorithm` | ***string***||
| `value` | ***string***||
## AutomationModuleStatus
##### (Appears on:[AutomationModule](#AutomationModule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationModuleSpec](#AutomationModuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
