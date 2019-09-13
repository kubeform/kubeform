---
title: GlueTrigger
menu:
  docs_{{ .version }}:
    identifier: gluetrigger-aws.kubeform.com
    name: GlueTrigger
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## GlueTrigger
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `GlueTrigger` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[GlueTriggerSpec](#GlueTriggerSpec)***||
| `status` | ***[GlueTriggerStatus](#GlueTriggerStatus)***||
## GlueTriggerSpec
##### (Appears on:[GlueTrigger](#GlueTrigger), [GlueTriggerStatus](#GlueTriggerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `actions` | ***[[]GlueTriggerSpecActions](#GlueTriggerSpecActions)***||
| `description` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `predicate` | ***[[]GlueTriggerSpecPredicate](#GlueTriggerSpecPredicate)***| ***(Optional)*** |
| `schedule` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## GlueTriggerSpecActions
##### (Appears on:[GlueTriggerSpec](#GlueTriggerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `arguments` | ***map[string]string***| ***(Optional)*** |
| `jobName` | ***string***||
| `timeout` | ***int***| ***(Optional)*** |
## GlueTriggerSpecPredicate
##### (Appears on:[GlueTriggerSpec](#GlueTriggerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `conditions` | ***[[]GlueTriggerSpecPredicateConditions](#GlueTriggerSpecPredicateConditions)***||
| `logical` | ***string***| ***(Optional)*** |
## GlueTriggerSpecPredicateConditions
##### (Appears on:[GlueTriggerSpecPredicate](#GlueTriggerSpecPredicate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `jobName` | ***string***||
| `logicalOperator` | ***string***| ***(Optional)*** |
| `state` | ***string***||
## GlueTriggerStatus
##### (Appears on:[GlueTrigger](#GlueTrigger))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueTriggerSpec](#GlueTriggerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
