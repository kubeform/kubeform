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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafRegexMatchSetSpec](#wafregexmatchsetspec)***||
| `status` | ***[WafRegexMatchSetStatus](#wafregexmatchsetstatus)***||
## Phase(`string` alias)

Appears on:[WafRegexMatchSetStatus](#wafregexmatchsetstatus)

## WafRegexMatchSetSpec

Appears on:[WafRegexMatchSet](#wafregexmatchset), [WafRegexMatchSetStatus](#wafregexmatchsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `regexMatchTuple` | ***[[]WafRegexMatchSetSpecRegexMatchTuple](#wafregexmatchsetspecregexmatchtuple)***| ***(Optional)*** |
## WafRegexMatchSetSpecRegexMatchTuple

Appears on:[WafRegexMatchSetSpec](#wafregexmatchsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafRegexMatchSetSpecRegexMatchTupleFieldToMatch](#wafregexmatchsetspecregexmatchtuplefieldtomatch)***||
| `regexPatternSetID` | ***string***||
| `textTransformation` | ***string***||
## WafRegexMatchSetSpecRegexMatchTupleFieldToMatch

Appears on:[WafRegexMatchSetSpecRegexMatchTuple](#wafregexmatchsetspecregexmatchtuple)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafRegexMatchSetStatus

Appears on:[WafRegexMatchSet](#wafregexmatchset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafRegexMatchSetSpec](#wafregexmatchsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
