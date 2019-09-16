---
title: FirewallApplicationRuleCollection
menu:
  docs_{{ .version }}:
    identifier: firewallapplicationrulecollection-azurerm.kubeform.com
    name: FirewallApplicationRuleCollection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FirewallApplicationRuleCollection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `FirewallApplicationRuleCollection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FirewallApplicationRuleCollectionSpec](#FirewallApplicationRuleCollectionSpec)***||
| `status` | ***[FirewallApplicationRuleCollectionStatus](#FirewallApplicationRuleCollectionStatus)***||
## FirewallApplicationRuleCollectionSpec

Appears on:[FirewallApplicationRuleCollection](#FirewallApplicationRuleCollection), [FirewallApplicationRuleCollectionStatus](#FirewallApplicationRuleCollectionStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `action` | ***string***||
| `azureFirewallName` | ***string***||
| `name` | ***string***||
| `priority` | ***int***||
| `resourceGroupName` | ***string***||
| `rule` | ***[[]FirewallApplicationRuleCollectionSpecRule](#FirewallApplicationRuleCollectionSpecRule)***||
## FirewallApplicationRuleCollectionSpecRule

Appears on:[FirewallApplicationRuleCollectionSpec](#FirewallApplicationRuleCollectionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `fqdnTags` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `protocol` | ***[[]FirewallApplicationRuleCollectionSpecRuleProtocol](#FirewallApplicationRuleCollectionSpecRuleProtocol)***| ***(Optional)*** |
| `sourceAddresses` | ***[]string***||
| `targetFqdns` | ***[]string***| ***(Optional)*** |
## FirewallApplicationRuleCollectionSpecRuleProtocol

Appears on:[FirewallApplicationRuleCollectionSpecRule](#FirewallApplicationRuleCollectionSpecRule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int***| ***(Optional)*** |
| `type` | ***string***||
## FirewallApplicationRuleCollectionStatus

Appears on:[FirewallApplicationRuleCollection](#FirewallApplicationRuleCollection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FirewallApplicationRuleCollectionSpec](#FirewallApplicationRuleCollectionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
