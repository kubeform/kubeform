---
title: WafByteMatchSet
menu:
  docs_{{ .version }}:
    identifier: wafbytematchset-aws.kubeform.com
    name: WafByteMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafByteMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafByteMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafByteMatchSetSpec](#wafbytematchsetspec)***||
| `status` | ***[WafByteMatchSetStatus](#wafbytematchsetstatus)***||
## WafByteMatchSetSpec

Appears on:[WafByteMatchSet](#wafbytematchset), [WafByteMatchSetStatus](#wafbytematchsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `byteMatchTuples` | ***[[]WafByteMatchSetSpecByteMatchTuples](#wafbytematchsetspecbytematchtuples)***| ***(Optional)*** |
| `name` | ***string***||
## WafByteMatchSetSpecByteMatchTuples

Appears on:[WafByteMatchSetSpec](#wafbytematchsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafByteMatchSetSpecByteMatchTuplesFieldToMatch](#wafbytematchsetspecbytematchtuplesfieldtomatch)***||
| `positionalConstraint` | ***string***||
| `targetString` | ***string***| ***(Optional)*** |
| `textTransformation` | ***string***||
## WafByteMatchSetSpecByteMatchTuplesFieldToMatch

Appears on:[WafByteMatchSetSpecByteMatchTuples](#wafbytematchsetspecbytematchtuples)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafByteMatchSetStatus

Appears on:[WafByteMatchSet](#wafbytematchset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafByteMatchSetSpec](#wafbytematchsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
