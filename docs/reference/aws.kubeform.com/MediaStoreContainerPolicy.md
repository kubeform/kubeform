---
title: MediaStoreContainerPolicy
menu:
  docs_v0.0.1:
    identifier: mediastorecontainerpolicy-aws.kubeform.com
    name: MediaStoreContainerPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## MediaStoreContainerPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MediaStoreContainerPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MediaStoreContainerPolicySpec](#MediaStoreContainerPolicySpec)***||
| `status` | ***[MediaStoreContainerPolicyStatus](#MediaStoreContainerPolicyStatus)***||
## MediaStoreContainerPolicySpec
##### (Appears on:[MediaStoreContainerPolicy](#MediaStoreContainerPolicy), [MediaStoreContainerPolicyStatus](#MediaStoreContainerPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `containerName` | ***string***||
| `policy` | ***string***||
## MediaStoreContainerPolicyStatus
##### (Appears on:[MediaStoreContainerPolicy](#MediaStoreContainerPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MediaStoreContainerPolicySpec](#MediaStoreContainerPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
