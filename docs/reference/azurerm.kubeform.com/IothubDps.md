---
title: IothubDps
menu:
  docs_{{ .version }}:
    identifier: iothubdps-azurerm.kubeform.com
    name: IothubDps
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IothubDps
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IothubDps` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IothubDpsSpec](#iothubdpsspec)***||
| `status` | ***[IothubDpsStatus](#iothubdpsstatus)***||
## IothubDpsSpec

Appears on:[IothubDps](#iothubdps), [IothubDpsStatus](#iothubdpsstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `allocationPolicy` | ***string***| ***(Optional)*** |
| `deviceProvisioningHostName` | ***string***| ***(Optional)*** |
| `IDScope` | ***string***| ***(Optional)*** |
| `linkedHub` | ***[[]IothubDpsSpecLinkedHub](#iothubdpsspeclinkedhub)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `serviceOperationsHostName` | ***string***| ***(Optional)*** |
| `sku` | ***[[]IothubDpsSpecSku](#iothubdpsspecsku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## IothubDpsSpecLinkedHub

Appears on:[IothubDpsSpec](#iothubdpsspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allocationWeight` | ***int64***| ***(Optional)*** |
| `applyAllocationPolicy` | ***bool***| ***(Optional)*** |
| `hostname` | ***string***| ***(Optional)*** |
| `location` | ***string***||
## IothubDpsSpecSku

Appears on:[IothubDpsSpec](#iothubdpsspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int64***||
| `name` | ***string***||
| `tier` | ***string***| ***(Optional)*** Deprecated|
## IothubDpsStatus

Appears on:[IothubDps](#iothubdps)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IothubDpsSpec](#iothubdpsspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IothubDpsStatus](#iothubdpsstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `linked_hub.<index>.connection_string` | ***string*** ||
