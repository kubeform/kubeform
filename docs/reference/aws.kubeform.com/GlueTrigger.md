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
| `spec` | ***[GlueTriggerSpec](#gluetriggerspec)***||
| `status` | ***[GlueTriggerStatus](#gluetriggerstatus)***||
## GlueTriggerSpec

Appears on:[GlueTrigger](#gluetrigger), [GlueTriggerStatus](#gluetriggerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `actions` | ***[[]GlueTriggerSpecActions](#gluetriggerspecactions)***||
| `description` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `predicate` | ***[[]GlueTriggerSpecPredicate](#gluetriggerspecpredicate)***| ***(Optional)*** |
| `schedule` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## GlueTriggerSpecActions

Appears on:[GlueTriggerSpec](#gluetriggerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `arguments` | ***map[string]string***| ***(Optional)*** |
| `jobName` | ***string***||
| `timeout` | ***int***| ***(Optional)*** |
## GlueTriggerSpecPredicate

Appears on:[GlueTriggerSpec](#gluetriggerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `conditions` | ***[[]GlueTriggerSpecPredicateConditions](#gluetriggerspecpredicateconditions)***||
| `logical` | ***string***| ***(Optional)*** |
## GlueTriggerSpecPredicateConditions

Appears on:[GlueTriggerSpecPredicate](#gluetriggerspecpredicate)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `jobName` | ***string***||
| `logicalOperator` | ***string***| ***(Optional)*** |
| `state` | ***string***||
## GlueTriggerStatus

Appears on:[GlueTrigger](#gluetrigger)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[GlueTriggerSpec](#gluetriggerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[GlueTriggerStatus](#gluetriggerstatus)

---
