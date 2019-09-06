## ResourcegroupsGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ResourcegroupsGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ResourcegroupsGroupSpec](#ResourcegroupsGroupSpec)***||
| `status` | ***[ResourcegroupsGroupStatus](#ResourcegroupsGroupStatus)***||
## ResourcegroupsGroupSpec
##### (Appears on:[ResourcegroupsGroup](#ResourcegroupsGroup), [ResourcegroupsGroupStatus](#ResourcegroupsGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceQuery` | ***[[]ResourcegroupsGroupSpecResourceQuery](#ResourcegroupsGroupSpecResourceQuery)***||
## ResourcegroupsGroupSpecResourceQuery
##### (Appears on:[ResourcegroupsGroupSpec](#ResourcegroupsGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `query` | ***string***||
| `type` | ***string***| ***(Optional)*** |
## ResourcegroupsGroupStatus
##### (Appears on:[ResourcegroupsGroup](#ResourcegroupsGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ResourcegroupsGroupSpec](#ResourcegroupsGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
