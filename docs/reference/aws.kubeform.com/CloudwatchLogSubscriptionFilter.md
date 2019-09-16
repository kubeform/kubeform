---
title: CloudwatchLogSubscriptionFilter
menu:
  docs_{{ .version }}:
    identifier: cloudwatchlogsubscriptionfilter-aws.kubeform.com
    name: CloudwatchLogSubscriptionFilter
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchLogSubscriptionFilter
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchLogSubscriptionFilter` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchLogSubscriptionFilterSpec](#CloudwatchLogSubscriptionFilterSpec)***||
| `status` | ***[CloudwatchLogSubscriptionFilterStatus](#CloudwatchLogSubscriptionFilterStatus)***||
## CloudwatchLogSubscriptionFilterSpec

Appears on:[CloudwatchLogSubscriptionFilter](#CloudwatchLogSubscriptionFilter), [CloudwatchLogSubscriptionFilterStatus](#CloudwatchLogSubscriptionFilterStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `destinationArn` | ***string***||
| `distribution` | ***string***| ***(Optional)*** |
| `filterPattern` | ***string***||
| `logGroupName` | ***string***||
| `name` | ***string***||
| `roleArn` | ***string***| ***(Optional)*** |
## CloudwatchLogSubscriptionFilterStatus

Appears on:[CloudwatchLogSubscriptionFilter](#CloudwatchLogSubscriptionFilter)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchLogSubscriptionFilterSpec](#CloudwatchLogSubscriptionFilterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
