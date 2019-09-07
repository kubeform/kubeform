---
title: ExpressRouteCircuitPeering
menu:
  docs_v0.0.1:
    identifier: expressroutecircuitpeering-azurerm.kubeform.com
    name: ExpressRouteCircuitPeering
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ExpressRouteCircuitPeering
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ExpressRouteCircuitPeering` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ExpressRouteCircuitPeeringSpec](#ExpressRouteCircuitPeeringSpec)***||
| `status` | ***[ExpressRouteCircuitPeeringStatus](#ExpressRouteCircuitPeeringStatus)***||
## ExpressRouteCircuitPeeringSpec
##### (Appears on:[ExpressRouteCircuitPeering](#ExpressRouteCircuitPeering), [ExpressRouteCircuitPeeringStatus](#ExpressRouteCircuitPeeringStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `azureAsn` | ***int***| ***(Optional)*** |
| `expressRouteCircuitName` | ***string***||
| `microsoftPeeringConfig` | ***[[]ExpressRouteCircuitPeeringSpecMicrosoftPeeringConfig](#ExpressRouteCircuitPeeringSpecMicrosoftPeeringConfig)***| ***(Optional)*** |
| `peerAsn` | ***int***| ***(Optional)*** |
| `peeringType` | ***string***||
| `primaryAzurePort` | ***string***| ***(Optional)*** |
| `primaryPeerAddressPrefix` | ***string***||
| `resourceGroupName` | ***string***||
| `secondaryAzurePort` | ***string***| ***(Optional)*** |
| `secondaryPeerAddressPrefix` | ***string***||
| `vlanID` | ***int***||
## ExpressRouteCircuitPeeringSpecMicrosoftPeeringConfig
##### (Appears on:[ExpressRouteCircuitPeeringSpec](#ExpressRouteCircuitPeeringSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `advertisedPublicPrefixes` | ***[]string***||
## ExpressRouteCircuitPeeringStatus
##### (Appears on:[ExpressRouteCircuitPeering](#ExpressRouteCircuitPeering))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ExpressRouteCircuitPeeringSpec](#ExpressRouteCircuitPeeringSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `shared_key` | ***string*** ||
