---
title: DataflowJob
menu:
  docs_{{ .version }}:
    identifier: dataflowjob-google.kubeform.com
    name: DataflowJob
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataflowJob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `DataflowJob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataflowJobSpec](#DataflowJobSpec)***||
| `status` | ***[DataflowJobStatus](#DataflowJobStatus)***||
## DataflowJobSpec

Appears on:[DataflowJob](#DataflowJob), [DataflowJobStatus](#DataflowJobStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `maxWorkers` | ***int***| ***(Optional)*** |
| `name` | ***string***||
| `onDelete` | ***string***| ***(Optional)*** |
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
| `tempGcsLocation` | ***string***||
| `templateGcsPath` | ***string***||
| `zone` | ***string***| ***(Optional)*** |
## DataflowJobStatus

Appears on:[DataflowJob](#DataflowJob)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataflowJobSpec](#DataflowJobSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
