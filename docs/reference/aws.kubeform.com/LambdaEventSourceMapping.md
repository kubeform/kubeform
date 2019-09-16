---
title: LambdaEventSourceMapping
menu:
  docs_{{ .version }}:
    identifier: lambdaeventsourcemapping-aws.kubeform.com
    name: LambdaEventSourceMapping
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LambdaEventSourceMapping
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LambdaEventSourceMapping` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LambdaEventSourceMappingSpec](#LambdaEventSourceMappingSpec)***||
| `status` | ***[LambdaEventSourceMappingStatus](#LambdaEventSourceMappingStatus)***||
## LambdaEventSourceMappingSpec

Appears on:[LambdaEventSourceMapping](#LambdaEventSourceMapping), [LambdaEventSourceMappingStatus](#LambdaEventSourceMappingStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `batchSize` | ***int***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `eventSourceArn` | ***string***||
| `functionArn` | ***string***| ***(Optional)*** |
| `functionName` | ***string***||
| `lastModified` | ***string***| ***(Optional)*** |
| `lastProcessingResult` | ***string***| ***(Optional)*** |
| `startingPosition` | ***string***| ***(Optional)*** |
| `startingPositionTimestamp` | ***string***| ***(Optional)*** |
| `state` | ***string***| ***(Optional)*** |
| `stateTransitionReason` | ***string***| ***(Optional)*** |
| `uuid` | ***string***| ***(Optional)*** |
## LambdaEventSourceMappingStatus

Appears on:[LambdaEventSourceMapping](#LambdaEventSourceMapping)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LambdaEventSourceMappingSpec](#LambdaEventSourceMappingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
