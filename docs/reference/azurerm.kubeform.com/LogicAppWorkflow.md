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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogicAppWorkflowSpec](#logicappworkflowspec)***||
| `status` | ***[LogicAppWorkflowStatus](#logicappworkflowstatus)***||
## LogicAppWorkflowSpec

Appears on:[LogicAppWorkflow](#logicappworkflow), [LogicAppWorkflowStatus](#logicappworkflowstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
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

Appears on:[LogicAppWorkflow](#logicappworkflow)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogicAppWorkflowSpec](#logicappworkflowspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LogicAppWorkflowStatus](#logicappworkflowstatus)

---
