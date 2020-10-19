---
title: StreamAnalyticsOutputServicebusTopic
menu:
  docs_{{ .version }}:
    identifier: streamanalyticsoutputservicebustopic-azurerm.kubeform.com
    name: StreamAnalyticsOutputServicebusTopic
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## StreamAnalyticsOutputServicebusTopic
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `StreamAnalyticsOutputServicebusTopic` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[StreamAnalyticsOutputServicebusTopicSpec](#streamanalyticsoutputservicebustopicspec)***||
| `status` | ***[StreamAnalyticsOutputServicebusTopicStatus](#streamanalyticsoutputservicebustopicstatus)***||
## Phase(`string` alias)

Appears on:[StreamAnalyticsOutputServicebusTopicStatus](#streamanalyticsoutputservicebustopicstatus)

## StreamAnalyticsOutputServicebusTopicSpec

Appears on:[StreamAnalyticsOutputServicebusTopic](#streamanalyticsoutputservicebustopic), [StreamAnalyticsOutputServicebusTopicStatus](#streamanalyticsoutputservicebustopicstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serialization` | ***[[]StreamAnalyticsOutputServicebusTopicSpecSerialization](#streamanalyticsoutputservicebustopicspecserialization)***||
| `servicebusNamespace` | ***string***||
| `sharedAccessPolicyName` | ***string***||
| `streamAnalyticsJobName` | ***string***||
| `topicName` | ***string***||
## StreamAnalyticsOutputServicebusTopicSpecSerialization

Appears on:[StreamAnalyticsOutputServicebusTopicSpec](#streamanalyticsoutputservicebustopicspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `encoding` | ***string***| ***(Optional)*** |
| `fieldDelimiter` | ***string***| ***(Optional)*** |
| `format` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## StreamAnalyticsOutputServicebusTopicStatus

Appears on:[StreamAnalyticsOutputServicebusTopic](#streamanalyticsoutputservicebustopic)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[StreamAnalyticsOutputServicebusTopicSpec](#streamanalyticsoutputservicebustopicspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `shared_access_policy_key` | ***string*** ||
