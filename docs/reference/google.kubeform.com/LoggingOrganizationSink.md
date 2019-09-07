---
title: LoggingOrganizationSink
menu:
  docs_v0.0.1:
    identifier: loggingorganizationsink-google.kubeform.com
    name: LoggingOrganizationSink
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## LoggingOrganizationSink
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `LoggingOrganizationSink` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoggingOrganizationSinkSpec](#LoggingOrganizationSinkSpec)***||
| `status` | ***[LoggingOrganizationSinkStatus](#LoggingOrganizationSinkStatus)***||
## LoggingOrganizationSinkSpec
##### (Appears on:[LoggingOrganizationSink](#LoggingOrganizationSink), [LoggingOrganizationSinkStatus](#LoggingOrganizationSinkStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `destination` | ***string***||
| `filter` | ***string***| ***(Optional)*** |
| `includeChildren` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `orgID` | ***string***||
| `writerIdentity` | ***string***| ***(Optional)*** |
## LoggingOrganizationSinkStatus
##### (Appears on:[LoggingOrganizationSink](#LoggingOrganizationSink))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoggingOrganizationSinkSpec](#LoggingOrganizationSinkSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
