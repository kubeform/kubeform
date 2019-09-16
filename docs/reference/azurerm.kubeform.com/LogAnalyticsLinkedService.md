---
title: LogAnalyticsLinkedService
menu:
  docs_{{ .version }}:
    identifier: loganalyticslinkedservice-azurerm.kubeform.com
    name: LogAnalyticsLinkedService
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LogAnalyticsLinkedService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `LogAnalyticsLinkedService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LogAnalyticsLinkedServiceSpec](#loganalyticslinkedservicespec)***||
| `status` | ***[LogAnalyticsLinkedServiceStatus](#loganalyticslinkedservicestatus)***||
## LogAnalyticsLinkedServiceSpec

Appears on:[LogAnalyticsLinkedService](#loganalyticslinkedservice), [LogAnalyticsLinkedServiceStatus](#loganalyticslinkedservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `linkedServiceName` | ***string***| ***(Optional)*** |
| `linkedServiceProperties` | ***[[]LogAnalyticsLinkedServiceSpecLinkedServiceProperties](#loganalyticslinkedservicespeclinkedserviceproperties)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `resourceID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `workspaceName` | ***string***||
## LogAnalyticsLinkedServiceSpecLinkedServiceProperties

Appears on:[LogAnalyticsLinkedServiceSpec](#loganalyticslinkedservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `resourceID` | ***string***||
## LogAnalyticsLinkedServiceStatus

Appears on:[LogAnalyticsLinkedService](#loganalyticslinkedservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LogAnalyticsLinkedServiceSpec](#loganalyticslinkedservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
