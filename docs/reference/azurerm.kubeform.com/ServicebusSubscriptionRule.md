---
title: ServicebusSubscriptionRule
menu:
  docs_v0.0.1:
    identifier: servicebussubscriptionrule-azurerm.kubeform.com
    name: ServicebusSubscriptionRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ServicebusSubscriptionRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ServicebusSubscriptionRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServicebusSubscriptionRuleSpec](#ServicebusSubscriptionRuleSpec)***||
| `status` | ***[ServicebusSubscriptionRuleStatus](#ServicebusSubscriptionRuleStatus)***||
## ServicebusSubscriptionRuleSpec
##### (Appears on:[ServicebusSubscriptionRule](#ServicebusSubscriptionRule), [ServicebusSubscriptionRuleStatus](#ServicebusSubscriptionRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***| ***(Optional)*** |
| `correlationFilter` | ***[[]ServicebusSubscriptionRuleSpecCorrelationFilter](#ServicebusSubscriptionRuleSpecCorrelationFilter)***| ***(Optional)*** |
| `filterType` | ***string***||
| `name` | ***string***||
| `namespaceName` | ***string***||
| `resourceGroupName` | ***string***||
| `sqlFilter` | ***string***| ***(Optional)*** |
| `subscriptionName` | ***string***||
| `topicName` | ***string***||
## ServicebusSubscriptionRuleSpecCorrelationFilter
##### (Appears on:[ServicebusSubscriptionRuleSpec](#ServicebusSubscriptionRuleSpec))
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
##### (Appears on:[ServicebusSubscriptionRule](#ServicebusSubscriptionRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServicebusSubscriptionRuleSpec](#ServicebusSubscriptionRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
