---
title: NeptuneParameterGroup
menu:
  docs_v0.0.1:
    identifier: neptuneparametergroup-aws.kubeform.com
    name: NeptuneParameterGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## NeptuneParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NeptuneParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NeptuneParameterGroupSpec](#NeptuneParameterGroupSpec)***||
| `status` | ***[NeptuneParameterGroupStatus](#NeptuneParameterGroupStatus)***||
## NeptuneParameterGroupSpec
##### (Appears on:[NeptuneParameterGroup](#NeptuneParameterGroup), [NeptuneParameterGroupStatus](#NeptuneParameterGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***||
| `parameter` | ***[[]NeptuneParameterGroupSpecParameter](#NeptuneParameterGroupSpecParameter)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## NeptuneParameterGroupSpecParameter
##### (Appears on:[NeptuneParameterGroupSpec](#NeptuneParameterGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyMethod` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `value` | ***string***||
## NeptuneParameterGroupStatus
##### (Appears on:[NeptuneParameterGroup](#NeptuneParameterGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NeptuneParameterGroupSpec](#NeptuneParameterGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
