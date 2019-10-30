---
title: SesTemplate
menu:
  docs_{{ .version }}:
    identifier: sestemplate-aws.kubeform.com
    name: SesTemplate
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SesTemplate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesTemplate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesTemplateSpec](#sestemplatespec)***||
| `status` | ***[SesTemplateStatus](#sestemplatestatus)***||
## Phase(`string` alias)

Appears on:[SesTemplateStatus](#sestemplatestatus)

## SesTemplateSpec

Appears on:[SesTemplate](#sestemplate), [SesTemplateStatus](#sestemplatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `html` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `subject` | ***string***| ***(Optional)*** |
| `text` | ***string***| ***(Optional)*** |
## SesTemplateStatus

Appears on:[SesTemplate](#sestemplate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesTemplateSpec](#sestemplatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
