---
title: NotificationHub
menu:
  docs_{{ .version }}:
    identifier: notificationhub-azurerm.kubeform.com
    name: NotificationHub
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NotificationHub
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NotificationHub` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NotificationHubSpec](#notificationhubspec)***||
| `status` | ***[NotificationHubStatus](#notificationhubstatus)***||
## NotificationHubSpec

Appears on:[NotificationHub](#notificationhub), [NotificationHubStatus](#notificationhubstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `apnsCredential` | ***[[]NotificationHubSpecApnsCredential](#notificationhubspecapnscredential)***| ***(Optional)*** |
| `gcmCredential` | ***[[]NotificationHubSpecGcmCredential](#notificationhubspecgcmcredential)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `namespaceName` | ***string***||
| `resourceGroupName` | ***string***||
## NotificationHubSpecApnsCredential

Appears on:[NotificationHubSpec](#notificationhubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `applicationMode` | ***string***||
| `bundleID` | ***string***||
| `keyID` | ***string***||
| `teamID` | ***string***||
## NotificationHubSpecGcmCredential

Appears on:[NotificationHubSpec](#notificationhubspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
## NotificationHubStatus

Appears on:[NotificationHub](#notificationhub)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NotificationHubSpec](#notificationhubspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NotificationHubStatus](#notificationhubstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `apns_credential.<index>.token` | ***string*** ||
| `gcm_credential.<index>.api_key` | ***string*** ||
