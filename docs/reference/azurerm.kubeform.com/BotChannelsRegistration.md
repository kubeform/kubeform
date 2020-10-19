---
title: BotChannelsRegistration
menu:
  docs_{{ .version }}:
    identifier: botchannelsregistration-azurerm.kubeform.com
    name: BotChannelsRegistration
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BotChannelsRegistration
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `BotChannelsRegistration` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BotChannelsRegistrationSpec](#botchannelsregistrationspec)***||
| `status` | ***[BotChannelsRegistrationStatus](#botchannelsregistrationstatus)***||
## BotChannelsRegistrationSpec

Appears on:[BotChannelsRegistration](#botchannelsregistration), [BotChannelsRegistrationStatus](#botchannelsregistrationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `developerAppInsightsApplicationID` | ***string***| ***(Optional)*** |
| `developerAppInsightsKey` | ***string***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `microsoftAppID` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## BotChannelsRegistrationStatus

Appears on:[BotChannelsRegistration](#botchannelsregistration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BotChannelsRegistrationSpec](#botchannelsregistrationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BotChannelsRegistrationStatus](#botchannelsregistrationstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `developer_app_insights_api_key` | ***string*** ||
