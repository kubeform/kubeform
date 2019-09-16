---
title: StreamAnalyticsStreamInputBlob
menu:
  docs_{{ .version }}:
    identifier: streamanalyticsstreaminputblob-azurerm.kubeform.com
    name: StreamAnalyticsStreamInputBlob
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StreamAnalyticsStreamInputBlob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StreamAnalyticsStreamInputBlob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsStreamInputBlobSpec](#streamanalyticsstreaminputblobspec)***||
| `status` | ***[StreamAnalyticsStreamInputBlobStatus](#streamanalyticsstreaminputblobstatus)***||
## StreamAnalyticsStreamInputBlobSpec

Appears on:[StreamAnalyticsStreamInputBlob](#streamanalyticsstreaminputblob), [StreamAnalyticsStreamInputBlobStatus](#streamanalyticsstreaminputblobstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `dateFormat` | ***string***||
| `name` | ***string***||
| `pathPattern` | ***string***||
| `resourceGroupName` | ***string***||
| `serialization` | ***[[]StreamAnalyticsStreamInputBlobSpecSerialization](#streamanalyticsstreaminputblobspecserialization)***||
| `storageAccountName` | ***string***||
| `storageContainerName` | ***string***||
| `streamAnalyticsJobName` | ***string***||
| `timeFormat` | ***string***||
## StreamAnalyticsStreamInputBlobSpecSerialization

Appears on:[StreamAnalyticsStreamInputBlobSpec](#streamanalyticsstreaminputblobspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `encoding` | ***string***| ***(Optional)*** |
| `fieldDelimiter` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## StreamAnalyticsStreamInputBlobStatus

Appears on:[StreamAnalyticsStreamInputBlob](#streamanalyticsstreaminputblob)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsStreamInputBlobSpec](#streamanalyticsstreaminputblobspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `storage_account_key` | ***string*** ||
