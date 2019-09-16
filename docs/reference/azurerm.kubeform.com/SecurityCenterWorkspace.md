---
title: SecurityCenterWorkspace
menu:
  docs_{{ .version }}:
    identifier: securitycenterworkspace-azurerm.kubeform.com
    name: SecurityCenterWorkspace
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SecurityCenterWorkspace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SecurityCenterWorkspace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityCenterWorkspaceSpec](#SecurityCenterWorkspaceSpec)***||
| `status` | ***[SecurityCenterWorkspaceStatus](#SecurityCenterWorkspaceStatus)***||
## SecurityCenterWorkspaceSpec

Appears on:[SecurityCenterWorkspace](#SecurityCenterWorkspace), [SecurityCenterWorkspaceStatus](#SecurityCenterWorkspaceStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `scope` | ***string***||
| `workspaceID` | ***string***||
## SecurityCenterWorkspaceStatus

Appears on:[SecurityCenterWorkspace](#SecurityCenterWorkspace)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityCenterWorkspaceSpec](#SecurityCenterWorkspaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
