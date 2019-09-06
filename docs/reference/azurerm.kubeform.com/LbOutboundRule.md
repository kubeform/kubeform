## LbOutboundRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LbOutboundRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbOutboundRuleSpec](#LbOutboundRuleSpec)***||
| `status` | ***[LbOutboundRuleStatus](#LbOutboundRuleStatus)***||
## LbOutboundRuleSpec
##### (Appears on:[LbOutboundRule](#LbOutboundRule), [LbOutboundRuleStatus](#LbOutboundRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocatedOutboundPorts` | ***int***| ***(Optional)*** |
| `backendAddressPoolID` | ***string***||
| `enableTcpReset` | ***bool***| ***(Optional)*** |
| `frontendIPConfiguration` | ***[[]LbOutboundRuleSpecFrontendIPConfiguration](#LbOutboundRuleSpecFrontendIPConfiguration)***| ***(Optional)*** |
| `idleTimeoutInMinutes` | ***int***| ***(Optional)*** |
| `loadbalancerID` | ***string***||
| `name` | ***string***||
| `protocol` | ***string***||
| `resourceGroupName` | ***string***||
## LbOutboundRuleSpecFrontendIPConfiguration
##### (Appears on:[LbOutboundRuleSpec](#LbOutboundRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## LbOutboundRuleStatus
##### (Appears on:[LbOutboundRule](#LbOutboundRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbOutboundRuleSpec](#LbOutboundRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
