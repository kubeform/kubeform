---
title: WafSizeConstraintSet
menu:
  docs_{{ .version }}:
    identifier: wafsizeconstraintset-aws.kubeform.com
    name: WafSizeConstraintSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafSizeConstraintSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafSizeConstraintSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafSizeConstraintSetSpec](#WafSizeConstraintSetSpec)***||
| `status` | ***[WafSizeConstraintSetStatus](#WafSizeConstraintSetStatus)***||
## WafSizeConstraintSetSpec

Appears on:[WafSizeConstraintSet](#WafSizeConstraintSet), [WafSizeConstraintSetStatus](#WafSizeConstraintSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `sizeConstraints` | ***[[]WafSizeConstraintSetSpecSizeConstraints](#WafSizeConstraintSetSpecSizeConstraints)***| ***(Optional)*** |
## WafSizeConstraintSetSpecSizeConstraints

Appears on:[WafSizeConstraintSetSpec](#WafSizeConstraintSetSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `comparisonOperator` | ***string***||
| `fieldToMatch` | ***[[]WafSizeConstraintSetSpecSizeConstraintsFieldToMatch](#WafSizeConstraintSetSpecSizeConstraintsFieldToMatch)***||
| `size` | ***int***||
| `textTransformation` | ***string***||
## WafSizeConstraintSetSpecSizeConstraintsFieldToMatch

Appears on:[WafSizeConstraintSetSpecSizeConstraints](#WafSizeConstraintSetSpecSizeConstraints)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafSizeConstraintSetStatus

Appears on:[WafSizeConstraintSet](#WafSizeConstraintSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafSizeConstraintSetSpec](#WafSizeConstraintSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
