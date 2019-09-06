## WafByteMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafByteMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafByteMatchSetSpec](#WafByteMatchSetSpec)***||
| `status` | ***[WafByteMatchSetStatus](#WafByteMatchSetStatus)***||
## WafByteMatchSetSpec
##### (Appears on:[WafByteMatchSet](#WafByteMatchSet), [WafByteMatchSetStatus](#WafByteMatchSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `byteMatchTuples` | ***[[]WafByteMatchSetSpecByteMatchTuples](#WafByteMatchSetSpecByteMatchTuples)***| ***(Optional)*** |
| `name` | ***string***||
## WafByteMatchSetSpecByteMatchTuples
##### (Appears on:[WafByteMatchSetSpec](#WafByteMatchSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafByteMatchSetSpecByteMatchTuplesFieldToMatch](#WafByteMatchSetSpecByteMatchTuplesFieldToMatch)***||
| `positionalConstraint` | ***string***||
| `targetString` | ***string***| ***(Optional)*** |
| `textTransformation` | ***string***||
## WafByteMatchSetSpecByteMatchTuplesFieldToMatch
##### (Appears on:[WafByteMatchSetSpecByteMatchTuples](#WafByteMatchSetSpecByteMatchTuples))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafByteMatchSetStatus
##### (Appears on:[WafByteMatchSet](#WafByteMatchSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafByteMatchSetSpec](#WafByteMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
