---
title: IothubConsumerGroup
menu:
  docs_{{ .version }}:
    identifier: iothubconsumergroup-azurerm.kubeform.com
    name: IothubConsumerGroup
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IothubConsumerGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IothubConsumerGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IothubConsumerGroupSpec](#iothubconsumergroupspec)***||
| `status` | ***[IothubConsumerGroupStatus](#iothubconsumergroupstatus)***||
## IothubConsumerGroupSpec

Appears on:[IothubConsumerGroup](#iothubconsumergroup), [IothubConsumerGroupStatus](#iothubconsumergroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `eventhubEndpointName` | ***string***||
| `iothubName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## IothubConsumerGroupStatus

Appears on:[IothubConsumerGroup](#iothubconsumergroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IothubConsumerGroupSpec](#iothubconsumergroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
