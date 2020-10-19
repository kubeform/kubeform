---
title: ServicebusSubscriptionRule
menu:
  docs_{{ .version }}:
    identifier: servicebussubscriptionrule-azurerm.kubeform.com
    name: ServicebusSubscriptionRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServicebusSubscriptionRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ServicebusSubscriptionRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServicebusSubscriptionRuleSpec](#servicebussubscriptionrulespec)***||
| `status` | ***[ServicebusSubscriptionRuleStatus](#servicebussubscriptionrulestatus)***||
## Phase(`string` alias)

Appears on:[ServicebusSubscriptionRuleStatus](#servicebussubscriptionrulestatus)

## ServicebusSubscriptionRuleSpec

Appears on:[ServicebusSubscriptionRule](#servicebussubscriptionrule), [ServicebusSubscriptionRuleStatus](#servicebussubscriptionrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***| ***(Optional)*** |
| `correlationFilter` | ***[[]ServicebusSubscriptionRuleSpecCorrelationFilter](#servicebussubscriptionrulespeccorrelationfilter)***| ***(Optional)*** |
| `filterType` | ***string***||
| `name` | ***string***||
| `namespaceName` | ***string***||
| `resourceGroupName` | ***string***||
| `sqlFilter` | ***string***| ***(Optional)*** |
| `subscriptionName` | ***string***||
| `topicName` | ***string***||
## ServicebusSubscriptionRuleSpecCorrelationFilter

Appears on:[ServicebusSubscriptionRuleSpec](#servicebussubscriptionrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***| ***(Optional)*** |
| `correlationID` | ***string***| ***(Optional)*** |
| `label` | ***string***| ***(Optional)*** |
| `messageID` | ***string***| ***(Optional)*** |
| `replyTo` | ***string***| ***(Optional)*** |
| `replyToSessionID` | ***string***| ***(Optional)*** |
| `sessionID` | ***string***| ***(Optional)*** |
| `to` | ***string***| ***(Optional)*** |
## ServicebusSubscriptionRuleStatus

Appears on:[ServicebusSubscriptionRule](#servicebussubscriptionrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServicebusSubscriptionRuleSpec](#servicebussubscriptionrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
