---
title: Lb
menu:
  docs_{{ .version }}:
    identifier: lb-azurerm.kubeform.com
    name: Lb
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Lb
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `Lb` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbSpec](#lbspec)***||
| `status` | ***[LbStatus](#lbstatus)***||
## LbSpec

Appears on:[Lb](#lb), [LbStatus](#lbstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `frontendIPConfiguration` | ***[[]LbSpecFrontendIPConfiguration](#lbspecfrontendipconfiguration)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `privateIPAddresses` | ***[]string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## LbSpecFrontendIPConfiguration

Appears on:[LbSpec](#lbspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `inboundNATRules` | ***[]string***| ***(Optional)*** |
| `loadBalancerRules` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `outboundRules` | ***[]string***| ***(Optional)*** |
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `privateIPAddressAllocation` | ***string***| ***(Optional)*** |
| `publicIPAddressID` | ***string***| ***(Optional)*** |
| `publicIPPrefixID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***| ***(Optional)*** |
| `zones` | ***[]string***| ***(Optional)*** |
## LbStatus

Appears on:[Lb](#lb)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbSpec](#lbspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LbStatus](#lbstatus)

---
