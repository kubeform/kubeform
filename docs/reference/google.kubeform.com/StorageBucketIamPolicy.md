---
title: StorageBucketIamPolicy
menu:
  docs_{{ .version }}:
    identifier: storagebucketiampolicy-google.kubeform.com
    name: StorageBucketIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageBucketIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageBucketIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageBucketIamPolicySpec](#StorageBucketIamPolicySpec)***||
| `status` | ***[StorageBucketIamPolicyStatus](#StorageBucketIamPolicyStatus)***||
## StorageBucketIamPolicySpec

Appears on:[StorageBucketIamPolicy](#StorageBucketIamPolicy), [StorageBucketIamPolicyStatus](#StorageBucketIamPolicyStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
## StorageBucketIamPolicyStatus

Appears on:[StorageBucketIamPolicy](#StorageBucketIamPolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageBucketIamPolicySpec](#StorageBucketIamPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
