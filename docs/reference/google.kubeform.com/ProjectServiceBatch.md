---
title: ProjectServiceBatch
menu:
  docs_{{ .version }}:
    identifier: projectservicebatch-google.kubeform.com
    name: ProjectServiceBatch
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ProjectServiceBatch
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ProjectServiceBatch` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectServiceBatchSpec](#projectservicebatchspec)***||
| `status` | ***[ProjectServiceBatchStatus](#projectservicebatchstatus)***||
## Phase(`string` alias)

Appears on:[ProjectServiceBatchStatus](#projectservicebatchstatus)

## ProjectServiceBatchSpec

Appears on:[ProjectServiceBatch](#projectservicebatch), [ProjectServiceBatchStatus](#projectservicebatchstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `disableOnDestroy` | ***bool***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `services` | ***[]string***||
## ProjectServiceBatchStatus

Appears on:[ProjectServiceBatch](#projectservicebatch)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectServiceBatchSpec](#projectservicebatchspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
