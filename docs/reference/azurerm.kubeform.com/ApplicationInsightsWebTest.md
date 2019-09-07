---
title: ApplicationInsightsWebTest
menu:
  docs_v0.0.1:
    identifier: applicationinsightswebtest-azurerm.kubeform.com
    name: ApplicationInsightsWebTest
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApplicationInsightsWebTest
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApplicationInsightsWebTest` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApplicationInsightsWebTestSpec](#ApplicationInsightsWebTestSpec)***||
| `status` | ***[ApplicationInsightsWebTestStatus](#ApplicationInsightsWebTestStatus)***||
## ApplicationInsightsWebTestSpec
##### (Appears on:[ApplicationInsightsWebTest](#ApplicationInsightsWebTest), [ApplicationInsightsWebTestStatus](#ApplicationInsightsWebTestStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationInsightsID` | ***string***||
| `configuration` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `frequency` | ***int***| ***(Optional)*** |
| `geoLocations` | ***[]string***||
| `kind` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `retryEnabled` | ***bool***| ***(Optional)*** |
| `syntheticMonitorID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `timeout` | ***int***| ***(Optional)*** |
## ApplicationInsightsWebTestStatus
##### (Appears on:[ApplicationInsightsWebTest](#ApplicationInsightsWebTest))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApplicationInsightsWebTestSpec](#ApplicationInsightsWebTestSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
