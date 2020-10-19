---
title: BotChannelEmail
menu:
  docs_{{ .version }}:
    identifier: botchannelemail-azurerm.kubeform.com
    name: BotChannelEmail
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BotChannelEmail
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BotChannelEmail` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BotChannelEmailSpec](#botchannelemailspec)***||
| `status` | ***[BotChannelEmailStatus](#botchannelemailstatus)***||
## BotChannelEmailSpec

Appears on:[BotChannelEmail](#botchannelemail), [BotChannelEmailStatus](#botchannelemailstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `botName` | ***string***||
| `emailAddress` | ***string***||
| `location` | ***string***||
| `resourceGroupName` | ***string***||
## BotChannelEmailStatus

Appears on:[BotChannelEmail](#botchannelemail)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BotChannelEmailSpec](#botchannelemailspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BotChannelEmailStatus](#botchannelemailstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `email_password` | ***string*** ||
