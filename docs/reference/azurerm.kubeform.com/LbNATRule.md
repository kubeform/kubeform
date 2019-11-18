---
title: LbNATRule
menu:
  docs_{{ .version }}:
    identifier: lbnatrule-azurerm.kubeform.com
    name: LbNATRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LbNATRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LbNATRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbNATRuleSpec](#lbnatrulespec)***||
| `status` | ***[LbNATRuleStatus](#lbnatrulestatus)***||
## LbNATRuleSpec

Appears on:[LbNATRule](#lbnatrule), [LbNATRuleStatus](#lbnatrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backendIPConfigurationID` | ***string***| ***(Optional)*** |
| `backendPort` | ***int64***||
| `enableFloatingIP` | ***bool***| ***(Optional)*** |
| `frontendIPConfigurationID` | ***string***| ***(Optional)*** |
| `frontendIPConfigurationName` | ***string***||
| `frontendPort` | ***int64***||
| `loadbalancerID` | ***string***||
| `location` | ***string***| ***(Optional)*** Deprecated|
| `name` | ***string***||
| `protocol` | ***string***||
| `resourceGroupName` | ***string***||
## LbNATRuleStatus

Appears on:[LbNATRule](#lbnatrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbNATRuleSpec](#lbnatrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LbNATRuleStatus](#lbnatrulestatus)

---
