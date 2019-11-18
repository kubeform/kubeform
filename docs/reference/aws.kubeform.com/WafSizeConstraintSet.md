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
| `spec` | ***[WafSizeConstraintSetSpec](#wafsizeconstraintsetspec)***||
| `status` | ***[WafSizeConstraintSetStatus](#wafsizeconstraintsetstatus)***||
## Phase(`string` alias)

Appears on:[WafSizeConstraintSetStatus](#wafsizeconstraintsetstatus)

## WafSizeConstraintSetSpec

Appears on:[WafSizeConstraintSet](#wafsizeconstraintset), [WafSizeConstraintSetStatus](#wafsizeconstraintsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `sizeConstraints` | ***[[]WafSizeConstraintSetSpecSizeConstraints](#wafsizeconstraintsetspecsizeconstraints)***| ***(Optional)*** |
## WafSizeConstraintSetSpecSizeConstraints

Appears on:[WafSizeConstraintSetSpec](#wafsizeconstraintsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `comparisonOperator` | ***string***||
| `fieldToMatch` | ***[[]WafSizeConstraintSetSpecSizeConstraintsFieldToMatch](#wafsizeconstraintsetspecsizeconstraintsfieldtomatch)***||
| `size` | ***int64***||
| `textTransformation` | ***string***||
## WafSizeConstraintSetSpecSizeConstraintsFieldToMatch

Appears on:[WafSizeConstraintSetSpecSizeConstraints](#wafsizeconstraintsetspecsizeconstraints)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafSizeConstraintSetStatus

Appears on:[WafSizeConstraintSet](#wafsizeconstraintset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafSizeConstraintSetSpec](#wafsizeconstraintsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
