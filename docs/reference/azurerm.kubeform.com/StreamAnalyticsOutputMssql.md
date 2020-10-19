---
title: StreamAnalyticsOutputMssql
menu:
  docs_{{ .version }}:
    identifier: streamanalyticsoutputmssql-azurerm.kubeform.com
    name: StreamAnalyticsOutputMssql
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StreamAnalyticsOutputMssql
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StreamAnalyticsOutputMssql` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsOutputMssqlSpec](#streamanalyticsoutputmssqlspec)***||
| `status` | ***[StreamAnalyticsOutputMssqlStatus](#streamanalyticsoutputmssqlstatus)***||
## Phase(`string` alias)

Appears on:[StreamAnalyticsOutputMssqlStatus](#streamanalyticsoutputmssqlstatus)

## StreamAnalyticsOutputMssqlSpec

Appears on:[StreamAnalyticsOutputMssql](#streamanalyticsoutputmssql), [StreamAnalyticsOutputMssqlStatus](#streamanalyticsoutputmssqlstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `database` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `server` | ***string***||
| `streamAnalyticsJobName` | ***string***||
| `table` | ***string***||
| `user` | ***string***||
## StreamAnalyticsOutputMssqlStatus

Appears on:[StreamAnalyticsOutputMssql](#streamanalyticsoutputmssql)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsOutputMssqlSpec](#streamanalyticsoutputmssqlspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
