---
title: StreamAnalyticsFunctionJavascriptUdf
menu:
  docs_{{ .version }}:
    identifier: streamanalyticsfunctionjavascriptudf-azurerm.kubeform.com
    name: StreamAnalyticsFunctionJavascriptUdf
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StreamAnalyticsFunctionJavascriptUdf
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StreamAnalyticsFunctionJavascriptUdf` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsFunctionJavascriptUdfSpec](#StreamAnalyticsFunctionJavascriptUdfSpec)***||
| `status` | ***[StreamAnalyticsFunctionJavascriptUdfStatus](#StreamAnalyticsFunctionJavascriptUdfStatus)***||
## StreamAnalyticsFunctionJavascriptUdfSpec
##### (Appears on:[StreamAnalyticsFunctionJavascriptUdf](#StreamAnalyticsFunctionJavascriptUdf), [StreamAnalyticsFunctionJavascriptUdfStatus](#StreamAnalyticsFunctionJavascriptUdfStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `input` | ***[[]StreamAnalyticsFunctionJavascriptUdfSpecInput](#StreamAnalyticsFunctionJavascriptUdfSpecInput)***||
| `name` | ***string***||
| `output` | ***[[]StreamAnalyticsFunctionJavascriptUdfSpecOutput](#StreamAnalyticsFunctionJavascriptUdfSpecOutput)***||
| `resourceGroupName` | ***string***||
| `script` | ***string***||
| `streamAnalyticsJobName` | ***string***||
## StreamAnalyticsFunctionJavascriptUdfSpecInput
##### (Appears on:[StreamAnalyticsFunctionJavascriptUdfSpec](#StreamAnalyticsFunctionJavascriptUdfSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## StreamAnalyticsFunctionJavascriptUdfSpecOutput
##### (Appears on:[StreamAnalyticsFunctionJavascriptUdfSpec](#StreamAnalyticsFunctionJavascriptUdfSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## StreamAnalyticsFunctionJavascriptUdfStatus
##### (Appears on:[StreamAnalyticsFunctionJavascriptUdf](#StreamAnalyticsFunctionJavascriptUdf))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsFunctionJavascriptUdfSpec](#StreamAnalyticsFunctionJavascriptUdfSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
