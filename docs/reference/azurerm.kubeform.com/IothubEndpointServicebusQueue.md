---
title: IothubEndpointServicebusQueue
menu:
  docs_{{ .version }}:
    identifier: iothubendpointservicebusqueue-azurerm.kubeform.com
    name: IothubEndpointServicebusQueue
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IothubEndpointServicebusQueue
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IothubEndpointServicebusQueue` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IothubEndpointServicebusQueueSpec](#iothubendpointservicebusqueuespec)***||
| `status` | ***[IothubEndpointServicebusQueueStatus](#iothubendpointservicebusqueuestatus)***||
## IothubEndpointServicebusQueueSpec

Appears on:[IothubEndpointServicebusQueue](#iothubendpointservicebusqueue), [IothubEndpointServicebusQueueStatus](#iothubendpointservicebusqueuestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `iothubName` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## IothubEndpointServicebusQueueStatus

Appears on:[IothubEndpointServicebusQueue](#iothubendpointservicebusqueue)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IothubEndpointServicebusQueueSpec](#iothubendpointservicebusqueuespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IothubEndpointServicebusQueueStatus](#iothubendpointservicebusqueuestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `connection_string` | ***string*** ||
