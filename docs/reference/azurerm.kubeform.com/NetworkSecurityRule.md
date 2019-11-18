---
title: NetworkSecurityRule
menu:
  docs_{{ .version }}:
    identifier: networksecurityrule-azurerm.kubeform.com
    name: NetworkSecurityRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkSecurityRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkSecurityRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkSecurityRuleSpec](#networksecurityrulespec)***||
| `status` | ***[NetworkSecurityRuleStatus](#networksecurityrulestatus)***||
## NetworkSecurityRuleSpec

Appears on:[NetworkSecurityRule](#networksecurityrule), [NetworkSecurityRuleStatus](#networksecurityrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `access` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `destinationAddressPrefix` | ***string***| ***(Optional)*** |
| `destinationAddressPrefixes` | ***[]string***| ***(Optional)*** |
| `destinationApplicationSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `destinationPortRange` | ***string***| ***(Optional)*** |
| `destinationPortRanges` | ***[]string***| ***(Optional)*** |
| `direction` | ***string***||
| `name` | ***string***||
| `networkSecurityGroupName` | ***string***||
| `priority` | ***int64***||
| `protocol` | ***string***||
| `resourceGroupName` | ***string***||
| `sourceAddressPrefix` | ***string***| ***(Optional)*** |
| `sourceAddressPrefixes` | ***[]string***| ***(Optional)*** |
| `sourceApplicationSecurityGroupIDS` | ***[]string***| ***(Optional)*** |
| `sourcePortRange` | ***string***| ***(Optional)*** |
| `sourcePortRanges` | ***[]string***| ***(Optional)*** |
## NetworkSecurityRuleStatus

Appears on:[NetworkSecurityRule](#networksecurityrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkSecurityRuleSpec](#networksecurityrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NetworkSecurityRuleStatus](#networksecurityrulestatus)

---
