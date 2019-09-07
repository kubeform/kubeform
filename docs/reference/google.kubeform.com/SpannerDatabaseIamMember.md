---
title: SpannerDatabaseIamMember
menu:
  docs_v0.0.1:
    identifier: spannerdatabaseiammember-google.kubeform.com
    name: SpannerDatabaseIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SpannerDatabaseIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SpannerDatabaseIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpannerDatabaseIamMemberSpec](#SpannerDatabaseIamMemberSpec)***||
| `status` | ***[SpannerDatabaseIamMemberStatus](#SpannerDatabaseIamMemberStatus)***||
## SpannerDatabaseIamMemberSpec
##### (Appears on:[SpannerDatabaseIamMember](#SpannerDatabaseIamMember), [SpannerDatabaseIamMemberStatus](#SpannerDatabaseIamMemberStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `database` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## SpannerDatabaseIamMemberStatus
##### (Appears on:[SpannerDatabaseIamMember](#SpannerDatabaseIamMember))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpannerDatabaseIamMemberSpec](#SpannerDatabaseIamMemberSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
