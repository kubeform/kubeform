---
title: DbParameterGroup
menu:
  docs_{{ .version }}:
    identifier: dbparametergroup-aws.kubeform.com
    name: DbParameterGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DbParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DbParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DbParameterGroupSpec](#dbparametergroupspec)***||
| `status` | ***[DbParameterGroupStatus](#dbparametergroupstatus)***||
## DbParameterGroupSpec

Appears on:[DbParameterGroup](#dbparametergroup), [DbParameterGroupStatus](#dbparametergroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `parameter` | ***[[]DbParameterGroupSpecParameter](#dbparametergroupspecparameter)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## DbParameterGroupSpecParameter

Appears on:[DbParameterGroupSpec](#dbparametergroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyMethod` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `value` | ***string***||
## DbParameterGroupStatus

Appears on:[DbParameterGroup](#dbparametergroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DbParameterGroupSpec](#dbparametergroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DbParameterGroupStatus](#dbparametergroupstatus)

---
