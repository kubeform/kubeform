---
title: StreamAnalyticsStreamInputIothub
menu:
  docs_{{ .version }}:
    identifier: streamanalyticsstreaminputiothub-azurerm.kubeform.com
    name: StreamAnalyticsStreamInputIothub
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StreamAnalyticsStreamInputIothub
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StreamAnalyticsStreamInputIothub` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsStreamInputIothubSpec](#StreamAnalyticsStreamInputIothubSpec)***||
| `status` | ***[StreamAnalyticsStreamInputIothubStatus](#StreamAnalyticsStreamInputIothubStatus)***||
## StreamAnalyticsStreamInputIothubSpec

Appears on:[StreamAnalyticsStreamInputIothub](#StreamAnalyticsStreamInputIothub), [StreamAnalyticsStreamInputIothubStatus](#StreamAnalyticsStreamInputIothubStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `endpoint` | ***string***||
| `eventhubConsumerGroupName` | ***string***||
| `iothubNamespace` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serialization` | ***[[]StreamAnalyticsStreamInputIothubSpecSerialization](#StreamAnalyticsStreamInputIothubSpecSerialization)***||
| `sharedAccessPolicyName` | ***string***||
| `streamAnalyticsJobName` | ***string***||
## StreamAnalyticsStreamInputIothubSpecSerialization

Appears on:[StreamAnalyticsStreamInputIothubSpec](#StreamAnalyticsStreamInputIothubSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `encoding` | ***string***| ***(Optional)*** |
| `fieldDelimiter` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## StreamAnalyticsStreamInputIothubStatus

Appears on:[StreamAnalyticsStreamInputIothub](#StreamAnalyticsStreamInputIothub)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsStreamInputIothubSpec](#StreamAnalyticsStreamInputIothubSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `shared_access_policy_key` | ***string*** ||
