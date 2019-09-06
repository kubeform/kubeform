## ComputeRouterNAT
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRouterNAT` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRouterNATSpec](#ComputeRouterNATSpec)***||
| `status` | ***[ComputeRouterNATStatus](#ComputeRouterNATStatus)***||
## ComputeRouterNATSpec
##### (Appears on:[ComputeRouterNAT](#ComputeRouterNAT), [ComputeRouterNATStatus](#ComputeRouterNATStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `icmpIdleTimeoutSec` | ***int***| ***(Optional)*** |
| `minPortsPerVm` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `natIPAllocateOption` | ***string***||
| `natIPS` | ***[]string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `router` | ***string***||
| `sourceSubnetworkIPRangesToNAT` | ***string***| ***(Optional)*** |
| `subnetwork` | ***[[]ComputeRouterNATSpecSubnetwork](#ComputeRouterNATSpecSubnetwork)***| ***(Optional)*** |
| `tcpEstablishedIdleTimeoutSec` | ***int***| ***(Optional)*** |
| `tcpTransitoryIdleTimeoutSec` | ***int***| ***(Optional)*** |
| `udpIdleTimeoutSec` | ***int***| ***(Optional)*** |
## ComputeRouterNATSpecSubnetwork
##### (Appears on:[ComputeRouterNATSpec](#ComputeRouterNATSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `secondaryIPRangeNames` | ***[]string***| ***(Optional)*** |
| `sourceIPRangesToNAT` | ***[]string***| ***(Optional)*** |
## ComputeRouterNATStatus
##### (Appears on:[ComputeRouterNAT](#ComputeRouterNAT))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRouterNATSpec](#ComputeRouterNATSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
