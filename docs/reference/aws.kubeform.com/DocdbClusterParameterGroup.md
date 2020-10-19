---
title: DocdbClusterParameterGroup
menu:
  docs_{{ .version }}:
    identifier: docdbclusterparametergroup-aws.kubeform.com
    name: DocdbClusterParameterGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DocdbClusterParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DocdbClusterParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DocdbClusterParameterGroupSpec](#docdbclusterparametergroupspec)***||
| `status` | ***[DocdbClusterParameterGroupStatus](#docdbclusterparametergroupstatus)***||
## DocdbClusterParameterGroupSpec

Appears on:[DocdbClusterParameterGroup](#docdbclusterparametergroup), [DocdbClusterParameterGroupStatus](#docdbclusterparametergroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `parameter` | ***[[]DocdbClusterParameterGroupSpecParameter](#docdbclusterparametergroupspecparameter)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## DocdbClusterParameterGroupSpecParameter

Appears on:[DocdbClusterParameterGroupSpec](#docdbclusterparametergroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyMethod` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `value` | ***string***||
## DocdbClusterParameterGroupStatus

Appears on:[DocdbClusterParameterGroup](#docdbclusterparametergroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DocdbClusterParameterGroupSpec](#docdbclusterparametergroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DocdbClusterParameterGroupStatus](#docdbclusterparametergroupstatus)

---
