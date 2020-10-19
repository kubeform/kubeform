---
title: LogAnalyticsWorkspace
menu:
  docs_{{ .version }}:
    identifier: loganalyticsworkspace-azurerm.kubeform.com
    name: LogAnalyticsWorkspace
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LogAnalyticsWorkspace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogAnalyticsWorkspace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogAnalyticsWorkspaceSpec](#loganalyticsworkspacespec)***||
| `status` | ***[LogAnalyticsWorkspaceStatus](#loganalyticsworkspacestatus)***||
## LogAnalyticsWorkspaceSpec

Appears on:[LogAnalyticsWorkspace](#loganalyticsworkspace), [LogAnalyticsWorkspaceStatus](#loganalyticsworkspacestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `location` | ***string***||
| `name` | ***string***||
| `portalURL` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `retentionInDays` | ***int64***| ***(Optional)*** |
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `workspaceID` | ***string***| ***(Optional)*** |
## LogAnalyticsWorkspaceStatus

Appears on:[LogAnalyticsWorkspace](#loganalyticsworkspace)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogAnalyticsWorkspaceSpec](#loganalyticsworkspacespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LogAnalyticsWorkspaceStatus](#loganalyticsworkspacestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_shared_key` | ***string*** ||
| `secondary_shared_key` | ***string*** ||
