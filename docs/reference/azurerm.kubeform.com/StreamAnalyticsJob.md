---
title: StreamAnalyticsJob
menu:
  docs_{{ .version }}:
    identifier: streamanalyticsjob-azurerm.kubeform.com
    name: StreamAnalyticsJob
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StreamAnalyticsJob
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StreamAnalyticsJob` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsJobSpec](#StreamAnalyticsJobSpec)***||
| `status` | ***[StreamAnalyticsJobStatus](#StreamAnalyticsJobStatus)***||
## StreamAnalyticsJobSpec

Appears on:[StreamAnalyticsJob](#StreamAnalyticsJob), [StreamAnalyticsJobStatus](#StreamAnalyticsJobStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `compatibilityLevel` | ***string***||
| `dataLocale` | ***string***||
| `eventsLateArrivalMaxDelayInSeconds` | ***int***||
| `eventsOutOfOrderMaxDelayInSeconds` | ***int***||
| `eventsOutOfOrderPolicy` | ***string***||
| `jobID` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `outputErrorPolicy` | ***string***||
| `resourceGroupName` | ***string***||
| `streamingUnits` | ***int***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `transformationQuery` | ***string***||
## StreamAnalyticsJobStatus

Appears on:[StreamAnalyticsJob](#StreamAnalyticsJob)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsJobSpec](#StreamAnalyticsJobSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
