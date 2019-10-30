---
title: CloudwatchLogDestinationPolicy
menu:
  docs_{{ .version }}:
    identifier: cloudwatchlogdestinationpolicy-aws.kubeform.com
    name: CloudwatchLogDestinationPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudwatchLogDestinationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudwatchLogDestinationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudwatchLogDestinationPolicySpec](#cloudwatchlogdestinationpolicyspec)***||
| `status` | ***[CloudwatchLogDestinationPolicyStatus](#cloudwatchlogdestinationpolicystatus)***||
## CloudwatchLogDestinationPolicySpec

Appears on:[CloudwatchLogDestinationPolicy](#cloudwatchlogdestinationpolicy), [CloudwatchLogDestinationPolicyStatus](#cloudwatchlogdestinationpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessPolicy` | ***string***||
| `destinationName` | ***string***||
## CloudwatchLogDestinationPolicyStatus

Appears on:[CloudwatchLogDestinationPolicy](#cloudwatchlogdestinationpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudwatchLogDestinationPolicySpec](#cloudwatchlogdestinationpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudwatchLogDestinationPolicyStatus](#cloudwatchlogdestinationpolicystatus)

---
