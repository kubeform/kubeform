---
title: NeptuneParameterGroup
menu:
  docs_{{ .version }}:
    identifier: neptuneparametergroup-aws.kubeform.com
    name: NeptuneParameterGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NeptuneParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NeptuneParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NeptuneParameterGroupSpec](#neptuneparametergroupspec)***||
| `status` | ***[NeptuneParameterGroupStatus](#neptuneparametergroupstatus)***||
## NeptuneParameterGroupSpec

Appears on:[NeptuneParameterGroup](#neptuneparametergroup), [NeptuneParameterGroupStatus](#neptuneparametergroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***||
| `parameter` | ***[[]NeptuneParameterGroupSpecParameter](#neptuneparametergroupspecparameter)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## NeptuneParameterGroupSpecParameter

Appears on:[NeptuneParameterGroupSpec](#neptuneparametergroupspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyMethod` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `value` | ***string***||
## NeptuneParameterGroupStatus

Appears on:[NeptuneParameterGroup](#neptuneparametergroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NeptuneParameterGroupSpec](#neptuneparametergroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NeptuneParameterGroupStatus](#neptuneparametergroupstatus)

---
