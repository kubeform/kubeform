---
title: WafregionalByteMatchSet
menu:
  docs_{{ .version }}:
    identifier: wafregionalbytematchset-aws.kubeform.com
    name: WafregionalByteMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalByteMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalByteMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalByteMatchSetSpec](#wafregionalbytematchsetspec)***||
| `status` | ***[WafregionalByteMatchSetStatus](#wafregionalbytematchsetstatus)***||
## Phase(`string` alias)

Appears on:[WafregionalByteMatchSetStatus](#wafregionalbytematchsetstatus)

## WafregionalByteMatchSetSpec

Appears on:[WafregionalByteMatchSet](#wafregionalbytematchset), [WafregionalByteMatchSetStatus](#wafregionalbytematchsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `byteMatchTuples` | ***[[]WafregionalByteMatchSetSpecByteMatchTuples](#wafregionalbytematchsetspecbytematchtuples)***| ***(Optional)*** |
| `name` | ***string***||
## WafregionalByteMatchSetSpecByteMatchTuples

Appears on:[WafregionalByteMatchSetSpec](#wafregionalbytematchsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch](#wafregionalbytematchsetspecbytematchtuplesfieldtomatch)***||
| `positionalConstraint` | ***string***||
| `targetString` | ***string***| ***(Optional)*** |
| `textTransformation` | ***string***||
## WafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch

Appears on:[WafregionalByteMatchSetSpecByteMatchTuples](#wafregionalbytematchsetspecbytematchtuples)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafregionalByteMatchSetStatus

Appears on:[WafregionalByteMatchSet](#wafregionalbytematchset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalByteMatchSetSpec](#wafregionalbytematchsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
