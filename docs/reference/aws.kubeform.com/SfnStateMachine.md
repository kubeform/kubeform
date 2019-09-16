---
title: SfnStateMachine
menu:
  docs_{{ .version }}:
    identifier: sfnstatemachine-aws.kubeform.com
    name: SfnStateMachine
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SfnStateMachine
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SfnStateMachine` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SfnStateMachineSpec](#sfnstatemachinespec)***||
| `status` | ***[SfnStateMachineStatus](#sfnstatemachinestatus)***||
## SfnStateMachineSpec

Appears on:[SfnStateMachine](#sfnstatemachine), [SfnStateMachineStatus](#sfnstatemachinestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `creationDate` | ***string***| ***(Optional)*** |
| `definition` | ***string***||
| `name` | ***string***||
| `roleArn` | ***string***||
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SfnStateMachineStatus

Appears on:[SfnStateMachine](#sfnstatemachine)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SfnStateMachineSpec](#sfnstatemachinespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
