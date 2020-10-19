---
title: FirewallNATRuleCollection
menu:
  docs_{{ .version }}:
    identifier: firewallnatrulecollection-azurerm.kubeform.com
    name: FirewallNATRuleCollection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FirewallNATRuleCollection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `FirewallNATRuleCollection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FirewallNATRuleCollectionSpec](#firewallnatrulecollectionspec)***||
| `status` | ***[FirewallNATRuleCollectionStatus](#firewallnatrulecollectionstatus)***||
## FirewallNATRuleCollectionSpec

Appears on:[FirewallNATRuleCollection](#firewallnatrulecollection), [FirewallNATRuleCollectionStatus](#firewallnatrulecollectionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***||
| `azureFirewallName` | ***string***||
| `name` | ***string***||
| `priority` | ***int64***||
| `resourceGroupName` | ***string***||
| `rule` | ***[[]FirewallNATRuleCollectionSpecRule](#firewallnatrulecollectionspecrule)***||
## FirewallNATRuleCollectionSpecRule

Appears on:[FirewallNATRuleCollectionSpec](#firewallnatrulecollectionspec)

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

Appears on:[FirewallNATRuleCollection](#firewallnatrulecollection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FirewallNATRuleCollectionSpec](#firewallnatrulecollectionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[FirewallNATRuleCollectionStatus](#firewallnatrulecollectionstatus)

---
