---
title: MediaStoreContainerPolicy
menu:
  docs_{{ .version }}:
    identifier: mediastorecontainerpolicy-aws.kubeform.com
    name: MediaStoreContainerPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MediaStoreContainerPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MediaStoreContainerPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MediaStoreContainerPolicySpec](#mediastorecontainerpolicyspec)***||
| `status` | ***[MediaStoreContainerPolicyStatus](#mediastorecontainerpolicystatus)***||
## MediaStoreContainerPolicySpec

Appears on:[MediaStoreContainerPolicy](#mediastorecontainerpolicy), [MediaStoreContainerPolicyStatus](#mediastorecontainerpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `containerName` | ***string***||
| `policy` | ***string***||
## MediaStoreContainerPolicyStatus

Appears on:[MediaStoreContainerPolicy](#mediastorecontainerpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MediaStoreContainerPolicySpec](#mediastorecontainerpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MediaStoreContainerPolicyStatus](#mediastorecontainerpolicystatus)

---
