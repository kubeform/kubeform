---
title: PostgresqlFirewallRule
menu:
  docs_v0.0.1:
    identifier: postgresqlfirewallrule-azurerm.kubeform.com
    name: PostgresqlFirewallRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## PostgresqlFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PostgresqlFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PostgresqlFirewallRuleSpec](#PostgresqlFirewallRuleSpec)***||
| `status` | ***[PostgresqlFirewallRuleStatus](#PostgresqlFirewallRuleStatus)***||
## PostgresqlFirewallRuleSpec
##### (Appears on:[PostgresqlFirewallRule](#PostgresqlFirewallRule), [PostgresqlFirewallRuleStatus](#PostgresqlFirewallRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `endIPAddress` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `startIPAddress` | ***string***||
## PostgresqlFirewallRuleStatus
##### (Appears on:[PostgresqlFirewallRule](#PostgresqlFirewallRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PostgresqlFirewallRuleSpec](#PostgresqlFirewallRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
