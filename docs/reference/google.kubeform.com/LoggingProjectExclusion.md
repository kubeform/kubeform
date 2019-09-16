---
title: LoggingProjectExclusion
menu:
  docs_{{ .version }}:
    identifier: loggingprojectexclusion-google.kubeform.com
    name: LoggingProjectExclusion
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LoggingProjectExclusion
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `LoggingProjectExclusion` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoggingProjectExclusionSpec](#LoggingProjectExclusionSpec)***||
| `status` | ***[LoggingProjectExclusionStatus](#LoggingProjectExclusionStatus)***||
## LoggingProjectExclusionSpec

Appears on:[LoggingProjectExclusion](#LoggingProjectExclusion), [LoggingProjectExclusionStatus](#LoggingProjectExclusionStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `disabled` | ***bool***| ***(Optional)*** |
| `filter` | ***string***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## LoggingProjectExclusionStatus

Appears on:[LoggingProjectExclusion](#LoggingProjectExclusion)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoggingProjectExclusionSpec](#LoggingProjectExclusionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
