---
title: WafSQLInjectionMatchSet
menu:
  docs_v0.0.1:
    identifier: wafsqlinjectionmatchset-aws.kubeform.com
    name: WafSQLInjectionMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## WafSQLInjectionMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafSQLInjectionMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafSQLInjectionMatchSetSpec](#WafSQLInjectionMatchSetSpec)***||
| `status` | ***[WafSQLInjectionMatchSetStatus](#WafSQLInjectionMatchSetStatus)***||
## WafSQLInjectionMatchSetSpec
##### (Appears on:[WafSQLInjectionMatchSet](#WafSQLInjectionMatchSet), [WafSQLInjectionMatchSetStatus](#WafSQLInjectionMatchSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `sqlInjectionMatchTuples` | ***[[]WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples](#WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples)***| ***(Optional)*** |
## WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples
##### (Appears on:[WafSQLInjectionMatchSetSpec](#WafSQLInjectionMatchSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafSQLInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch](#WafSQLInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch)***||
| `textTransformation` | ***string***||
## WafSQLInjectionMatchSetSpecSqlInjectionMatchTuplesFieldToMatch
##### (Appears on:[WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples](#WafSQLInjectionMatchSetSpecSqlInjectionMatchTuples))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafSQLInjectionMatchSetStatus
##### (Appears on:[WafSQLInjectionMatchSet](#WafSQLInjectionMatchSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafSQLInjectionMatchSetSpec](#WafSQLInjectionMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
