## VpnConnectionRoute
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpnConnectionRoute` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpnConnectionRouteSpec](#VpnConnectionRouteSpec)***||
| `status` | ***[VpnConnectionRouteStatus](#VpnConnectionRouteStatus)***||
## VpnConnectionRouteSpec
##### (Appears on:[VpnConnectionRoute](#VpnConnectionRoute), [VpnConnectionRouteStatus](#VpnConnectionRouteStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `destinationCIDRBlock` | ***string***||
| `vpnConnectionID` | ***string***||
## VpnConnectionRouteStatus
##### (Appears on:[VpnConnectionRoute](#VpnConnectionRoute))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpnConnectionRouteSpec](#VpnConnectionRouteSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
