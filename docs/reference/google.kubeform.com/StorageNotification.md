---
title: StorageNotification
menu:
  docs_{{ .version }}:
    identifier: storagenotification-google.kubeform.com
    name: StorageNotification
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StorageNotification
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `StorageNotification` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StorageNotificationSpec](#StorageNotificationSpec)***||
| `status` | ***[StorageNotificationStatus](#StorageNotificationStatus)***||
## StorageNotificationSpec

Appears on:[StorageNotification](#StorageNotification), [StorageNotificationStatus](#StorageNotificationStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `customAttributes` | ***map[string]string***| ***(Optional)*** |
| `eventTypes` | ***[]string***| ***(Optional)*** |
| `objectNamePrefix` | ***string***| ***(Optional)*** |
| `payloadFormat` | ***string***||
| `selfLink` | ***string***| ***(Optional)*** |
| `topic` | ***string***||
## StorageNotificationStatus

Appears on:[StorageNotification](#StorageNotification)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StorageNotificationSpec](#StorageNotificationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
