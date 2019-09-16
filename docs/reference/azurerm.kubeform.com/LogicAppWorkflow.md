---
title: LogicAppWorkflow
menu:
  docs_{{ .version }}:
    identifier: logicappworkflow-azurerm.kubeform.com
    name: LogicAppWorkflow
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LogicAppWorkflow
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogicAppWorkflow` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogicAppWorkflowSpec](#LogicAppWorkflowSpec)***||
| `status` | ***[LogicAppWorkflowStatus](#LogicAppWorkflowStatus)***||
## LogicAppWorkflowSpec

Appears on:[LogicAppWorkflow](#LogicAppWorkflow), [LogicAppWorkflowStatus](#LogicAppWorkflowStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `parameters` | ***map[string]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `workflowSchema` | ***string***| ***(Optional)*** |
| `workflowVersion` | ***string***| ***(Optional)*** |
## LogicAppWorkflowStatus

Appears on:[LogicAppWorkflow](#LogicAppWorkflow)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogicAppWorkflowSpec](#LogicAppWorkflowSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
