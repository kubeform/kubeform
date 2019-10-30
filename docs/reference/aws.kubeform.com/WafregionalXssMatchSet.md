---
title: WafregionalXssMatchSet
menu:
  docs_{{ .version }}:
    identifier: wafregionalxssmatchset-aws.kubeform.com
    name: WafregionalXssMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalXssMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalXssMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalXssMatchSetSpec](#wafregionalxssmatchsetspec)***||
| `status` | ***[WafregionalXssMatchSetStatus](#wafregionalxssmatchsetstatus)***||
## Phase(`string` alias)

Appears on:[WafregionalXssMatchSetStatus](#wafregionalxssmatchsetstatus)

## WafregionalXssMatchSetSpec

Appears on:[WafregionalXssMatchSet](#wafregionalxssmatchset), [WafregionalXssMatchSetStatus](#wafregionalxssmatchsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `xssMatchTuple` | ***[[]WafregionalXssMatchSetSpecXssMatchTuple](#wafregionalxssmatchsetspecxssmatchtuple)***| ***(Optional)*** |
## WafregionalXssMatchSetSpecXssMatchTuple

Appears on:[WafregionalXssMatchSetSpec](#wafregionalxssmatchsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafregionalXssMatchSetSpecXssMatchTupleFieldToMatch](#wafregionalxssmatchsetspecxssmatchtuplefieldtomatch)***||
| `textTransformation` | ***string***||
## WafregionalXssMatchSetSpecXssMatchTupleFieldToMatch

Appears on:[WafregionalXssMatchSetSpecXssMatchTuple](#wafregionalxssmatchsetspecxssmatchtuple)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafregionalXssMatchSetStatus

Appears on:[WafregionalXssMatchSet](#wafregionalxssmatchset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalXssMatchSetSpec](#wafregionalxssmatchsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
