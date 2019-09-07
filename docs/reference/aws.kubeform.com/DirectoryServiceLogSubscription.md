---
title: DirectoryServiceLogSubscription
menu:
  docs_v0.0.1:
    identifier: directoryservicelogsubscription-aws.kubeform.com
    name: DirectoryServiceLogSubscription
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DirectoryServiceLogSubscription
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DirectoryServiceLogSubscription` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DirectoryServiceLogSubscriptionSpec](#DirectoryServiceLogSubscriptionSpec)***||
| `status` | ***[DirectoryServiceLogSubscriptionStatus](#DirectoryServiceLogSubscriptionStatus)***||
## DirectoryServiceLogSubscriptionSpec
##### (Appears on:[DirectoryServiceLogSubscription](#DirectoryServiceLogSubscription), [DirectoryServiceLogSubscriptionStatus](#DirectoryServiceLogSubscriptionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `directoryID` | ***string***||
| `logGroupName` | ***string***||
## DirectoryServiceLogSubscriptionStatus
##### (Appears on:[DirectoryServiceLogSubscription](#DirectoryServiceLogSubscription))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DirectoryServiceLogSubscriptionSpec](#DirectoryServiceLogSubscriptionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
