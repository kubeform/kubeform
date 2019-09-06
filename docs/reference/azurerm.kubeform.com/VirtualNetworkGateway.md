## VirtualNetworkGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualNetworkGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualNetworkGatewaySpec](#VirtualNetworkGatewaySpec)***||
| `status` | ***[VirtualNetworkGatewayStatus](#VirtualNetworkGatewayStatus)***||
## VirtualNetworkGatewaySpec
##### (Appears on:[VirtualNetworkGateway](#VirtualNetworkGateway), [VirtualNetworkGatewayStatus](#VirtualNetworkGatewayStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activeActive` | ***bool***| ***(Optional)*** |
| `bgpSettings` | ***[[]VirtualNetworkGatewaySpecBgpSettings](#VirtualNetworkGatewaySpecBgpSettings)***| ***(Optional)*** |
| `defaultLocalNetworkGatewayID` | ***string***| ***(Optional)*** |
| `enableBGP` | ***bool***| ***(Optional)*** |
| `ipConfiguration` | ***[[]VirtualNetworkGatewaySpecIpConfiguration](#VirtualNetworkGatewaySpecIpConfiguration)***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `type` | ***string***||
| `vpnClientConfiguration` | ***[[]VirtualNetworkGatewaySpecVpnClientConfiguration](#VirtualNetworkGatewaySpecVpnClientConfiguration)***| ***(Optional)*** |
| `vpnType` | ***string***| ***(Optional)*** |
## VirtualNetworkGatewaySpecBgpSettings
##### (Appears on:[VirtualNetworkGatewaySpec](#VirtualNetworkGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `asn` | ***int***| ***(Optional)*** |
| `peerWeight` | ***int***| ***(Optional)*** |
| `peeringAddress` | ***string***| ***(Optional)*** |
## VirtualNetworkGatewaySpecIpConfiguration
##### (Appears on:[VirtualNetworkGatewaySpec](#VirtualNetworkGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `privateIPAddressAllocation` | ***string***| ***(Optional)*** |
| `publicIPAddressID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***||
## VirtualNetworkGatewaySpecVpnClientConfiguration
##### (Appears on:[VirtualNetworkGatewaySpec](#VirtualNetworkGatewaySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `addressSpace` | ***[]string***||
| `radiusServerAddress` | ***string***| ***(Optional)*** |
| `radiusServerSecret` | ***string***| ***(Optional)*** |
| `revokedCertificate` | ***[[]VirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate](#VirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate)***| ***(Optional)*** |
| `rootCertificate` | ***[[]VirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate](#VirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate)***| ***(Optional)*** |
| `vpnClientProtocols` | ***[]string***| ***(Optional)*** |
## VirtualNetworkGatewaySpecVpnClientConfigurationRevokedCertificate
##### (Appears on:[VirtualNetworkGatewaySpecVpnClientConfiguration](#VirtualNetworkGatewaySpecVpnClientConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `thumbprint` | ***string***||
## VirtualNetworkGatewaySpecVpnClientConfigurationRootCertificate
##### (Appears on:[VirtualNetworkGatewaySpecVpnClientConfiguration](#VirtualNetworkGatewaySpecVpnClientConfiguration))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `publicCertData` | ***string***||
## VirtualNetworkGatewayStatus
##### (Appears on:[VirtualNetworkGateway](#VirtualNetworkGateway))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualNetworkGatewaySpec](#VirtualNetworkGatewaySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
