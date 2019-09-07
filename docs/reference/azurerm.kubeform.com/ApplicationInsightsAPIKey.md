---
title: ApplicationInsightsAPIKey
menu:
  docs_v0.0.1:
    identifier: applicationinsightsapikey-azurerm.kubeform.com
    name: ApplicationInsightsAPIKey
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApplicationInsightsAPIKey
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApplicationInsightsAPIKey` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApplicationInsightsAPIKeySpec](#ApplicationInsightsAPIKeySpec)***||
| `status` | ***[ApplicationInsightsAPIKeyStatus](#ApplicationInsightsAPIKeyStatus)***||
## ApplicationInsightsAPIKeySpec
##### (Appears on:[ApplicationInsightsAPIKey](#ApplicationInsightsAPIKey), [ApplicationInsightsAPIKeyStatus](#ApplicationInsightsAPIKeyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `applicationInsightsID` | ***string***||
| `name` | ***string***||
| `readPermissions` | ***[]string***| ***(Optional)*** |
| `writePermissions` | ***[]string***| ***(Optional)*** |
## ApplicationInsightsAPIKeyStatus
##### (Appears on:[ApplicationInsightsAPIKey](#ApplicationInsightsAPIKey))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApplicationInsightsAPIKeySpec](#ApplicationInsightsAPIKeySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `api_key` | ***string*** ||
