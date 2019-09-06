## ComputeRoute
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRoute` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRouteSpec](#ComputeRouteSpec)***||
| `status` | ***[ComputeRouteStatus](#ComputeRouteStatus)***||
## ComputeRouteSpec
##### (Appears on:[ComputeRoute](#ComputeRoute), [ComputeRouteStatus](#ComputeRouteStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `destRange` | ***string***||
| `name` | ***string***||
| `network` | ***string***||
| `nextHopGateway` | ***string***| ***(Optional)*** |
| `nextHopInstance` | ***string***| ***(Optional)*** |
| `nextHopInstanceZone` | ***string***| ***(Optional)*** |
| `nextHopIP` | ***string***| ***(Optional)*** |
| `nextHopNetwork` | ***string***| ***(Optional)*** |
| `nextHopVPNTunnel` | ***string***| ***(Optional)*** |
| `priority` | ***int***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
## ComputeRouteStatus
##### (Appears on:[ComputeRoute](#ComputeRoute))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRouteSpec](#ComputeRouteSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
