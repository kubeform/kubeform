---
title: IotDps
menu:
  docs_{{ .version }}:
    identifier: iotdps-azurerm.kubeform.com
    name: IotDps
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IotDps
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IotDps` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IotDpsSpec](#iotdpsspec)***||
| `status` | ***[IotDpsStatus](#iotdpsstatus)***||
## IotDpsSpec

Appears on:[IotDps](#iotdps), [IotDpsStatus](#iotdpsstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***[[]IotDpsSpecSku](#iotdpsspecsku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## IotDpsSpecSku

Appears on:[IotDpsSpec](#iotdpsspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `name` | ***string***||
| `tier` | ***string***||
## IotDpsStatus

Appears on:[IotDps](#iotdps)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IotDpsSpec](#iotdpsspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---