---
title: DaxParameterGroup
menu:
  docs_{{ .version }}:
    identifier: daxparametergroup-aws.kubeform.com
    name: DaxParameterGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DaxParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DaxParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DaxParameterGroupSpec](#daxparametergroupspec)***||
| `status` | ***[DaxParameterGroupStatus](#daxparametergroupstatus)***||
## DaxParameterGroupSpec

Appears on:[DaxParameterGroup](#daxparametergroup), [DaxParameterGroupStatus](#daxparametergroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parameters` | ***[[]DaxParameterGroupSpecParameters](#daxparametergroupspecparameters)***| ***(Optional)*** |
## DaxParameterGroupSpecParameters

Appears on:[DaxParameterGroupSpec](#daxparametergroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## DaxParameterGroupStatus

Appears on:[DaxParameterGroup](#daxparametergroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DaxParameterGroupSpec](#daxparametergroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DaxParameterGroupStatus](#daxparametergroupstatus)

---
