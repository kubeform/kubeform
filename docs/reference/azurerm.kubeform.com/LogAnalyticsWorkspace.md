---
title: LogAnalyticsWorkspace
menu:
  docs_v0.0.1:
    identifier: loganalyticsworkspace-azurerm.kubeform.com
    name: LogAnalyticsWorkspace
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## LogAnalyticsWorkspace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogAnalyticsWorkspace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogAnalyticsWorkspaceSpec](#LogAnalyticsWorkspaceSpec)***||
| `status` | ***[LogAnalyticsWorkspaceStatus](#LogAnalyticsWorkspaceStatus)***||
## LogAnalyticsWorkspaceSpec
##### (Appears on:[LogAnalyticsWorkspace](#LogAnalyticsWorkspace), [LogAnalyticsWorkspaceStatus](#LogAnalyticsWorkspaceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `location` | ***string***||
| `name` | ***string***||
| `portalURL` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `retentionInDays` | ***int***| ***(Optional)*** |
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `workspaceID` | ***string***| ***(Optional)*** |
## LogAnalyticsWorkspaceStatus
##### (Appears on:[LogAnalyticsWorkspace](#LogAnalyticsWorkspace))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogAnalyticsWorkspaceSpec](#LogAnalyticsWorkspaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_shared_key` | ***string*** ||
| `secondary_shared_key` | ***string*** ||
