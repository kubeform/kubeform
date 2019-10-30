---
title: PubsubSubscriptionIamBinding
menu:
  docs_{{ .version }}:
    identifier: pubsubsubscriptioniambinding-google.kubeform.com
    name: PubsubSubscriptionIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PubsubSubscriptionIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `PubsubSubscriptionIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PubsubSubscriptionIamBindingSpec](#pubsubsubscriptioniambindingspec)***||
| `status` | ***[PubsubSubscriptionIamBindingStatus](#pubsubsubscriptioniambindingstatus)***||
## Phase(`string` alias)

Appears on:[PubsubSubscriptionIamBindingStatus](#pubsubsubscriptioniambindingstatus)

## PubsubSubscriptionIamBindingSpec

Appears on:[PubsubSubscriptionIamBinding](#pubsubsubscriptioniambinding), [PubsubSubscriptionIamBindingStatus](#pubsubsubscriptioniambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `subscription` | ***string***||
## PubsubSubscriptionIamBindingStatus

Appears on:[PubsubSubscriptionIamBinding](#pubsubsubscriptioniambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PubsubSubscriptionIamBindingSpec](#pubsubsubscriptioniambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
