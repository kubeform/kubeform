---
title: SnsTopicPolicy
menu:
  docs_{{ .version }}:
    identifier: snstopicpolicy-aws.kubeform.com
    name: SnsTopicPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SnsTopicPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SnsTopicPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SnsTopicPolicySpec](#snstopicpolicyspec)***||
| `status` | ***[SnsTopicPolicyStatus](#snstopicpolicystatus)***||
## Phase(`string` alias)

Appears on:[SnsTopicPolicyStatus](#snstopicpolicystatus)

## SnsTopicPolicySpec

Appears on:[SnsTopicPolicy](#snstopicpolicy), [SnsTopicPolicyStatus](#snstopicpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***||
| `policy` | ***string***||
## SnsTopicPolicyStatus

Appears on:[SnsTopicPolicy](#snstopicpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SnsTopicPolicySpec](#snstopicpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
