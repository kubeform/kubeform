---
title: StreamAnalyticsReferenceInputBlob
menu:
  docs_{{ .version }}:
    identifier: streamanalyticsreferenceinputblob-azurerm.kubeform.com
    name: StreamAnalyticsReferenceInputBlob
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StreamAnalyticsReferenceInputBlob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StreamAnalyticsReferenceInputBlob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsReferenceInputBlobSpec](#streamanalyticsreferenceinputblobspec)***||
| `status` | ***[StreamAnalyticsReferenceInputBlobStatus](#streamanalyticsreferenceinputblobstatus)***||
## Phase(`string` alias)

Appears on:[StreamAnalyticsReferenceInputBlobStatus](#streamanalyticsreferenceinputblobstatus)

## StreamAnalyticsReferenceInputBlobSpec

Appears on:[StreamAnalyticsReferenceInputBlob](#streamanalyticsreferenceinputblob), [StreamAnalyticsReferenceInputBlobStatus](#streamanalyticsreferenceinputblobstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `dateFormat` | ***string***||
| `name` | ***string***||
| `pathPattern` | ***string***||
| `resourceGroupName` | ***string***||
| `serialization` | ***[[]StreamAnalyticsReferenceInputBlobSpecSerialization](#streamanalyticsreferenceinputblobspecserialization)***||
| `storageAccountName` | ***string***||
| `storageContainerName` | ***string***||
| `streamAnalyticsJobName` | ***string***||
| `timeFormat` | ***string***||
## StreamAnalyticsReferenceInputBlobSpecSerialization

Appears on:[StreamAnalyticsReferenceInputBlobSpec](#streamanalyticsreferenceinputblobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `encoding` | ***string***| ***(Optional)*** |
| `fieldDelimiter` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## StreamAnalyticsReferenceInputBlobStatus

Appears on:[StreamAnalyticsReferenceInputBlob](#streamanalyticsreferenceinputblob)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsReferenceInputBlobSpec](#streamanalyticsreferenceinputblobspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `storage_account_key` | ***string*** ||
