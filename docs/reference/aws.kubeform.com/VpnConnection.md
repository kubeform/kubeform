---
title: VpnConnection
menu:
  docs_{{ .version }}:
    identifier: vpnconnection-aws.kubeform.com
    name: VpnConnection
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpnConnection
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpnConnection` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpnConnectionSpec](#vpnconnectionspec)***||
| `status` | ***[VpnConnectionStatus](#vpnconnectionstatus)***||
## Phase(`string` alias)

Appears on:[VpnConnectionStatus](#vpnconnectionstatus)

## VpnConnectionSpec

Appears on:[VpnConnection](#vpnconnection), [VpnConnectionStatus](#vpnconnectionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `customerGatewayConfiguration` | ***string***| ***(Optional)*** |
| `customerGatewayID` | ***string***||
| `routes` | ***[[]VpnConnectionSpecRoutes](#vpnconnectionspecroutes)***| ***(Optional)*** |
| `staticRoutesOnly` | ***bool***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `transitGatewayAttachmentID` | ***string***| ***(Optional)*** |
| `transitGatewayID` | ***string***| ***(Optional)*** |
| `tunnel1Address` | ***string***| ***(Optional)*** |
| `tunnel1BGPAsn` | ***string***| ***(Optional)*** |
| `tunnel1BGPHoldtime` | ***int64***| ***(Optional)*** |
| `tunnel1CgwInsideAddress` | ***string***| ***(Optional)*** |
| `tunnel1InsideCIDR` | ***string***| ***(Optional)*** |
| `tunnel1VgwInsideAddress` | ***string***| ***(Optional)*** |
| `tunnel2Address` | ***string***| ***(Optional)*** |
| `tunnel2BGPAsn` | ***string***| ***(Optional)*** |
| `tunnel2BGPHoldtime` | ***int64***| ***(Optional)*** |
| `tunnel2CgwInsideAddress` | ***string***| ***(Optional)*** |
| `tunnel2InsideCIDR` | ***string***| ***(Optional)*** |
| `tunnel2VgwInsideAddress` | ***string***| ***(Optional)*** |
| `type` | ***string***||
| `vgwTelemetry` | ***[[]VpnConnectionSpecVgwTelemetry](#vpnconnectionspecvgwtelemetry)***| ***(Optional)*** |
| `vpnGatewayID` | ***string***| ***(Optional)*** |
## VpnConnectionSpecRoutes

Appears on:[VpnConnectionSpec](#vpnconnectionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `destinationCIDRBlock` | ***string***| ***(Optional)*** |
| `source` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
## VpnConnectionSpecVgwTelemetry

Appears on:[VpnConnectionSpec](#vpnconnectionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `acceptedRouteCount` | ***int64***| ***(Optional)*** |
| `lastStatusChange` | ***string***| ***(Optional)*** |
| `outsideIPAddress` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `statusMessage` | ***string***| ***(Optional)*** |
## VpnConnectionStatus

Appears on:[VpnConnection](#vpnconnection)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpnConnectionSpec](#vpnconnectionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `tunnel1_preshared_key` | ***string*** ||
| `tunnel2_preshared_key` | ***string*** ||
