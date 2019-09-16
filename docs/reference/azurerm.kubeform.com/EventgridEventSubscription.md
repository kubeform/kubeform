---
title: EventgridEventSubscription
menu:
  docs_{{ .version }}:
    identifier: eventgrideventsubscription-azurerm.kubeform.com
    name: EventgridEventSubscription
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## EventgridEventSubscription
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `EventgridEventSubscription` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)***||
| `status` | ***[EventgridEventSubscriptionStatus](#EventgridEventSubscriptionStatus)***||
## EventgridEventSubscriptionSpec

Appears on:[EventgridEventSubscription](#EventgridEventSubscription), [EventgridEventSubscriptionStatus](#EventgridEventSubscriptionStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `eventDeliverySchema` | ***string***| ***(Optional)*** |
| `eventhubEndpoint` | ***[[]EventgridEventSubscriptionSpecEventhubEndpoint](#EventgridEventSubscriptionSpecEventhubEndpoint)***| ***(Optional)*** |
| `hybridConnectionEndpoint` | ***[[]EventgridEventSubscriptionSpecHybridConnectionEndpoint](#EventgridEventSubscriptionSpecHybridConnectionEndpoint)***| ***(Optional)*** |
| `includedEventTypes` | ***[]string***| ***(Optional)*** |
| `labels` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `retryPolicy` | ***[[]EventgridEventSubscriptionSpecRetryPolicy](#EventgridEventSubscriptionSpecRetryPolicy)***| ***(Optional)*** |
| `scope` | ***string***||
| `storageBlobDeadLetterDestination` | ***[[]EventgridEventSubscriptionSpecStorageBlobDeadLetterDestination](#EventgridEventSubscriptionSpecStorageBlobDeadLetterDestination)***| ***(Optional)*** |
| `storageQueueEndpoint` | ***[[]EventgridEventSubscriptionSpecStorageQueueEndpoint](#EventgridEventSubscriptionSpecStorageQueueEndpoint)***| ***(Optional)*** |
| `subjectFilter` | ***[[]EventgridEventSubscriptionSpecSubjectFilter](#EventgridEventSubscriptionSpecSubjectFilter)***| ***(Optional)*** |
| `topicName` | ***string***| ***(Optional)*** |
| `webhookEndpoint` | ***[[]EventgridEventSubscriptionSpecWebhookEndpoint](#EventgridEventSubscriptionSpecWebhookEndpoint)***| ***(Optional)*** |
## EventgridEventSubscriptionSpecEventhubEndpoint

Appears on:[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `eventhubID` | ***string***||
## EventgridEventSubscriptionSpecHybridConnectionEndpoint

Appears on:[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hybridConnectionID` | ***string***||
## EventgridEventSubscriptionSpecRetryPolicy

Appears on:[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `eventTimeToLive` | ***int***||
| `maxDeliveryAttempts` | ***int***||
## EventgridEventSubscriptionSpecStorageBlobDeadLetterDestination

Appears on:[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `storageAccountID` | ***string***||
| `storageBlobContainerName` | ***string***||
## EventgridEventSubscriptionSpecStorageQueueEndpoint

Appears on:[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `queueName` | ***string***||
| `storageAccountID` | ***string***||
## EventgridEventSubscriptionSpecSubjectFilter

Appears on:[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `caseSensitive` | ***bool***| ***(Optional)*** |
| `subjectBeginsWith` | ***string***| ***(Optional)*** |
| `subjectEndsWith` | ***string***| ***(Optional)*** |
## EventgridEventSubscriptionSpecWebhookEndpoint

Appears on:[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `url` | ***string***||
## EventgridEventSubscriptionStatus

Appears on:[EventgridEventSubscription](#EventgridEventSubscription)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EventgridEventSubscriptionSpec](#EventgridEventSubscriptionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
