---
title: AutomationRunbook
menu:
  docs_v0.0.1:
    identifier: automationrunbook-azurerm.kubeform.com
    name: AutomationRunbook
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AutomationRunbook
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationRunbook` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationRunbookSpec](#AutomationRunbookSpec)***||
| `status` | ***[AutomationRunbookStatus](#AutomationRunbookStatus)***||
## AutomationRunbookSpec
##### (Appears on:[AutomationRunbook](#AutomationRunbook), [AutomationRunbookStatus](#AutomationRunbookStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `content` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `logProgress` | ***bool***||
| `logVerbose` | ***bool***||
| `name` | ***string***||
| `publishContentLink` | ***[[]AutomationRunbookSpecPublishContentLink](#AutomationRunbookSpecPublishContentLink)***||
| `resourceGroupName` | ***string***||
| `runbookType` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## AutomationRunbookSpecPublishContentLink
##### (Appears on:[AutomationRunbookSpec](#AutomationRunbookSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `hash` | ***[[]AutomationRunbookSpecPublishContentLinkHash](#AutomationRunbookSpecPublishContentLinkHash)***| ***(Optional)*** |
| `uri` | ***string***||
| `version` | ***string***| ***(Optional)*** |
## AutomationRunbookSpecPublishContentLinkHash
##### (Appears on:[AutomationRunbookSpecPublishContentLink](#AutomationRunbookSpecPublishContentLink))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `algorithm` | ***string***||
| `value` | ***string***||
## AutomationRunbookStatus
##### (Appears on:[AutomationRunbook](#AutomationRunbook))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationRunbookSpec](#AutomationRunbookSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
