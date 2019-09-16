---
title: LogicAppActionCustom
menu:
  docs_{{ .version }}:
    identifier: logicappactioncustom-azurerm.kubeform.com
    name: LogicAppActionCustom
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LogicAppActionCustom
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogicAppActionCustom` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogicAppActionCustomSpec](#logicappactioncustomspec)***||
| `status` | ***[LogicAppActionCustomStatus](#logicappactioncustomstatus)***||
## LogicAppActionCustomSpec

Appears on:[LogicAppActionCustom](#logicappactioncustom), [LogicAppActionCustomStatus](#logicappactioncustomstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `body` | ***string***||
| `logicAppID` | ***string***||
| `name` | ***string***||
## LogicAppActionCustomStatus

Appears on:[LogicAppActionCustom](#logicappactioncustom)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogicAppActionCustomSpec](#logicappactioncustomspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
