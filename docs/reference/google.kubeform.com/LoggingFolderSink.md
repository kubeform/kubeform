---
title: LoggingFolderSink
menu:
  docs_{{ .version }}:
    identifier: loggingfoldersink-google.kubeform.com
    name: LoggingFolderSink
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LoggingFolderSink
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `LoggingFolderSink` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoggingFolderSinkSpec](#loggingfoldersinkspec)***||
| `status` | ***[LoggingFolderSinkStatus](#loggingfoldersinkstatus)***||
## LoggingFolderSinkSpec

Appears on:[LoggingFolderSink](#loggingfoldersink), [LoggingFolderSinkStatus](#loggingfoldersinkstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `destination` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `folder` | ***string***||
| `includeChildren` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `writerIdentity` | ***string***| ***(Optional)*** |
## LoggingFolderSinkStatus

Appears on:[LoggingFolderSink](#loggingfoldersink)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoggingFolderSinkSpec](#loggingfoldersinkspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LoggingFolderSinkStatus](#loggingfoldersinkstatus)

---
