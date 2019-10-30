---
title: PubsubSubscriptionIamMember
menu:
  docs_{{ .version }}:
    identifier: pubsubsubscriptioniammember-google.kubeform.com
    name: PubsubSubscriptionIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PubsubSubscriptionIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `PubsubSubscriptionIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PubsubSubscriptionIamMemberSpec](#pubsubsubscriptioniammemberspec)***||
| `status` | ***[PubsubSubscriptionIamMemberStatus](#pubsubsubscriptioniammemberstatus)***||
## Phase(`string` alias)

Appears on:[PubsubSubscriptionIamMemberStatus](#pubsubsubscriptioniammemberstatus)

## PubsubSubscriptionIamMemberSpec

Appears on:[PubsubSubscriptionIamMember](#pubsubsubscriptioniammember), [PubsubSubscriptionIamMemberStatus](#pubsubsubscriptioniammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `subscription` | ***string***||
## PubsubSubscriptionIamMemberStatus

Appears on:[PubsubSubscriptionIamMember](#pubsubsubscriptioniammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PubsubSubscriptionIamMemberSpec](#pubsubsubscriptioniammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
