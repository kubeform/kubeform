---
title: InspectorAssessmentTarget
menu:
  docs_{{ .version }}:
    identifier: inspectorassessmenttarget-aws.kubeform.com
    name: InspectorAssessmentTarget
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## InspectorAssessmentTarget
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `InspectorAssessmentTarget` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[InspectorAssessmentTargetSpec](#InspectorAssessmentTargetSpec)***||
| `status` | ***[InspectorAssessmentTargetStatus](#InspectorAssessmentTargetStatus)***||
## InspectorAssessmentTargetSpec

Appears on:[InspectorAssessmentTarget](#InspectorAssessmentTarget), [InspectorAssessmentTargetStatus](#InspectorAssessmentTargetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupArn` | ***string***| ***(Optional)*** |
## InspectorAssessmentTargetStatus

Appears on:[InspectorAssessmentTarget](#InspectorAssessmentTarget)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[InspectorAssessmentTargetSpec](#InspectorAssessmentTargetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
