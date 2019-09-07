---
title: PubsubTopic
menu:
  docs_v0.0.1:
    identifier: pubsubtopic-google.kubeform.com
    name: PubsubTopic
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## PubsubTopic
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `PubsubTopic` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PubsubTopicSpec](#PubsubTopicSpec)***||
| `status` | ***[PubsubTopicStatus](#PubsubTopicStatus)***||
## PubsubTopicSpec
##### (Appears on:[PubsubTopic](#PubsubTopic), [PubsubTopicStatus](#PubsubTopicStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## PubsubTopicStatus
##### (Appears on:[PubsubTopic](#PubsubTopic))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PubsubTopicSpec](#PubsubTopicSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
