---
title: VirtualNetworkGateway
menu:
  docs_{{ .version }}:
    identifier: virtualnetworkgateway-azurerm.kubeform.com
    name: VirtualNetworkGateway
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VirtualNetworkGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualNetworkGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualNetworkGatewaySpec](#virtualnetworkgatewayspec)***||
| `status` | ***[VirtualNetworkGatewayStatus](#virtualnetworkgatewaystatus)***||
## Phase(`string` alias)

Appears on:[VirtualNetworkGatewayStatus](#virtualnetworkgatewaystatus)

## VirtualNetworkGatewaySpec

Appears on:[VirtualNetworkGateway](#virtualnetworkgateway), [VirtualNetworkGatewayStatus](#virtualnetworkgatewaystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activeActive` | ***bool***| ***(Optional)*** |
| `bgpSettings` | ***[[]VirtualNetworkGatewaySpecBgpSettings](#virtualnetworkgatewayspecbgpsettings)***| ***(Optional)*** |
| `defaultLocalNetworkGatewayID` | ***string***| ***(Optional)*** |
| `enableBGP` | ***bool***| ***(Optional)*** |
| `generation` | ***string***| ***(Optional)*** |
| `ipConfiguration` | ***[[]VirtualNetworkGatewaySpecIpConfiguration](#virtualnetworkgatewayspecipconfiguration)***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***||
| `vpnClientConfiguration` | ***[[]VirtualNetworkGatewaySpecVpnClientConfiguration](#virtualnetworkgatewayspecvpnclientconfiguration)***| ***(Optional)*** |
| `vpnType` | ***string***| ***(Optional)*** |
## VirtualNetworkGatewaySpecBgpSettings

Appears on:[VirtualNetworkGatewaySpec](#virtualnetworkgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `asn` | ***int64***| ***(Optional)*** |
| `peerWeight` | ***int64***| ***(Optional)*** |
| `peeringAddress` | ***string***| ***(Optional)*** |
## VirtualNetworkGatewaySpecIpConfiguration

Appears on:[VirtualNetworkGatewaySpec](#virtualnetworkgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `privateIPAddressAllocation` | ***string***| ***(Optional)*** |
| `publicIPAddressID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***||
## VirtualNetworkGatewaySpecVpnClientConfiguration

Appears on:[VirtualNetworkGatewaySpec](#virtualnetworkgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `addressSpace` | ***[]string***||
| `radiusServerAddress` | ***string***| ***(Optional)*** |
| `radiusServerSecret` | ***string***| ***(Optional)*** |
| `revokedCertificate` | ***[[]VirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate](#virtualnetworkgatewayspecvpnclientconfigurationrevokedcertificate)***| ***(Optional)*** |
| `rootCertificate` | ***[[]VirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate](#virtualnetworkgatewayspecvpnclientconfigurationrootcertificate)***| ***(Optional)*** |
| `vpnClientProtocols` | ***[]string***| ***(Optional)*** |
## VirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate

Appears on:[VirtualNetworkGatewaySpecVpnClientConfiguration](#virtualnetworkgatewayspecvpnclientconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `thumbprint` | ***string***||
## VirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate

Appears on:[VirtualNetworkGatewaySpecVpnClientConfiguration](#virtualnetworkgatewayspecvpnclientconfiguration)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `publicCertData` | ***string***||
## VirtualNetworkGatewayStatus

Appears on:[VirtualNetworkGateway](#virtualnetworkgateway)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualNetworkGatewaySpec](#virtualnetworkgatewayspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
