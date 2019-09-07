---
title: LogAnalyticsWorkspaceLinkedService
menu:
  docs_v0.0.1:
    identifier: loganalyticsworkspacelinkedservice-azurerm.kubeform.com
    name: LogAnalyticsWorkspaceLinkedService
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## LogAnalyticsWorkspaceLinkedService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogAnalyticsWorkspaceLinkedService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogAnalyticsWorkspaceLinkedServiceSpec](#LogAnalyticsWorkspaceLinkedServiceSpec)***||
| `status` | ***[LogAnalyticsWorkspaceLinkedServiceStatus](#LogAnalyticsWorkspaceLinkedServiceStatus)***||
## LogAnalyticsWorkspaceLinkedServiceSpec
##### (Appears on:[LogAnalyticsWorkspaceLinkedService](#LogAnalyticsWorkspaceLinkedService), [LogAnalyticsWorkspaceLinkedServiceStatus](#LogAnalyticsWorkspaceLinkedServiceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `linkedServiceName` | ***string***| ***(Optional)*** |
| `linkedServiceProperties` | ***[[]LogAnalyticsWorkspaceLinkedServiceSpecLinkedServiceProperties](#LogAnalyticsWorkspaceLinkedServiceSpecLinkedServiceProperties)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `resourceID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `workspaceName` | ***string***||
## LogAnalyticsWorkspaceLinkedServiceSpecLinkedServiceProperties
##### (Appears on:[LogAnalyticsWorkspaceLinkedServiceSpec](#LogAnalyticsWorkspaceLinkedServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceID` | ***string***||
## LogAnalyticsWorkspaceLinkedServiceStatus
##### (Appears on:[LogAnalyticsWorkspaceLinkedService](#LogAnalyticsWorkspaceLinkedService))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogAnalyticsWorkspaceLinkedServiceSpec](#LogAnalyticsWorkspaceLinkedServiceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
