---
title: SpannerInstanceIamMember
menu:
  docs_{{ .version }}:
    identifier: spannerinstanceiammember-google.kubeform.com
    name: SpannerInstanceIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SpannerInstanceIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SpannerInstanceIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpannerInstanceIamMemberSpec](#SpannerInstanceIamMemberSpec)***||
| `status` | ***[SpannerInstanceIamMemberStatus](#SpannerInstanceIamMemberStatus)***||
## SpannerInstanceIamMemberSpec
##### (Appears on:[SpannerInstanceIamMember](#SpannerInstanceIamMember), [SpannerInstanceIamMemberStatus](#SpannerInstanceIamMemberStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## SpannerInstanceIamMemberStatus
##### (Appears on:[SpannerInstanceIamMember](#SpannerInstanceIamMember))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpannerInstanceIamMemberSpec](#SpannerInstanceIamMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
