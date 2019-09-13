---
title: SnsTopicSubscription
menu:
  docs_{{ .version }}:
    identifier: snstopicsubscription-aws.kubeform.com
    name: SnsTopicSubscription
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SnsTopicSubscription
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SnsTopicSubscription` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SnsTopicSubscriptionSpec](#SnsTopicSubscriptionSpec)***||
| `status` | ***[SnsTopicSubscriptionStatus](#SnsTopicSubscriptionStatus)***||
## SnsTopicSubscriptionSpec
##### (Appears on:[SnsTopicSubscription](#SnsTopicSubscription), [SnsTopicSubscriptionStatus](#SnsTopicSubscriptionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `confirmationTimeoutInMinutes` | ***int***| ***(Optional)*** |
| `deliveryPolicy` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***||
| `endpointAutoConfirms` | ***bool***| ***(Optional)*** |
| `filterPolicy` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `rawMessageDelivery` | ***bool***| ***(Optional)*** |
| `topicArn` | ***string***||
## SnsTopicSubscriptionStatus
##### (Appears on:[SnsTopicSubscription](#SnsTopicSubscription))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SnsTopicSubscriptionSpec](#SnsTopicSubscriptionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
