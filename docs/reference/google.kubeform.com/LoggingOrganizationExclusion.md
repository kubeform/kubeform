---
title: LoggingOrganizationExclusion
menu:
  docs_{{ .version }}:
    identifier: loggingorganizationexclusion-google.kubeform.com
    name: LoggingOrganizationExclusion
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LoggingOrganizationExclusion
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `LoggingOrganizationExclusion` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoggingOrganizationExclusionSpec](#loggingorganizationexclusionspec)***||
| `status` | ***[LoggingOrganizationExclusionStatus](#loggingorganizationexclusionstatus)***||
## LoggingOrganizationExclusionSpec

Appears on:[LoggingOrganizationExclusion](#loggingorganizationexclusion), [LoggingOrganizationExclusionStatus](#loggingorganizationexclusionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `disabled` | ***bool***| ***(Optional)*** |
| `filter` | ***string***||
| `name` | ***string***||
| `orgID` | ***string***||
## LoggingOrganizationExclusionStatus

Appears on:[LoggingOrganizationExclusion](#loggingorganizationexclusion)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoggingOrganizationExclusionSpec](#loggingorganizationexclusionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
