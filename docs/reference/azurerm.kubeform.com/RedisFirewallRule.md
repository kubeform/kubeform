---
title: RedisFirewallRule
menu:
  docs_{{ .version }}:
    identifier: redisfirewallrule-azurerm.kubeform.com
    name: RedisFirewallRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RedisFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `RedisFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RedisFirewallRuleSpec](#redisfirewallrulespec)***||
| `status` | ***[RedisFirewallRuleStatus](#redisfirewallrulestatus)***||
## Phase(`string` alias)

Appears on:[RedisFirewallRuleStatus](#redisfirewallrulestatus)

## RedisFirewallRuleSpec

Appears on:[RedisFirewallRule](#redisfirewallrule), [RedisFirewallRuleStatus](#redisfirewallrulestatus)

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

Appears on:[RedisFirewallRule](#redisfirewallrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RedisFirewallRuleSpec](#redisfirewallrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
