## WafregionalRegexPatternSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalRegexPatternSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalRegexPatternSetSpec](#WafregionalRegexPatternSetSpec)***||
| `status` | ***[WafregionalRegexPatternSetStatus](#WafregionalRegexPatternSetStatus)***||
## WafregionalRegexPatternSetSpec
##### (Appears on:[WafregionalRegexPatternSet](#WafregionalRegexPatternSet), [WafregionalRegexPatternSetStatus](#WafregionalRegexPatternSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `regexPatternStrings` | ***[]string***| ***(Optional)*** |
## WafregionalRegexPatternSetStatus
##### (Appears on:[WafregionalRegexPatternSet](#WafregionalRegexPatternSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalRegexPatternSetSpec](#WafregionalRegexPatternSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
