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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalByteMatchSetSpec](#WafregionalByteMatchSetSpec)***||
| `status` | ***[WafregionalByteMatchSetStatus](#WafregionalByteMatchSetStatus)***||
## WafregionalByteMatchSetSpec

Appears on:[WafregionalByteMatchSet](#WafregionalByteMatchSet), [WafregionalByteMatchSetStatus](#WafregionalByteMatchSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `byteMatchTuples` | ***[[]WafregionalByteMatchSetSpecByteMatchTuples](#WafregionalByteMatchSetSpecByteMatchTuples)***| ***(Optional)*** |
| `name` | ***string***||
## WafregionalByteMatchSetSpecByteMatchTuples

Appears on:[WafregionalByteMatchSetSpec](#WafregionalByteMatchSetSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `fieldToMatch` | ***[[]WafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch](#WafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch)***||
| `positionalConstraint` | ***string***||
| `targetString` | ***string***| ***(Optional)*** |
| `textTransformation` | ***string***||
## WafregionalByteMatchSetSpecByteMatchTuplesFieldToMatch

Appears on:[WafregionalByteMatchSetSpecByteMatchTuples](#WafregionalByteMatchSetSpecByteMatchTuples)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafregionalByteMatchSetStatus

Appears on:[WafregionalByteMatchSet](#WafregionalByteMatchSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalByteMatchSetSpec](#WafregionalByteMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
