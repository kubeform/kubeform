---
title: MysqlVirtualNetworkRule
menu:
  docs_{{ .version }}:
    identifier: mysqlvirtualnetworkrule-azurerm.kubeform.com
    name: MysqlVirtualNetworkRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MysqlVirtualNetworkRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MysqlVirtualNetworkRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MysqlVirtualNetworkRuleSpec](#mysqlvirtualnetworkrulespec)***||
| `status` | ***[MysqlVirtualNetworkRuleStatus](#mysqlvirtualnetworkrulestatus)***||
## MysqlVirtualNetworkRuleSpec

Appears on:[MysqlVirtualNetworkRule](#mysqlvirtualnetworkrule), [MysqlVirtualNetworkRuleStatus](#mysqlvirtualnetworkrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serverName` | ***string***||
| `subnetID` | ***string***||
## MysqlVirtualNetworkRuleStatus

Appears on:[MysqlVirtualNetworkRule](#mysqlvirtualnetworkrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MysqlVirtualNetworkRuleSpec](#mysqlvirtualnetworkrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MysqlVirtualNetworkRuleStatus](#mysqlvirtualnetworkrulestatus)

---
