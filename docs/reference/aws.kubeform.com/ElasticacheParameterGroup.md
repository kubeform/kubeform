## ElasticacheParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticacheParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticacheParameterGroupSpec](#ElasticacheParameterGroupSpec)***||
| `status` | ***[ElasticacheParameterGroupStatus](#ElasticacheParameterGroupStatus)***||
## ElasticacheParameterGroupSpec
##### (Appears on:[ElasticacheParameterGroup](#ElasticacheParameterGroup), [ElasticacheParameterGroupStatus](#ElasticacheParameterGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***||
| `parameter` | ***[[]ElasticacheParameterGroupSpecParameter](#ElasticacheParameterGroupSpecParameter)***| ***(Optional)*** |
## ElasticacheParameterGroupSpecParameter
##### (Appears on:[ElasticacheParameterGroupSpec](#ElasticacheParameterGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## ElasticacheParameterGroupStatus
##### (Appears on:[ElasticacheParameterGroup](#ElasticacheParameterGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticacheParameterGroupSpec](#ElasticacheParameterGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
