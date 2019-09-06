## SfnActivity
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SfnActivity` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SfnActivitySpec](#SfnActivitySpec)***||
| `status` | ***[SfnActivityStatus](#SfnActivityStatus)***||
## SfnActivitySpec
##### (Appears on:[SfnActivity](#SfnActivity), [SfnActivityStatus](#SfnActivityStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `creationDate` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## SfnActivityStatus
##### (Appears on:[SfnActivity](#SfnActivity))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SfnActivitySpec](#SfnActivitySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
