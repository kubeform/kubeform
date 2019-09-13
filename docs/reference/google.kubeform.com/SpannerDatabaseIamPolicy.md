---
title: SpannerDatabaseIamPolicy
menu:
  docs_{{ .version }}:
    identifier: spannerdatabaseiampolicy-google.kubeform.com
    name: SpannerDatabaseIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SpannerDatabaseIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SpannerDatabaseIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpannerDatabaseIamPolicySpec](#SpannerDatabaseIamPolicySpec)***||
| `status` | ***[SpannerDatabaseIamPolicyStatus](#SpannerDatabaseIamPolicyStatus)***||
## SpannerDatabaseIamPolicySpec
##### (Appears on:[SpannerDatabaseIamPolicy](#SpannerDatabaseIamPolicy), [SpannerDatabaseIamPolicyStatus](#SpannerDatabaseIamPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `database` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## SpannerDatabaseIamPolicyStatus
##### (Appears on:[SpannerDatabaseIamPolicy](#SpannerDatabaseIamPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpannerDatabaseIamPolicySpec](#SpannerDatabaseIamPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
