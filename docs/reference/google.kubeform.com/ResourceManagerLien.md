## ResourceManagerLien
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ResourceManagerLien` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ResourceManagerLienSpec](#ResourceManagerLienSpec)***||
| `status` | ***[ResourceManagerLienStatus](#ResourceManagerLienStatus)***||
## ResourceManagerLienSpec
##### (Appears on:[ResourceManagerLien](#ResourceManagerLien), [ResourceManagerLienStatus](#ResourceManagerLienStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createTime` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `origin` | ***string***||
| `parent` | ***string***||
| `reason` | ***string***||
| `restrictions` | ***[]string***||
## ResourceManagerLienStatus
##### (Appears on:[ResourceManagerLien](#ResourceManagerLien))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ResourceManagerLienSpec](#ResourceManagerLienSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
