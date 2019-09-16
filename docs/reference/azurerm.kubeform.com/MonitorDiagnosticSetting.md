---
title: MonitorDiagnosticSetting
menu:
  docs_{{ .version }}:
    identifier: monitordiagnosticsetting-azurerm.kubeform.com
    name: MonitorDiagnosticSetting
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MonitorDiagnosticSetting
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitorDiagnosticSetting` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitorDiagnosticSettingSpec](#MonitorDiagnosticSettingSpec)***||
| `status` | ***[MonitorDiagnosticSettingStatus](#MonitorDiagnosticSettingStatus)***||
## MonitorDiagnosticSettingSpec

Appears on:[MonitorDiagnosticSetting](#MonitorDiagnosticSetting), [MonitorDiagnosticSettingStatus](#MonitorDiagnosticSettingStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `eventhubAuthorizationRuleID` | ***string***| ***(Optional)*** |
| `eventhubName` | ***string***| ***(Optional)*** |
| `log` | ***[[]MonitorDiagnosticSettingSpecLog](#MonitorDiagnosticSettingSpecLog)***| ***(Optional)*** |
| `logAnalyticsWorkspaceID` | ***string***| ***(Optional)*** |
| `metric` | ***[[]MonitorDiagnosticSettingSpecMetric](#MonitorDiagnosticSettingSpecMetric)***| ***(Optional)*** |
| `name` | ***string***||
| `storageAccountID` | ***string***| ***(Optional)*** |
| `targetResourceID` | ***string***||
## MonitorDiagnosticSettingSpecLog

Appears on:[MonitorDiagnosticSettingSpec](#MonitorDiagnosticSettingSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `category` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `retentionPolicy` | ***[[]MonitorDiagnosticSettingSpecLogRetentionPolicy](#MonitorDiagnosticSettingSpecLogRetentionPolicy)***||
## MonitorDiagnosticSettingSpecLogRetentionPolicy

Appears on:[MonitorDiagnosticSettingSpecLog](#MonitorDiagnosticSettingSpecLog)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int***| ***(Optional)*** |
| `enabled` | ***bool***||
## MonitorDiagnosticSettingSpecMetric

Appears on:[MonitorDiagnosticSettingSpec](#MonitorDiagnosticSettingSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `category` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `retentionPolicy` | ***[[]MonitorDiagnosticSettingSpecMetricRetentionPolicy](#MonitorDiagnosticSettingSpecMetricRetentionPolicy)***||
## MonitorDiagnosticSettingSpecMetricRetentionPolicy

Appears on:[MonitorDiagnosticSettingSpecMetric](#MonitorDiagnosticSettingSpecMetric)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `days` | ***int***| ***(Optional)*** |
| `enabled` | ***bool***||
## MonitorDiagnosticSettingStatus

Appears on:[MonitorDiagnosticSetting](#MonitorDiagnosticSetting)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitorDiagnosticSettingSpec](#MonitorDiagnosticSettingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
