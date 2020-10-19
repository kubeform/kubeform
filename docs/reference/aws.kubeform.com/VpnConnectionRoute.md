---
title: VpnConnectionRoute
menu:
  docs_{{ .version }}:
    identifier: vpnconnectionroute-aws.kubeform.com
    name: VpnConnectionRoute
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpnConnectionRoute
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpnConnectionRoute` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpnConnectionRouteSpec](#vpnconnectionroutespec)***||
| `status` | ***[VpnConnectionRouteStatus](#vpnconnectionroutestatus)***||
## Phase(`string` alias)

Appears on:[VpnConnectionRouteStatus](#vpnconnectionroutestatus)

## VpnConnectionRouteSpec

Appears on:[VpnConnectionRoute](#vpnconnectionroute), [VpnConnectionRouteStatus](#vpnconnectionroutestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `destinationCIDRBlock` | ***string***||
| `vpnConnectionID` | ***string***||
## VpnConnectionRouteStatus

Appears on:[VpnConnectionRoute](#vpnconnectionroute)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpnConnectionRouteSpec](#vpnconnectionroutespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
