---
title: NeptuneClusterParameterGroup
menu:
  docs_{{ .version }}:
    identifier: neptuneclusterparametergroup-aws.kubeform.com
    name: NeptuneClusterParameterGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NeptuneClusterParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NeptuneClusterParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NeptuneClusterParameterGroupSpec](#neptuneclusterparametergroupspec)***||
| `status` | ***[NeptuneClusterParameterGroupStatus](#neptuneclusterparametergroupstatus)***||
## NeptuneClusterParameterGroupSpec

Appears on:[NeptuneClusterParameterGroup](#neptuneclusterparametergroup), [NeptuneClusterParameterGroupStatus](#neptuneclusterparametergroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `parameter` | ***[[]NeptuneClusterParameterGroupSpecParameter](#neptuneclusterparametergroupspecparameter)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## NeptuneClusterParameterGroupSpecParameter

Appears on:[NeptuneClusterParameterGroupSpec](#neptuneclusterparametergroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyMethod` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `value` | ***string***||
## NeptuneClusterParameterGroupStatus

Appears on:[NeptuneClusterParameterGroup](#neptuneclusterparametergroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NeptuneClusterParameterGroupSpec](#neptuneclusterparametergroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
