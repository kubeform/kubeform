---
title: BotChannelSlack
menu:
  docs_{{ .version }}:
    identifier: botchannelslack-azurerm.kubeform.com
    name: BotChannelSlack
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BotChannelSlack
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BotChannelSlack` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BotChannelSlackSpec](#botchannelslackspec)***||
| `status` | ***[BotChannelSlackStatus](#botchannelslackstatus)***||
## BotChannelSlackSpec

Appears on:[BotChannelSlack](#botchannelslack), [BotChannelSlackStatus](#botchannelslackstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `botName` | ***string***||
| `clientID` | ***string***||
| `landingPageURL` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `resourceGroupName` | ***string***||
## BotChannelSlackStatus

Appears on:[BotChannelSlack](#botchannelslack)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BotChannelSlackSpec](#botchannelslackspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BotChannelSlackStatus](#botchannelslackstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `client_secret` | ***string*** ||
| `verification_token` | ***string*** ||
