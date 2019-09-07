---
title: FirewallNATRuleCollection
menu:
  docs_v0.0.1:
    identifier: firewallnatrulecollection-azurerm.kubeform.com
    name: FirewallNATRuleCollection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## FirewallNATRuleCollection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `FirewallNATRuleCollection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FirewallNATRuleCollectionSpec](#FirewallNATRuleCollectionSpec)***||
| `status` | ***[FirewallNATRuleCollectionStatus](#FirewallNATRuleCollectionStatus)***||
## FirewallNATRuleCollectionSpec
##### (Appears on:[FirewallNATRuleCollection](#FirewallNATRuleCollection), [FirewallNATRuleCollectionStatus](#FirewallNATRuleCollectionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***||
| `azureFirewallName` | ***string***||
| `name` | ***string***||
| `priority` | ***int***||
| `resourceGroupName` | ***string***||
| `rule` | ***[[]FirewallNATRuleCollectionSpecRule](#FirewallNATRuleCollectionSpecRule)***||
## FirewallNATRuleCollectionSpecRule
##### (Appears on:[FirewallNATRuleCollectionSpec](#FirewallNATRuleCollectionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `destinationAddresses` | ***[]string***||
| `destinationPorts` | ***[]string***||
| `name` | ***string***||
| `protocols` | ***[]string***||
| `sourceAddresses` | ***[]string***||
| `translatedAddress` | ***string***||
| `translatedPort` | ***string***||
## FirewallNATRuleCollectionStatus
##### (Appears on:[FirewallNATRuleCollection](#FirewallNATRuleCollection))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FirewallNATRuleCollectionSpec](#FirewallNATRuleCollectionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
