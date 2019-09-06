## SesDomainDkim
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesDomainDkim` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesDomainDkimSpec](#SesDomainDkimSpec)***||
| `status` | ***[SesDomainDkimStatus](#SesDomainDkimStatus)***||
## SesDomainDkimSpec
##### (Appears on:[SesDomainDkim](#SesDomainDkim), [SesDomainDkimStatus](#SesDomainDkimStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dkimTokens` | ***[]string***| ***(Optional)*** |
| `domain` | ***string***||
## SesDomainDkimStatus
##### (Appears on:[SesDomainDkim](#SesDomainDkim))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesDomainDkimSpec](#SesDomainDkimSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
