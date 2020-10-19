---
title: VpnGateway
menu:
  docs_{{ .version }}:
    identifier: vpngateway-azurerm.kubeform.com
    name: VpnGateway
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpnGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VpnGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpnGatewaySpec](#vpngatewayspec)***||
| `status` | ***[VpnGatewayStatus](#vpngatewaystatus)***||
## Phase(`string` alias)

Appears on:[VpnGatewayStatus](#vpngatewaystatus)

## VpnGatewaySpec

Appears on:[VpnGateway](#vpngateway), [VpnGatewayStatus](#vpngatewaystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bgpSettings` | ***[[]VpnGatewaySpecBgpSettings](#vpngatewayspecbgpsettings)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `scaleUnit` | ***int64***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `virtualHubID` | ***string***||
## VpnGatewaySpecBgpSettings

Appears on:[VpnGatewaySpec](#vpngatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `asn` | ***int64***||
| `bgpPeeringAddress` | ***string***| ***(Optional)*** |
| `peerWeight` | ***int64***||
## VpnGatewayStatus

Appears on:[VpnGateway](#vpngateway)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpnGatewaySpec](#vpngatewayspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
