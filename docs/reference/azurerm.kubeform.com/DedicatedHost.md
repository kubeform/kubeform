---
title: DedicatedHost
menu:
  docs_{{ .version }}:
    identifier: dedicatedhost-azurerm.kubeform.com
    name: DedicatedHost
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DedicatedHost
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DedicatedHost` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DedicatedHostSpec](#dedicatedhostspec)***||
| `status` | ***[DedicatedHostStatus](#dedicatedhoststatus)***||
## DedicatedHostSpec

Appears on:[DedicatedHost](#dedicatedhost), [DedicatedHostStatus](#dedicatedhoststatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `autoReplaceOnFailure` | ***bool***| ***(Optional)*** |
| `dedicatedHostGroupID` | ***string***||
| `licenseType` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `platformFaultDomain` | ***int64***||
| `skuName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## DedicatedHostStatus

Appears on:[DedicatedHost](#dedicatedhost)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DedicatedHostSpec](#dedicatedhostspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DedicatedHostStatus](#dedicatedhoststatus)

---
