## DocdbClusterParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DocdbClusterParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DocdbClusterParameterGroupSpec](#DocdbClusterParameterGroupSpec)***||
| `status` | ***[DocdbClusterParameterGroupStatus](#DocdbClusterParameterGroupStatus)***||
## DocdbClusterParameterGroupSpec
##### (Appears on:[DocdbClusterParameterGroup](#DocdbClusterParameterGroup), [DocdbClusterParameterGroupStatus](#DocdbClusterParameterGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `parameter` | ***[[]DocdbClusterParameterGroupSpecParameter](#DocdbClusterParameterGroupSpecParameter)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## DocdbClusterParameterGroupSpecParameter
##### (Appears on:[DocdbClusterParameterGroupSpec](#DocdbClusterParameterGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyMethod` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `value` | ***string***||
## DocdbClusterParameterGroupStatus
##### (Appears on:[DocdbClusterParameterGroup](#DocdbClusterParameterGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DocdbClusterParameterGroupSpec](#DocdbClusterParameterGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
