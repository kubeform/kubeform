---
title: ServicebusTopicAuthorizationRule
menu:
  docs_{{ .version }}:
    identifier: servicebustopicauthorizationrule-azurerm.kubeform.com
    name: ServicebusTopicAuthorizationRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServicebusTopicAuthorizationRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ServicebusTopicAuthorizationRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServicebusTopicAuthorizationRuleSpec](#servicebustopicauthorizationrulespec)***||
| `status` | ***[ServicebusTopicAuthorizationRuleStatus](#servicebustopicauthorizationrulestatus)***||
## Phase(`string` alias)

Appears on:[ServicebusTopicAuthorizationRuleStatus](#servicebustopicauthorizationrulestatus)

## ServicebusTopicAuthorizationRuleSpec

Appears on:[ServicebusTopicAuthorizationRule](#servicebustopicauthorizationrule), [ServicebusTopicAuthorizationRuleStatus](#servicebustopicauthorizationrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `listen` | ***bool***| ***(Optional)*** |
| `manage` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `namespaceName` | ***string***||
| `resourceGroupName` | ***string***||
| `send` | ***bool***| ***(Optional)*** |
| `topicName` | ***string***||
## ServicebusTopicAuthorizationRuleStatus

Appears on:[ServicebusTopicAuthorizationRule](#servicebustopicauthorizationrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServicebusTopicAuthorizationRuleSpec](#servicebustopicauthorizationrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_connection_string` | ***string*** ||
| `primary_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
| `secondary_key` | ***string*** ||
