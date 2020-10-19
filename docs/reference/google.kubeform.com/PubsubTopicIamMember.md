---
title: PubsubTopicIamMember
menu:
  docs_{{ .version }}:
    identifier: pubsubtopiciammember-google.kubeform.com
    name: PubsubTopicIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PubsubTopicIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `PubsubTopicIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PubsubTopicIamMemberSpec](#pubsubtopiciammemberspec)***||
| `status` | ***[PubsubTopicIamMemberStatus](#pubsubtopiciammemberstatus)***||
## Phase(`string` alias)

Appears on:[PubsubTopicIamMemberStatus](#pubsubtopiciammemberstatus)

## PubsubTopicIamMemberSpec

Appears on:[PubsubTopicIamMember](#pubsubtopiciammember), [PubsubTopicIamMemberStatus](#pubsubtopiciammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `topic` | ***string***||
## PubsubTopicIamMemberStatus

Appears on:[PubsubTopicIamMember](#pubsubtopiciammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PubsubTopicIamMemberSpec](#pubsubtopiciammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
