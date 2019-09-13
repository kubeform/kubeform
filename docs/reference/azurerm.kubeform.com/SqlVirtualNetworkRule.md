---
title: SqlVirtualNetworkRule
menu:
  docs_{{ .version }}:
    identifier: sqlvirtualnetworkrule-azurerm.kubeform.com
    name: SqlVirtualNetworkRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SqlVirtualNetworkRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlVirtualNetworkRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlVirtualNetworkRuleSpec](#SqlVirtualNetworkRuleSpec)***||
| `status` | ***[SqlVirtualNetworkRuleStatus](#SqlVirtualNetworkRuleStatus)***||
## SqlVirtualNetworkRuleSpec
##### (Appears on:[SqlVirtualNetworkRule](#SqlVirtualNetworkRule), [SqlVirtualNetworkRuleStatus](#SqlVirtualNetworkRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ignoreMissingVnetServiceEndpoint` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `subnetID` | ***string***||
## SqlVirtualNetworkRuleStatus
##### (Appears on:[SqlVirtualNetworkRule](#SqlVirtualNetworkRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlVirtualNetworkRuleSpec](#SqlVirtualNetworkRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
