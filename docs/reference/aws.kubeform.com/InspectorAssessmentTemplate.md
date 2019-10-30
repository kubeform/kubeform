---
title: InspectorAssessmentTemplate
menu:
  docs_{{ .version }}:
    identifier: inspectorassessmenttemplate-aws.kubeform.com
    name: InspectorAssessmentTemplate
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## InspectorAssessmentTemplate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `InspectorAssessmentTemplate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[InspectorAssessmentTemplateSpec](#inspectorassessmenttemplatespec)***||
| `status` | ***[InspectorAssessmentTemplateStatus](#inspectorassessmenttemplatestatus)***||
## InspectorAssessmentTemplateSpec

Appears on:[InspectorAssessmentTemplate](#inspectorassessmenttemplate), [InspectorAssessmentTemplateStatus](#inspectorassessmenttemplatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `duration` | ***int***||
| `name` | ***string***||
| `rulesPackageArns` | ***[]string***||
| `targetArn` | ***string***||
## InspectorAssessmentTemplateStatus

Appears on:[InspectorAssessmentTemplate](#inspectorassessmenttemplate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[InspectorAssessmentTemplateSpec](#inspectorassessmenttemplatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[InspectorAssessmentTemplateStatus](#inspectorassessmenttemplatestatus)

---
