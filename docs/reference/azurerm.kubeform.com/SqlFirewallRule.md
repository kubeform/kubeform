---
title: SqlFirewallRule
menu:
  docs_v0.0.1:
    identifier: sqlfirewallrule-azurerm.kubeform.com
    name: SqlFirewallRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SqlFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlFirewallRuleSpec](#SqlFirewallRuleSpec)***||
| `status` | ***[SqlFirewallRuleStatus](#SqlFirewallRuleStatus)***||
## SqlFirewallRuleSpec
##### (Appears on:[SqlFirewallRule](#SqlFirewallRule), [SqlFirewallRuleStatus](#SqlFirewallRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `endIPAddress` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `startIPAddress` | ***string***||
## SqlFirewallRuleStatus
##### (Appears on:[SqlFirewallRule](#SqlFirewallRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlFirewallRuleSpec](#SqlFirewallRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
