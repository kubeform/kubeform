---
title: WafXssMatchSet
menu:
  docs_v0.0.1:
    identifier: wafxssmatchset-aws.kubeform.com
    name: WafXssMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## WafXssMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafXssMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafXssMatchSetSpec](#WafXssMatchSetSpec)***||
| `status` | ***[WafXssMatchSetStatus](#WafXssMatchSetStatus)***||
## WafXssMatchSetSpec
##### (Appears on:[WafXssMatchSet](#WafXssMatchSet), [WafXssMatchSetStatus](#WafXssMatchSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `xssMatchTuples` | ***[[]WafXssMatchSetSpecXssMatchTuples](#WafXssMatchSetSpecXssMatchTuples)***| ***(Optional)*** |
## WafXssMatchSetSpecXssMatchTuples
##### (Appears on:[WafXssMatchSetSpec](#WafXssMatchSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafXssMatchSetSpecXssMatchTuplesFieldToMatch](#WafXssMatchSetSpecXssMatchTuplesFieldToMatch)***||
| `textTransformation` | ***string***||
## WafXssMatchSetSpecXssMatchTuplesFieldToMatch
##### (Appears on:[WafXssMatchSetSpecXssMatchTuples](#WafXssMatchSetSpecXssMatchTuples))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafXssMatchSetStatus
##### (Appears on:[WafXssMatchSet](#WafXssMatchSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafXssMatchSetSpec](#WafXssMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
