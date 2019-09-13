---
title: ServicebusQueueAuthorizationRule
menu:
  docs_{{ .version }}:
    identifier: servicebusqueueauthorizationrule-azurerm.kubeform.com
    name: ServicebusQueueAuthorizationRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServicebusQueueAuthorizationRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ServicebusQueueAuthorizationRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServicebusQueueAuthorizationRuleSpec](#ServicebusQueueAuthorizationRuleSpec)***||
| `status` | ***[ServicebusQueueAuthorizationRuleStatus](#ServicebusQueueAuthorizationRuleStatus)***||
## ServicebusQueueAuthorizationRuleSpec
##### (Appears on:[ServicebusQueueAuthorizationRule](#ServicebusQueueAuthorizationRule), [ServicebusQueueAuthorizationRuleStatus](#ServicebusQueueAuthorizationRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `listen` | ***bool***| ***(Optional)*** |
| `manage` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `namespaceName` | ***string***||
| `queueName` | ***string***||
| `resourceGroupName` | ***string***||
| `send` | ***bool***| ***(Optional)*** |
## ServicebusQueueAuthorizationRuleStatus
##### (Appears on:[ServicebusQueueAuthorizationRule](#ServicebusQueueAuthorizationRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServicebusQueueAuthorizationRuleSpec](#ServicebusQueueAuthorizationRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_connection_string` | ***string*** ||
| `primary_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
| `secondary_key` | ***string*** ||
