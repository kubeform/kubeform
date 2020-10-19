---
title: ApplicationInsightsWebTest
menu:
  docs_{{ .version }}:
    identifier: applicationinsightswebtest-azurerm.kubeform.com
    name: ApplicationInsightsWebTest
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApplicationInsightsWebTest
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApplicationInsightsWebTest` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApplicationInsightsWebTestSpec](#applicationinsightswebtestspec)***||
| `status` | ***[ApplicationInsightsWebTestStatus](#applicationinsightswebteststatus)***||
## ApplicationInsightsWebTestSpec

Appears on:[ApplicationInsightsWebTest](#applicationinsightswebtest), [ApplicationInsightsWebTestStatus](#applicationinsightswebteststatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationInsightsID` | ***string***||
| `configuration` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `frequency` | ***int64***| ***(Optional)*** |
| `geoLocations` | ***[]string***||
| `kind` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `retryEnabled` | ***bool***| ***(Optional)*** |
| `syntheticMonitorID` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `timeout` | ***int64***| ***(Optional)*** |
## ApplicationInsightsWebTestStatus

Appears on:[ApplicationInsightsWebTest](#applicationinsightswebtest)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApplicationInsightsWebTestSpec](#applicationinsightswebtestspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApplicationInsightsWebTestStatus](#applicationinsightswebteststatus)

---
