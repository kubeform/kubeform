## ComputeSharedVpcHostProject
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeSharedVpcHostProject` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeSharedVpcHostProjectSpec](#ComputeSharedVpcHostProjectSpec)***||
| `status` | ***[ComputeSharedVpcHostProjectStatus](#ComputeSharedVpcHostProjectStatus)***||
## ComputeSharedVpcHostProjectSpec
##### (Appears on:[ComputeSharedVpcHostProject](#ComputeSharedVpcHostProject), [ComputeSharedVpcHostProjectStatus](#ComputeSharedVpcHostProjectStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `project` | ***string***||
## ComputeSharedVpcHostProjectStatus
##### (Appears on:[ComputeSharedVpcHostProject](#ComputeSharedVpcHostProject))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeSharedVpcHostProjectSpec](#ComputeSharedVpcHostProjectSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
