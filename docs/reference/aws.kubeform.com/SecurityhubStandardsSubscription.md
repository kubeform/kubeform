---
title: SecurityhubStandardsSubscription
menu:
  docs_v0.0.1:
    identifier: securityhubstandardssubscription-aws.kubeform.com
    name: SecurityhubStandardsSubscription
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SecurityhubStandardsSubscription
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SecurityhubStandardsSubscription` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityhubStandardsSubscriptionSpec](#SecurityhubStandardsSubscriptionSpec)***||
| `status` | ***[SecurityhubStandardsSubscriptionStatus](#SecurityhubStandardsSubscriptionStatus)***||
## SecurityhubStandardsSubscriptionSpec
##### (Appears on:[SecurityhubStandardsSubscription](#SecurityhubStandardsSubscription), [SecurityhubStandardsSubscriptionStatus](#SecurityhubStandardsSubscriptionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `standardsArn` | ***string***||
## SecurityhubStandardsSubscriptionStatus
##### (Appears on:[SecurityhubStandardsSubscription](#SecurityhubStandardsSubscription))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityhubStandardsSubscriptionSpec](#SecurityhubStandardsSubscriptionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
