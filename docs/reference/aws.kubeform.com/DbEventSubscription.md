---
title: DbEventSubscription
menu:
  docs_{{ .version }}:
    identifier: dbeventsubscription-aws.kubeform.com
    name: DbEventSubscription
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DbEventSubscription
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DbEventSubscription` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DbEventSubscriptionSpec](#dbeventsubscriptionspec)***||
| `status` | ***[DbEventSubscriptionStatus](#dbeventsubscriptionstatus)***||
## DbEventSubscriptionSpec

Appears on:[DbEventSubscription](#dbeventsubscription), [DbEventSubscriptionStatus](#dbeventsubscriptionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `customerAwsID` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `eventCategories` | ***[]string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `snsTopic` | ***string***||
| `sourceIDS` | ***[]string***| ***(Optional)*** |
| `sourceType` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## DbEventSubscriptionStatus

Appears on:[DbEventSubscription](#dbeventsubscription)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DbEventSubscriptionSpec](#dbeventsubscriptionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
