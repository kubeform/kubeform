---
title: LbNATRule
menu:
  docs_v0.0.1:
    identifier: lbnatrule-azurerm.kubeform.com
    name: LbNATRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## LbNATRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LbNATRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbNATRuleSpec](#LbNATRuleSpec)***||
| `status` | ***[LbNATRuleStatus](#LbNATRuleStatus)***||
## LbNATRuleSpec
##### (Appears on:[LbNATRule](#LbNATRule), [LbNATRuleStatus](#LbNATRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backendIPConfigurationID` | ***string***| ***(Optional)*** |
| `backendPort` | ***int***||
| `enableFloatingIP` | ***bool***| ***(Optional)*** |
| `frontendIPConfigurationID` | ***string***| ***(Optional)*** |
| `frontendIPConfigurationName` | ***string***||
| `frontendPort` | ***int***||
| `loadbalancerID` | ***string***||
| `location` | ***string***| ***(Optional)*** Deprecated|
| `name` | ***string***||
| `protocol` | ***string***||
| `resourceGroupName` | ***string***||
## LbNATRuleStatus
##### (Appears on:[LbNATRule](#LbNATRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbNATRuleSpec](#LbNATRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
