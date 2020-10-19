---
title: ComputeNodeTemplate
menu:
  docs_{{ .version }}:
    identifier: computenodetemplate-google.kubeform.com
    name: ComputeNodeTemplate
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeNodeTemplate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeNodeTemplate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeNodeTemplateSpec](#computenodetemplatespec)***||
| `status` | ***[ComputeNodeTemplateStatus](#computenodetemplatestatus)***||
## ComputeNodeTemplateSpec

Appears on:[ComputeNodeTemplate](#computenodetemplate), [ComputeNodeTemplateStatus](#computenodetemplatestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `nodeAffinityLabels` | ***map[string]string***| ***(Optional)*** |
| `nodeType` | ***string***| ***(Optional)*** |
| `nodeTypeFlexibility` | ***[[]ComputeNodeTemplateSpecNodeTypeFlexibility](#computenodetemplatespecnodetypeflexibility)***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeNodeTemplateSpecNodeTypeFlexibility

Appears on:[ComputeNodeTemplateSpec](#computenodetemplatespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cpus` | ***string***| ***(Optional)*** |
| `localSsd` | ***string***| ***(Optional)*** |
| `memory` | ***string***| ***(Optional)*** |
## ComputeNodeTemplateStatus

Appears on:[ComputeNodeTemplate](#computenodetemplate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeNodeTemplateSpec](#computenodetemplatespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeNodeTemplateStatus](#computenodetemplatestatus)

---
