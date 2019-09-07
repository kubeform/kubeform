---
title: LogAnalyticsSolution
menu:
  docs_v0.0.1:
    identifier: loganalyticssolution-azurerm.kubeform.com
    name: LogAnalyticsSolution
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## LogAnalyticsSolution
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogAnalyticsSolution` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogAnalyticsSolutionSpec](#LogAnalyticsSolutionSpec)***||
| `status` | ***[LogAnalyticsSolutionStatus](#LogAnalyticsSolutionStatus)***||
## LogAnalyticsSolutionSpec
##### (Appears on:[LogAnalyticsSolution](#LogAnalyticsSolution), [LogAnalyticsSolutionStatus](#LogAnalyticsSolutionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `plan` | ***[[]LogAnalyticsSolutionSpecPlan](#LogAnalyticsSolutionSpecPlan)***||
| `resourceGroupName` | ***string***||
| `solutionName` | ***string***||
| `workspaceName` | ***string***||
| `workspaceResourceID` | ***string***||
## LogAnalyticsSolutionSpecPlan
##### (Appears on:[LogAnalyticsSolutionSpec](#LogAnalyticsSolutionSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `product` | ***string***||
| `promotionCode` | ***string***| ***(Optional)*** |
| `publisher` | ***string***||
## LogAnalyticsSolutionStatus
##### (Appears on:[LogAnalyticsSolution](#LogAnalyticsSolution))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogAnalyticsSolutionSpec](#LogAnalyticsSolutionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
