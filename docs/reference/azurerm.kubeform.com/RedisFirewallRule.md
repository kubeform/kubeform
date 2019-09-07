---
title: RedisFirewallRule
menu:
  docs_v0.0.1:
    identifier: redisfirewallrule-azurerm.kubeform.com
    name: RedisFirewallRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## RedisFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RedisFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RedisFirewallRuleSpec](#RedisFirewallRuleSpec)***||
| `status` | ***[RedisFirewallRuleStatus](#RedisFirewallRuleStatus)***||
## RedisFirewallRuleSpec
##### (Appears on:[RedisFirewallRule](#RedisFirewallRule), [RedisFirewallRuleStatus](#RedisFirewallRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `endIP` | ***string***||
| `name` | ***string***||
| `redisCacheName` | ***string***||
| `resourceGroupName` | ***string***||
| `startIP` | ***string***||
## RedisFirewallRuleStatus
##### (Appears on:[RedisFirewallRule](#RedisFirewallRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RedisFirewallRuleSpec](#RedisFirewallRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
