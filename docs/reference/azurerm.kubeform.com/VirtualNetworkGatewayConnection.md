---
title: VirtualNetworkGatewayConnection
menu:
  docs_{{ .version }}:
    identifier: virtualnetworkgatewayconnection-azurerm.kubeform.com
    name: VirtualNetworkGatewayConnection
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VirtualNetworkGatewayConnection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualNetworkGatewayConnection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualNetworkGatewayConnectionSpec](#VirtualNetworkGatewayConnectionSpec)***||
| `status` | ***[VirtualNetworkGatewayConnectionStatus](#VirtualNetworkGatewayConnectionStatus)***||
## VirtualNetworkGatewayConnectionSpec
##### (Appears on:[VirtualNetworkGatewayConnection](#VirtualNetworkGatewayConnection), [VirtualNetworkGatewayConnectionStatus](#VirtualNetworkGatewayConnectionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `enableBGP` | ***bool***| ***(Optional)*** |
| `expressRouteCircuitID` | ***string***| ***(Optional)*** |
| `expressRouteGatewayBypass` | ***bool***| ***(Optional)*** |
| `ipsecPolicy` | ***[[]VirtualNetworkGatewayConnectionSpecIpsecPolicy](#VirtualNetworkGatewayConnectionSpecIpsecPolicy)***| ***(Optional)*** |
| `localNetworkGatewayID` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `peerVirtualNetworkGatewayID` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `routingWeight` | ***int***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***||
| `usePolicyBasedTrafficSelectors` | ***bool***| ***(Optional)*** |
| `virtualNetworkGatewayID` | ***string***||
## VirtualNetworkGatewayConnectionSpecIpsecPolicy
##### (Appears on:[VirtualNetworkGatewayConnectionSpec](#VirtualNetworkGatewayConnectionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dhGroup` | ***string***||
| `ikeEncryption` | ***string***||
| `ikeIntegrity` | ***string***||
| `ipsecEncryption` | ***string***||
| `ipsecIntegrity` | ***string***||
| `pfsGroup` | ***string***||
| `saDatasize` | ***int***| ***(Optional)*** |
| `saLifetime` | ***int***| ***(Optional)*** |
## VirtualNetworkGatewayConnectionStatus
##### (Appears on:[VirtualNetworkGatewayConnection](#VirtualNetworkGatewayConnection))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualNetworkGatewayConnectionSpec](#VirtualNetworkGatewayConnectionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `authorization_key` | ***string*** ||
| `shared_key` | ***string*** ||
