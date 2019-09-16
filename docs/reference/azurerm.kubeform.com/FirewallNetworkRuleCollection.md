---
title: FirewallNetworkRuleCollection
menu:
  docs_{{ .version }}:
    identifier: firewallnetworkrulecollection-azurerm.kubeform.com
    name: FirewallNetworkRuleCollection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FirewallNetworkRuleCollection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `FirewallNetworkRuleCollection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FirewallNetworkRuleCollectionSpec](#FirewallNetworkRuleCollectionSpec)***||
| `status` | ***[FirewallNetworkRuleCollectionStatus](#FirewallNetworkRuleCollectionStatus)***||
## FirewallNetworkRuleCollectionSpec

Appears on:[FirewallNetworkRuleCollection](#FirewallNetworkRuleCollection), [FirewallNetworkRuleCollectionStatus](#FirewallNetworkRuleCollectionStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***||
| `azureFirewallName` | ***string***||
| `name` | ***string***||
| `priority` | ***int***||
| `resourceGroupName` | ***string***||
| `rule` | ***[[]FirewallNetworkRuleCollectionSpecRule](#FirewallNetworkRuleCollectionSpecRule)***||
## FirewallNetworkRuleCollectionSpecRule

Appears on:[FirewallNetworkRuleCollectionSpec](#FirewallNetworkRuleCollectionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `destinationAddresses` | ***[]string***||
| `destinationPorts` | ***[]string***||
| `name` | ***string***||
| `protocols` | ***[]string***||
| `sourceAddresses` | ***[]string***||
## FirewallNetworkRuleCollectionStatus

Appears on:[FirewallNetworkRuleCollection](#FirewallNetworkRuleCollection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FirewallNetworkRuleCollectionSpec](#FirewallNetworkRuleCollectionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
