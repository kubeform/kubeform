---
title: CloudwatchLogDestination
menu:
  docs_{{ .version }}:
    identifier: cloudwatchlogdestination-aws.kubeform.com
    name: CloudwatchLogDestination
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchLogDestination
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchLogDestination` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchLogDestinationSpec](#cloudwatchlogdestinationspec)***||
| `status` | ***[CloudwatchLogDestinationStatus](#cloudwatchlogdestinationstatus)***||
## CloudwatchLogDestinationSpec

Appears on:[CloudwatchLogDestination](#cloudwatchlogdestination), [CloudwatchLogDestinationStatus](#cloudwatchlogdestinationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `roleArn` | ***string***||
| `targetArn` | ***string***||
## CloudwatchLogDestinationStatus

Appears on:[CloudwatchLogDestination](#cloudwatchlogdestination)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchLogDestinationSpec](#cloudwatchlogdestinationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudwatchLogDestinationStatus](#cloudwatchlogdestinationstatus)

---
