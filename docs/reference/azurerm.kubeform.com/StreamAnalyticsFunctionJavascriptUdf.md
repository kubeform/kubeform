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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsFunctionJavascriptUdfSpec](#streamanalyticsfunctionjavascriptudfspec)***||
| `status` | ***[StreamAnalyticsFunctionJavascriptUdfStatus](#streamanalyticsfunctionjavascriptudfstatus)***||
## Phase(`string` alias)

Appears on:[StreamAnalyticsFunctionJavascriptUdfStatus](#streamanalyticsfunctionjavascriptudfstatus)

## StreamAnalyticsFunctionJavascriptUdfSpec

Appears on:[StreamAnalyticsFunctionJavascriptUdf](#streamanalyticsfunctionjavascriptudf), [StreamAnalyticsFunctionJavascriptUdfStatus](#streamanalyticsfunctionjavascriptudfstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `input` | ***[[]StreamAnalyticsFunctionJavascriptUdfSpecInput](#streamanalyticsfunctionjavascriptudfspecinput)***||
| `name` | ***string***||
| `output` | ***[[]StreamAnalyticsFunctionJavascriptUdfSpecOutput](#streamanalyticsfunctionjavascriptudfspecoutput)***||
| `resourceGroupName` | ***string***||
| `script` | ***string***||
| `streamAnalyticsJobName` | ***string***||
## StreamAnalyticsFunctionJavascriptUdfSpecInput

Appears on:[StreamAnalyticsFunctionJavascriptUdfSpec](#streamanalyticsfunctionjavascriptudfspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## StreamAnalyticsFunctionJavascriptUdfSpecOutput

Appears on:[StreamAnalyticsFunctionJavascriptUdfSpec](#streamanalyticsfunctionjavascriptudfspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## StreamAnalyticsFunctionJavascriptUdfStatus

Appears on:[StreamAnalyticsFunctionJavascriptUdf](#streamanalyticsfunctionjavascriptudf)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsFunctionJavascriptUdfSpec](#streamanalyticsfunctionjavascriptudfspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
