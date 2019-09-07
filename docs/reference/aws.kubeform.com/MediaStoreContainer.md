---
title: MediaStoreContainer
menu:
  docs_v0.0.1:
    identifier: mediastorecontainer-aws.kubeform.com
    name: MediaStoreContainer
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## MediaStoreContainer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `MediaStoreContainer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MediaStoreContainerSpec](#MediaStoreContainerSpec)***||
| `status` | ***[MediaStoreContainerStatus](#MediaStoreContainerStatus)***||
## MediaStoreContainerSpec
##### (Appears on:[MediaStoreContainer](#MediaStoreContainer), [MediaStoreContainerStatus](#MediaStoreContainerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## MediaStoreContainerStatus
##### (Appears on:[MediaStoreContainer](#MediaStoreContainer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MediaStoreContainerSpec](#MediaStoreContainerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
