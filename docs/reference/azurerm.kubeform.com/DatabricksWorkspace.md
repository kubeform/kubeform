---
title: DatabricksWorkspace
menu:
  docs_v0.0.1:
    identifier: databricksworkspace-azurerm.kubeform.com
    name: DatabricksWorkspace
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DatabricksWorkspace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DatabricksWorkspace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatabricksWorkspaceSpec](#DatabricksWorkspaceSpec)***||
| `status` | ***[DatabricksWorkspaceStatus](#DatabricksWorkspaceStatus)***||
## DatabricksWorkspaceSpec
##### (Appears on:[DatabricksWorkspace](#DatabricksWorkspace), [DatabricksWorkspaceStatus](#DatabricksWorkspaceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `managedResourceGroupID` | ***string***| ***(Optional)*** |
| `managedResourceGroupName` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## DatabricksWorkspaceStatus
##### (Appears on:[DatabricksWorkspace](#DatabricksWorkspace))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatabricksWorkspaceSpec](#DatabricksWorkspaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
