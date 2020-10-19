---
title: DataFactoryTriggerSchedule
menu:
  docs_{{ .version }}:
    identifier: datafactorytriggerschedule-azurerm.kubeform.com
    name: DataFactoryTriggerSchedule
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataFactoryTriggerSchedule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataFactoryTriggerSchedule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataFactoryTriggerScheduleSpec](#datafactorytriggerschedulespec)***||
| `status` | ***[DataFactoryTriggerScheduleStatus](#datafactorytriggerschedulestatus)***||
## DataFactoryTriggerScheduleSpec

Appears on:[DataFactoryTriggerSchedule](#datafactorytriggerschedule), [DataFactoryTriggerScheduleStatus](#datafactorytriggerschedulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `annotations` | ***[]string***| ***(Optional)*** |
| `dataFactoryName` | ***string***||
| `endTime` | ***string***| ***(Optional)*** |
| `frequency` | ***string***| ***(Optional)*** |
| `interval` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `pipelineName` | ***string***||
| `pipelineParameters` | ***map[string]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `startTime` | ***string***| ***(Optional)*** |
## DataFactoryTriggerScheduleStatus

Appears on:[DataFactoryTriggerSchedule](#datafactorytriggerschedule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataFactoryTriggerScheduleSpec](#datafactorytriggerschedulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DataFactoryTriggerScheduleStatus](#datafactorytriggerschedulestatus)

---
