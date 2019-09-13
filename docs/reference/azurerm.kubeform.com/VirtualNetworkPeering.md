---
title: VirtualNetworkPeering
menu:
  docs_{{ .version }}:
    identifier: virtualnetworkpeering-azurerm.kubeform.com
    name: VirtualNetworkPeering
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VirtualNetworkPeering
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualNetworkPeering` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualNetworkPeeringSpec](#VirtualNetworkPeeringSpec)***||
| `status` | ***[VirtualNetworkPeeringStatus](#VirtualNetworkPeeringStatus)***||
## VirtualNetworkPeeringSpec
##### (Appears on:[VirtualNetworkPeering](#VirtualNetworkPeering), [VirtualNetworkPeeringStatus](#VirtualNetworkPeeringStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowForwardedTraffic` | ***bool***| ***(Optional)*** |
| `allowGatewayTransit` | ***bool***| ***(Optional)*** |
| `allowVirtualNetworkAccess` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `remoteVirtualNetworkID` | ***string***||
| `resourceGroupName` | ***string***||
| `useRemoteGateways` | ***bool***| ***(Optional)*** |
| `virtualNetworkName` | ***string***||
## VirtualNetworkPeeringStatus
##### (Appears on:[VirtualNetworkPeering](#VirtualNetworkPeering))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualNetworkPeeringSpec](#VirtualNetworkPeeringSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
