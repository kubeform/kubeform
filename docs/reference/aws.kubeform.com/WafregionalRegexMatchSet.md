---
title: WafregionalRegexMatchSet
menu:
  docs_{{ .version }}:
    identifier: wafregionalregexmatchset-aws.kubeform.com
    name: WafregionalRegexMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalRegexMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalRegexMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalRegexMatchSetSpec](#WafregionalRegexMatchSetSpec)***||
| `status` | ***[WafregionalRegexMatchSetStatus](#WafregionalRegexMatchSetStatus)***||
## WafregionalRegexMatchSetSpec
##### (Appears on:[WafregionalRegexMatchSet](#WafregionalRegexMatchSet), [WafregionalRegexMatchSetStatus](#WafregionalRegexMatchSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `regexMatchTuple` | ***[[]WafregionalRegexMatchSetSpecRegexMatchTuple](#WafregionalRegexMatchSetSpecRegexMatchTuple)***| ***(Optional)*** |
## WafregionalRegexMatchSetSpecRegexMatchTuple
##### (Appears on:[WafregionalRegexMatchSetSpec](#WafregionalRegexMatchSetSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafregionalRegexMatchSetSpecRegexMatchTupleFieldToMatch](#WafregionalRegexMatchSetSpecRegexMatchTupleFieldToMatch)***||
| `regexPatternSetID` | ***string***||
| `textTransformation` | ***string***||
## WafregionalRegexMatchSetSpecRegexMatchTupleFieldToMatch
##### (Appears on:[WafregionalRegexMatchSetSpecRegexMatchTuple](#WafregionalRegexMatchSetSpecRegexMatchTuple))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafregionalRegexMatchSetStatus
##### (Appears on:[WafregionalRegexMatchSet](#WafregionalRegexMatchSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalRegexMatchSetSpec](#WafregionalRegexMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
