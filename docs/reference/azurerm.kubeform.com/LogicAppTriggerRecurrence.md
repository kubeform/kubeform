## LogicAppTriggerRecurrence
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogicAppTriggerRecurrence` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogicAppTriggerRecurrenceSpec](#LogicAppTriggerRecurrenceSpec)***||
| `status` | ***[LogicAppTriggerRecurrenceStatus](#LogicAppTriggerRecurrenceStatus)***||
## LogicAppTriggerRecurrenceSpec
##### (Appears on:[LogicAppTriggerRecurrence](#LogicAppTriggerRecurrence), [LogicAppTriggerRecurrenceStatus](#LogicAppTriggerRecurrenceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `frequency` | ***string***||
| `interval` | ***int***||
| `logicAppID` | ***string***||
| `name` | ***string***||
## LogicAppTriggerRecurrenceStatus
##### (Appears on:[LogicAppTriggerRecurrence](#LogicAppTriggerRecurrence))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogicAppTriggerRecurrenceSpec](#LogicAppTriggerRecurrenceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
