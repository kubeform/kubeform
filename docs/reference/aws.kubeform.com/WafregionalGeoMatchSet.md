## WafregionalGeoMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalGeoMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalGeoMatchSetSpec](#WafregionalGeoMatchSetSpec)***||
| `status` | ***[WafregionalGeoMatchSetStatus](#WafregionalGeoMatchSetStatus)***||
## WafregionalGeoMatchSetSpec
##### (Appears on:[WafregionalGeoMatchSet](#WafregionalGeoMatchSet), [WafregionalGeoMatchSetStatus](#WafregionalGeoMatchSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `geoMatchConstraint` | ***[[]WafregionalGeoMatchSetSpecGeoMatchConstraint](#WafregionalGeoMatchSetSpecGeoMatchConstraint)***| ***(Optional)*** |
| `name` | ***string***||
## WafregionalGeoMatchSetSpecGeoMatchConstraint
##### (Appears on:[WafregionalGeoMatchSetSpec](#WafregionalGeoMatchSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
| `value` | ***string***||
## WafregionalGeoMatchSetStatus
##### (Appears on:[WafregionalGeoMatchSet](#WafregionalGeoMatchSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalGeoMatchSetSpec](#WafregionalGeoMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
