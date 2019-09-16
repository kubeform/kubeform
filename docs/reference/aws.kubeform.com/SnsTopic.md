---
title: SnsTopic
menu:
  docs_{{ .version }}:
    identifier: snstopic-aws.kubeform.com
    name: SnsTopic
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SnsTopic
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SnsTopic` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SnsTopicSpec](#SnsTopicSpec)***||
| `status` | ***[SnsTopicStatus](#SnsTopicStatus)***||
## SnsTopicSpec

Appears on:[SnsTopic](#SnsTopic), [SnsTopicStatus](#SnsTopicStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationFailureFeedbackRoleArn` | ***string***| ***(Optional)*** |
| `applicationSuccessFeedbackRoleArn` | ***string***| ***(Optional)*** |
| `applicationSuccessFeedbackSampleRate` | ***int***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `deliveryPolicy` | ***string***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `httpFailureFeedbackRoleArn` | ***string***| ***(Optional)*** |
| `httpSuccessFeedbackRoleArn` | ***string***| ***(Optional)*** |
| `httpSuccessFeedbackSampleRate` | ***int***| ***(Optional)*** |
| `kmsMasterKeyID` | ***string***| ***(Optional)*** |
| `lambdaFailureFeedbackRoleArn` | ***string***| ***(Optional)*** |
| `lambdaSuccessFeedbackRoleArn` | ***string***| ***(Optional)*** |
| `lambdaSuccessFeedbackSampleRate` | ***int***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `policy` | ***string***| ***(Optional)*** |
| `sqsFailureFeedbackRoleArn` | ***string***| ***(Optional)*** |
| `sqsSuccessFeedbackRoleArn` | ***string***| ***(Optional)*** |
| `sqsSuccessFeedbackSampleRate` | ***int***| ***(Optional)*** |
## SnsTopicStatus

Appears on:[SnsTopic](#SnsTopic)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SnsTopicSpec](#SnsTopicSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
