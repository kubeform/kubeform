---
title: SpannerInstanceIamPolicy
menu:
  docs_{{ .version }}:
    identifier: spannerinstanceiampolicy-google.kubeform.com
    name: SpannerInstanceIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SpannerInstanceIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SpannerInstanceIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpannerInstanceIamPolicySpec](#SpannerInstanceIamPolicySpec)***||
| `status` | ***[SpannerInstanceIamPolicyStatus](#SpannerInstanceIamPolicyStatus)***||
## SpannerInstanceIamPolicySpec
##### (Appears on:[SpannerInstanceIamPolicy](#SpannerInstanceIamPolicy), [SpannerInstanceIamPolicyStatus](#SpannerInstanceIamPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## SpannerInstanceIamPolicyStatus
##### (Appears on:[SpannerInstanceIamPolicy](#SpannerInstanceIamPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpannerInstanceIamPolicySpec](#SpannerInstanceIamPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
