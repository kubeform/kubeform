---
title: SesIdentityNotificationTopic
menu:
  docs_v0.0.1:
    identifier: sesidentitynotificationtopic-aws.kubeform.com
    name: SesIdentityNotificationTopic
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SesIdentityNotificationTopic
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesIdentityNotificationTopic` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesIdentityNotificationTopicSpec](#SesIdentityNotificationTopicSpec)***||
| `status` | ***[SesIdentityNotificationTopicStatus](#SesIdentityNotificationTopicStatus)***||
## SesIdentityNotificationTopicSpec
##### (Appears on:[SesIdentityNotificationTopic](#SesIdentityNotificationTopic), [SesIdentityNotificationTopicStatus](#SesIdentityNotificationTopicStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `identity` | ***string***||
| `notificationType` | ***string***||
| `topicArn` | ***string***| ***(Optional)*** |
## SesIdentityNotificationTopicStatus
##### (Appears on:[SesIdentityNotificationTopic](#SesIdentityNotificationTopic))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesIdentityNotificationTopicSpec](#SesIdentityNotificationTopicSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
