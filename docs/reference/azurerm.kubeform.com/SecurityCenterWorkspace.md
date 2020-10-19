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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityCenterWorkspaceSpec](#securitycenterworkspacespec)***||
| `status` | ***[SecurityCenterWorkspaceStatus](#securitycenterworkspacestatus)***||
## Phase(`string` alias)

Appears on:[SecurityCenterWorkspaceStatus](#securitycenterworkspacestatus)

## SecurityCenterWorkspaceSpec

Appears on:[SecurityCenterWorkspace](#securitycenterworkspace), [SecurityCenterWorkspaceStatus](#securitycenterworkspacestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `scope` | ***string***||
| `workspaceID` | ***string***||
## SecurityCenterWorkspaceStatus

Appears on:[SecurityCenterWorkspace](#securitycenterworkspace)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityCenterWorkspaceSpec](#securitycenterworkspacespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
