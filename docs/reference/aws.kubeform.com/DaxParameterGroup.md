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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DaxParameterGroupSpec](#DaxParameterGroupSpec)***||
| `status` | ***[DaxParameterGroupStatus](#DaxParameterGroupStatus)***||
## DaxParameterGroupSpec

Appears on:[DaxParameterGroup](#DaxParameterGroup), [DaxParameterGroupStatus](#DaxParameterGroupStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `parameters` | ***[[]DaxParameterGroupSpecParameters](#DaxParameterGroupSpecParameters)***| ***(Optional)*** |
## DaxParameterGroupSpecParameters

Appears on:[DaxParameterGroupSpec](#DaxParameterGroupSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## DaxParameterGroupStatus

Appears on:[DaxParameterGroup](#DaxParameterGroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DaxParameterGroupSpec](#DaxParameterGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
