---
title: WafRegexMatchSet
menu:
  docs_{{ .version }}:
    identifier: wafregexmatchset-aws.kubeform.com
    name: WafRegexMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafRegexMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafRegexMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafRegexMatchSetSpec](#WafRegexMatchSetSpec)***||
| `status` | ***[WafRegexMatchSetStatus](#WafRegexMatchSetStatus)***||
## WafRegexMatchSetSpec

Appears on:[WafRegexMatchSet](#WafRegexMatchSet), [WafRegexMatchSetStatus](#WafRegexMatchSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `regexMatchTuple` | ***[[]WafRegexMatchSetSpecRegexMatchTuple](#WafRegexMatchSetSpecRegexMatchTuple)***| ***(Optional)*** |
## WafRegexMatchSetSpecRegexMatchTuple

Appears on:[WafRegexMatchSetSpec](#WafRegexMatchSetSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafRegexMatchSetSpecRegexMatchTupleFieldToMatch](#WafRegexMatchSetSpecRegexMatchTupleFieldToMatch)***||
| `regexPatternSetID` | ***string***||
| `textTransformation` | ***string***||
## WafRegexMatchSetSpecRegexMatchTupleFieldToMatch

Appears on:[WafRegexMatchSetSpecRegexMatchTuple](#WafRegexMatchSetSpecRegexMatchTuple)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafRegexMatchSetStatus

Appears on:[WafRegexMatchSet](#WafRegexMatchSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafRegexMatchSetSpec](#WafRegexMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
