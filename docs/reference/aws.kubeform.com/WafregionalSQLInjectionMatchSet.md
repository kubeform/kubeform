## WafregionalSQLInjectionMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalSQLInjectionMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalSQLInjectionMatchSetSpec](#WafregionalSQLInjectionMatchSetSpec)***||
| `status` | ***[WafregionalSQLInjectionMatchSetStatus](#WafregionalSQLInjectionMatchSetStatus)***||
## WafregionalSQLInjectionMatchSetSpec
##### (Appears on:[WafregionalSQLInjectionMatchSet](#WafregionalSQLInjectionMatchSet), [WafregionalSQLInjectionMatchSetStatus](#WafregionalSQLInjectionMatchSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `sqlInjectionMatchTuple` | ***[[]WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTuple](#WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTuple)***| ***(Optional)*** |
## WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTuple
##### (Appears on:[WafregionalSQLInjectionMatchSetSpec](#WafregionalSQLInjectionMatchSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch](#WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch)***||
| `textTransformation` | ***string***||
## WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTupleFieldToMatch
##### (Appears on:[WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTuple](#WafregionalSQLInjectionMatchSetSpecSqlInjectionMatchTuple))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafregionalSQLInjectionMatchSetStatus
##### (Appears on:[WafregionalSQLInjectionMatchSet](#WafregionalSQLInjectionMatchSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalSQLInjectionMatchSetSpec](#WafregionalSQLInjectionMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
