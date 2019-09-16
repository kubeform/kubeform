---
title: DataLakeStoreFirewallRule
menu:
  docs_{{ .version }}:
    identifier: datalakestorefirewallrule-azurerm.kubeform.com
    name: DataLakeStoreFirewallRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataLakeStoreFirewallRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataLakeStoreFirewallRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataLakeStoreFirewallRuleSpec](#DataLakeStoreFirewallRuleSpec)***||
| `status` | ***[DataLakeStoreFirewallRuleStatus](#DataLakeStoreFirewallRuleStatus)***||
## DataLakeStoreFirewallRuleSpec

Appears on:[DataLakeStoreFirewallRule](#DataLakeStoreFirewallRule), [DataLakeStoreFirewallRuleStatus](#DataLakeStoreFirewallRuleStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `endIPAddress` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `startIPAddress` | ***string***||
## DataLakeStoreFirewallRuleStatus

Appears on:[DataLakeStoreFirewallRule](#DataLakeStoreFirewallRule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataLakeStoreFirewallRuleSpec](#DataLakeStoreFirewallRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
