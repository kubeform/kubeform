---
title: AutomationRunbook
menu:
  docs_{{ .version }}:
    identifier: automationrunbook-azurerm.kubeform.com
    name: AutomationRunbook
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationRunbook
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationRunbook` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationRunbookSpec](#automationrunbookspec)***||
| `status` | ***[AutomationRunbookStatus](#automationrunbookstatus)***||
## AutomationRunbookSpec

Appears on:[AutomationRunbook](#automationrunbook), [AutomationRunbookStatus](#automationrunbookstatus)

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
| `publishContentLink` | ***[[]AutomationRunbookSpecPublishContentLink](#automationrunbookspecpublishcontentlink)***||
| `resourceGroupName` | ***string***||
| `runbookType` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## AutomationRunbookSpecPublishContentLink

Appears on:[AutomationRunbookSpec](#automationrunbookspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hash` | ***[[]AutomationRunbookSpecPublishContentLinkHash](#automationrunbookspecpublishcontentlinkhash)***| ***(Optional)*** |
| `uri` | ***string***||
| `version` | ***string***| ***(Optional)*** |
## AutomationRunbookSpecPublishContentLinkHash

Appears on:[AutomationRunbookSpecPublishContentLink](#automationrunbookspecpublishcontentlink)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `algorithm` | ***string***||
| `value` | ***string***||
## AutomationRunbookStatus

Appears on:[AutomationRunbook](#automationrunbook)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationRunbookSpec](#automationrunbookspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AutomationRunbookStatus](#automationrunbookstatus)

---
