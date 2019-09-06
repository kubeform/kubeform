## RedshiftParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RedshiftParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RedshiftParameterGroupSpec](#RedshiftParameterGroupSpec)***||
| `status` | ***[RedshiftParameterGroupStatus](#RedshiftParameterGroupStatus)***||
## RedshiftParameterGroupSpec
##### (Appears on:[RedshiftParameterGroup](#RedshiftParameterGroup), [RedshiftParameterGroupStatus](#RedshiftParameterGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***||
| `parameter` | ***[[]RedshiftParameterGroupSpecParameter](#RedshiftParameterGroupSpecParameter)***| ***(Optional)*** |
## RedshiftParameterGroupSpecParameter
##### (Appears on:[RedshiftParameterGroupSpec](#RedshiftParameterGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## RedshiftParameterGroupStatus
##### (Appears on:[RedshiftParameterGroup](#RedshiftParameterGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RedshiftParameterGroupSpec](#RedshiftParameterGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
