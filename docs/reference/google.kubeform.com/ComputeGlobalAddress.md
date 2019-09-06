## ComputeGlobalAddress
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeGlobalAddress` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeGlobalAddressSpec](#ComputeGlobalAddressSpec)***||
| `status` | ***[ComputeGlobalAddressStatus](#ComputeGlobalAddressStatus)***||
## ComputeGlobalAddressSpec
##### (Appears on:[ComputeGlobalAddress](#ComputeGlobalAddress), [ComputeGlobalAddressStatus](#ComputeGlobalAddressStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `address` | ***string***| ***(Optional)*** |
| `addressType` | ***string***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `ipVersion` | ***string***| ***(Optional)*** |
| `labelFingerprint` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** Deprecated|
| `name` | ***string***||
| `network` | ***string***| ***(Optional)*** Deprecated|
| `prefixLength` | ***int***| ***(Optional)*** Deprecated|
| `project` | ***string***| ***(Optional)*** |
| `purpose` | ***string***| ***(Optional)*** Deprecated|
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeGlobalAddressStatus
##### (Appears on:[ComputeGlobalAddress](#ComputeGlobalAddress))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeGlobalAddressSpec](#ComputeGlobalAddressSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
