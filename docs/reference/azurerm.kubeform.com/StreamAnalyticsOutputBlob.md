---
title: StreamAnalyticsOutputBlob
menu:
  docs_{{ .version }}:
    identifier: streamanalyticsoutputblob-azurerm.kubeform.com
    name: StreamAnalyticsOutputBlob
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StreamAnalyticsOutputBlob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StreamAnalyticsOutputBlob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsOutputBlobSpec](#streamanalyticsoutputblobspec)***||
| `status` | ***[StreamAnalyticsOutputBlobStatus](#streamanalyticsoutputblobstatus)***||
## Phase(`string` alias)

Appears on:[StreamAnalyticsOutputBlobStatus](#streamanalyticsoutputblobstatus)

## StreamAnalyticsOutputBlobSpec

Appears on:[StreamAnalyticsOutputBlob](#streamanalyticsoutputblob), [StreamAnalyticsOutputBlobStatus](#streamanalyticsoutputblobstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `dateFormat` | ***string***||
| `name` | ***string***||
| `pathPattern` | ***string***||
| `resourceGroupName` | ***string***||
| `serialization` | ***[[]StreamAnalyticsOutputBlobSpecSerialization](#streamanalyticsoutputblobspecserialization)***||
| `storageAccountName` | ***string***||
| `storageContainerName` | ***string***||
| `streamAnalyticsJobName` | ***string***||
| `timeFormat` | ***string***||
## StreamAnalyticsOutputBlobSpecSerialization

Appears on:[StreamAnalyticsOutputBlobSpec](#streamanalyticsoutputblobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `encoding` | ***string***| ***(Optional)*** |
| `fieldDelimiter` | ***string***| ***(Optional)*** |
| `format` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## StreamAnalyticsOutputBlobStatus

Appears on:[StreamAnalyticsOutputBlob](#streamanalyticsoutputblob)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsOutputBlobSpec](#streamanalyticsoutputblobspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `storage_account_key` | ***string*** ||
