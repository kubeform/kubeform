---
title: WafregionalSizeConstraintSet
menu:
  docs_{{ .version }}:
    identifier: wafregionalsizeconstraintset-aws.kubeform.com
    name: WafregionalSizeConstraintSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalSizeConstraintSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalSizeConstraintSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalSizeConstraintSetSpec](#wafregionalsizeconstraintsetspec)***||
| `status` | ***[WafregionalSizeConstraintSetStatus](#wafregionalsizeconstraintsetstatus)***||
## Phase(`string` alias)

Appears on:[WafregionalSizeConstraintSetStatus](#wafregionalsizeconstraintsetstatus)

## WafregionalSizeConstraintSetSpec

Appears on:[WafregionalSizeConstraintSet](#wafregionalsizeconstraintset), [WafregionalSizeConstraintSetStatus](#wafregionalsizeconstraintsetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `sizeConstraints` | ***[[]WafregionalSizeConstraintSetSpecSizeConstraints](#wafregionalsizeconstraintsetspecsizeconstraints)***| ***(Optional)*** |
## WafregionalSizeConstraintSetSpecSizeConstraints

Appears on:[WafregionalSizeConstraintSetSpec](#wafregionalsizeconstraintsetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `comparisonOperator` | ***string***||
| `fieldToMatch` | ***[[]WafregionalSizeConstraintSetSpecSizeConstraintsFieldToMatch](#wafregionalsizeconstraintsetspecsizeconstraintsfieldtomatch)***||
| `size` | ***int***||
| `textTransformation` | ***string***||
## WafregionalSizeConstraintSetSpecSizeConstraintsFieldToMatch

Appears on:[WafregionalSizeConstraintSetSpecSizeConstraints](#wafregionalsizeconstraintsetspecsizeconstraints)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `data` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## WafregionalSizeConstraintSetStatus

Appears on:[WafregionalSizeConstraintSet](#wafregionalsizeconstraintset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalSizeConstraintSetSpec](#wafregionalsizeconstraintsetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
