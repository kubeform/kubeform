---
title: LbOutboundRule
menu:
  docs_{{ .version }}:
    identifier: lboutboundrule-azurerm.kubeform.com
    name: LbOutboundRule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LbOutboundRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LbOutboundRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbOutboundRuleSpec](#lboutboundrulespec)***||
| `status` | ***[LbOutboundRuleStatus](#lboutboundrulestatus)***||
## LbOutboundRuleSpec

Appears on:[LbOutboundRule](#lboutboundrule), [LbOutboundRuleStatus](#lboutboundrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocatedOutboundPorts` | ***int***| ***(Optional)*** |
| `backendAddressPoolID` | ***string***||
| `enableTcpReset` | ***bool***| ***(Optional)*** |
| `frontendIPConfiguration` | ***[[]LbOutboundRuleSpecFrontendIPConfiguration](#lboutboundrulespecfrontendipconfiguration)***| ***(Optional)*** |
| `idleTimeoutInMinutes` | ***int***| ***(Optional)*** |
| `loadbalancerID` | ***string***||
| `name` | ***string***||
| `protocol` | ***string***||
| `resourceGroupName` | ***string***||
## LbOutboundRuleSpecFrontendIPConfiguration

Appears on:[LbOutboundRuleSpec](#lboutboundrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## LbOutboundRuleStatus

Appears on:[LbOutboundRule](#lboutboundrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbOutboundRuleSpec](#lboutboundrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LbOutboundRuleStatus](#lboutboundrulestatus)

---
