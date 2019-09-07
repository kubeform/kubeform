---
title: NeptuneClusterParameterGroup
menu:
  docs_v0.0.1:
    identifier: neptuneclusterparametergroup-aws.kubeform.com
    name: NeptuneClusterParameterGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## NeptuneClusterParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NeptuneClusterParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NeptuneClusterParameterGroupSpec](#NeptuneClusterParameterGroupSpec)***||
| `status` | ***[NeptuneClusterParameterGroupStatus](#NeptuneClusterParameterGroupStatus)***||
## NeptuneClusterParameterGroupSpec
##### (Appears on:[NeptuneClusterParameterGroup](#NeptuneClusterParameterGroup), [NeptuneClusterParameterGroupStatus](#NeptuneClusterParameterGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `parameter` | ***[[]NeptuneClusterParameterGroupSpecParameter](#NeptuneClusterParameterGroupSpecParameter)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## NeptuneClusterParameterGroupSpecParameter
##### (Appears on:[NeptuneClusterParameterGroupSpec](#NeptuneClusterParameterGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyMethod` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `value` | ***string***||
## NeptuneClusterParameterGroupStatus
##### (Appears on:[NeptuneClusterParameterGroup](#NeptuneClusterParameterGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NeptuneClusterParameterGroupSpec](#NeptuneClusterParameterGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
