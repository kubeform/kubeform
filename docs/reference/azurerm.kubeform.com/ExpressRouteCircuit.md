---
title: ExpressRouteCircuit
menu:
  docs_{{ .version }}:
    identifier: expressroutecircuit-azurerm.kubeform.com
    name: ExpressRouteCircuit
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ExpressRouteCircuit
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ExpressRouteCircuit` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ExpressRouteCircuitSpec](#ExpressRouteCircuitSpec)***||
| `status` | ***[ExpressRouteCircuitStatus](#ExpressRouteCircuitStatus)***||
## ExpressRouteCircuitSpec
##### (Appears on:[ExpressRouteCircuit](#ExpressRouteCircuit), [ExpressRouteCircuitStatus](#ExpressRouteCircuitStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `allowClassicOperations` | ***bool***| ***(Optional)*** |
| `bandwidthInMbps` | ***int***||
| `location` | ***string***||
| `name` | ***string***||
| `peeringLocation` | ***string***||
| `resourceGroupName` | ***string***||
| `serviceProviderName` | ***string***||
| `serviceProviderProvisioningState` | ***string***| ***(Optional)*** |
| `sku` | ***[[]ExpressRouteCircuitSpecSku](#ExpressRouteCircuitSpecSku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## ExpressRouteCircuitSpecSku
##### (Appears on:[ExpressRouteCircuitSpec](#ExpressRouteCircuitSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `family` | ***string***||
| `tier` | ***string***||
## ExpressRouteCircuitStatus
##### (Appears on:[ExpressRouteCircuit](#ExpressRouteCircuit))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ExpressRouteCircuitSpec](#ExpressRouteCircuitSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `service_key` | ***string*** ||
