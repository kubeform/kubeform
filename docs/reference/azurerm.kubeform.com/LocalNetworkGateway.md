---
title: LocalNetworkGateway
menu:
  docs_{{ .version }}:
    identifier: localnetworkgateway-azurerm.kubeform.com
    name: LocalNetworkGateway
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LocalNetworkGateway
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LocalNetworkGateway` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LocalNetworkGatewaySpec](#localnetworkgatewayspec)***||
| `status` | ***[LocalNetworkGatewayStatus](#localnetworkgatewaystatus)***||
## LocalNetworkGatewaySpec

Appears on:[LocalNetworkGateway](#localnetworkgateway), [LocalNetworkGatewayStatus](#localnetworkgatewaystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addressSpace` | ***[]string***||
| `bgpSettings` | ***[[]LocalNetworkGatewaySpecBgpSettings](#localnetworkgatewayspecbgpsettings)***| ***(Optional)*** |
| `gatewayAddress` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## LocalNetworkGatewaySpecBgpSettings

Appears on:[LocalNetworkGatewaySpec](#localnetworkgatewayspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `asn` | ***int64***||
| `bgpPeeringAddress` | ***string***||
| `peerWeight` | ***int64***| ***(Optional)*** |
## LocalNetworkGatewayStatus

Appears on:[LocalNetworkGateway](#localnetworkgateway)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LocalNetworkGatewaySpec](#localnetworkgatewayspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LocalNetworkGatewayStatus](#localnetworkgatewaystatus)

---
