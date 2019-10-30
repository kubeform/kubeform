---
title: CloudwatchLogMetricFilter
menu:
  docs_{{ .version }}:
    identifier: cloudwatchlogmetricfilter-aws.kubeform.com
    name: CloudwatchLogMetricFilter
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchLogMetricFilter
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchLogMetricFilter` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchLogMetricFilterSpec](#cloudwatchlogmetricfilterspec)***||
| `status` | ***[CloudwatchLogMetricFilterStatus](#cloudwatchlogmetricfilterstatus)***||
## CloudwatchLogMetricFilterSpec

Appears on:[CloudwatchLogMetricFilter](#cloudwatchlogmetricfilter), [CloudwatchLogMetricFilterStatus](#cloudwatchlogmetricfilterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `logGroupName` | ***string***||
| `metricTransformation` | ***[[]CloudwatchLogMetricFilterSpecMetricTransformation](#cloudwatchlogmetricfilterspecmetrictransformation)***||
| `name` | ***string***||
| `pattern` | ***string***||
## CloudwatchLogMetricFilterSpecMetricTransformation

Appears on:[CloudwatchLogMetricFilterSpec](#cloudwatchlogmetricfilterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `defaultValue` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `namespace` | ***string***||
| `value` | ***string***||
## CloudwatchLogMetricFilterStatus

Appears on:[CloudwatchLogMetricFilter](#cloudwatchlogmetricfilter)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchLogMetricFilterSpec](#cloudwatchlogmetricfilterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudwatchLogMetricFilterStatus](#cloudwatchlogmetricfilterstatus)

---
