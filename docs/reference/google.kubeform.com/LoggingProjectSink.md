---
title: LoggingProjectSink
menu:
  docs_{{ .version }}:
    identifier: loggingprojectsink-google.kubeform.com
    name: LoggingProjectSink
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LoggingProjectSink
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `LoggingProjectSink` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoggingProjectSinkSpec](#loggingprojectsinkspec)***||
| `status` | ***[LoggingProjectSinkStatus](#loggingprojectsinkstatus)***||
## LoggingProjectSinkSpec

Appears on:[LoggingProjectSink](#loggingprojectsink), [LoggingProjectSinkStatus](#loggingprojectsinkstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `destination` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `uniqueWriterIdentity` | ***bool***| ***(Optional)*** |
| `writerIdentity` | ***string***| ***(Optional)*** |
## LoggingProjectSinkStatus

Appears on:[LoggingProjectSink](#loggingprojectsink)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoggingProjectSinkSpec](#loggingprojectsinkspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LoggingProjectSinkStatus](#loggingprojectsinkstatus)

---
