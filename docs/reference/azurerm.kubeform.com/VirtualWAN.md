---
title: VirtualWAN
menu:
  docs_{{ .version }}:
    identifier: virtualwan-azurerm.kubeform.com
    name: VirtualWAN
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VirtualWAN
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `VirtualWAN` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VirtualWANSpec](#virtualwanspec)***||
| `status` | ***[VirtualWANStatus](#virtualwanstatus)***||
## Phase(`string` alias)

Appears on:[VirtualWANStatus](#virtualwanstatus)

## VirtualWANSpec

Appears on:[VirtualWAN](#virtualwan), [VirtualWANStatus](#virtualwanstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowBranchToBranchTraffic` | ***bool***| ***(Optional)*** |
| `allowVnetToVnetTraffic` | ***bool***| ***(Optional)*** |
| `disableVPNEncryption` | ***bool***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `office365LocalBreakoutCategory` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `securityProviderName` | ***string***| ***(Optional)*** Deprecated|
| `tags` | ***map[string]string***| ***(Optional)*** |
## VirtualWANStatus

Appears on:[VirtualWAN](#virtualwan)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VirtualWANSpec](#virtualwanspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
