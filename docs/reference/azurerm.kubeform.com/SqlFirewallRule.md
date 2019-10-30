---
title: SqlFirewallRule
menu:
  docs_{{ .version }}:
    identifier: sqlfirewallrule-azurerm.kubeform.com
    name: SqlFirewallRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SqlFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlFirewallRuleSpec](#sqlfirewallrulespec)***||
| `status` | ***[SqlFirewallRuleStatus](#sqlfirewallrulestatus)***||
## Phase(`string` alias)

Appears on:[SqlFirewallRuleStatus](#sqlfirewallrulestatus)

## SqlFirewallRuleSpec

Appears on:[SqlFirewallRule](#sqlfirewallrule), [SqlFirewallRuleStatus](#sqlfirewallrulestatus)

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

Appears on:[SqlFirewallRule](#sqlfirewallrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlFirewallRuleSpec](#sqlfirewallrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
