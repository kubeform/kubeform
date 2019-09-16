---
title: SpotDatafeedSubscription
menu:
  docs_{{ .version }}:
    identifier: spotdatafeedsubscription-aws.kubeform.com
    name: SpotDatafeedSubscription
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SpotDatafeedSubscription
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SpotDatafeedSubscription` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SpotDatafeedSubscriptionSpec](#SpotDatafeedSubscriptionSpec)***||
| `status` | ***[SpotDatafeedSubscriptionStatus](#SpotDatafeedSubscriptionStatus)***||
## SpotDatafeedSubscriptionSpec

Appears on:[SpotDatafeedSubscription](#SpotDatafeedSubscription), [SpotDatafeedSubscriptionStatus](#SpotDatafeedSubscriptionStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `prefix` | ***string***| ***(Optional)*** |
## SpotDatafeedSubscriptionStatus

Appears on:[SpotDatafeedSubscription](#SpotDatafeedSubscription)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SpotDatafeedSubscriptionSpec](#SpotDatafeedSubscriptionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
