---
title: PubsubTopicIamBinding
menu:
  docs_{{ .version }}:
    identifier: pubsubtopiciambinding-google.kubeform.com
    name: PubsubTopicIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PubsubTopicIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `PubsubTopicIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PubsubTopicIamBindingSpec](#pubsubtopiciambindingspec)***||
| `status` | ***[PubsubTopicIamBindingStatus](#pubsubtopiciambindingstatus)***||
## PubsubTopicIamBindingSpec

Appears on:[PubsubTopicIamBinding](#pubsubtopiciambinding), [PubsubTopicIamBindingStatus](#pubsubtopiciambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `topic` | ***string***||
## PubsubTopicIamBindingStatus

Appears on:[PubsubTopicIamBinding](#pubsubtopiciambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PubsubTopicIamBindingSpec](#pubsubtopiciambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
