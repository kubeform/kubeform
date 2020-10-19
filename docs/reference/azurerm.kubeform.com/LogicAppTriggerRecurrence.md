---
title: LogicAppTriggerRecurrence
menu:
  docs_{{ .version }}:
    identifier: logicapptriggerrecurrence-azurerm.kubeform.com
    name: LogicAppTriggerRecurrence
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LogicAppTriggerRecurrence
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogicAppTriggerRecurrence` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogicAppTriggerRecurrenceSpec](#logicapptriggerrecurrencespec)***||
| `status` | ***[LogicAppTriggerRecurrenceStatus](#logicapptriggerrecurrencestatus)***||
## LogicAppTriggerRecurrenceSpec

Appears on:[LogicAppTriggerRecurrence](#logicapptriggerrecurrence), [LogicAppTriggerRecurrenceStatus](#logicapptriggerrecurrencestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `frequency` | ***string***||
| `interval` | ***int64***||
| `logicAppID` | ***string***||
| `name` | ***string***||
| `startTime` | ***string***| ***(Optional)*** |
## LogicAppTriggerRecurrenceStatus

Appears on:[LogicAppTriggerRecurrence](#logicapptriggerrecurrence)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogicAppTriggerRecurrenceSpec](#logicapptriggerrecurrencespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LogicAppTriggerRecurrenceStatus](#logicapptriggerrecurrencestatus)

---
