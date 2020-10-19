---
title: DedicatedHostGroup
menu:
  docs_{{ .version }}:
    identifier: dedicatedhostgroup-azurerm.kubeform.com
    name: DedicatedHostGroup
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DedicatedHostGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DedicatedHostGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DedicatedHostGroupSpec](#dedicatedhostgroupspec)***||
| `status` | ***[DedicatedHostGroupStatus](#dedicatedhostgroupstatus)***||
## DedicatedHostGroupSpec

Appears on:[DedicatedHostGroup](#dedicatedhostgroup), [DedicatedHostGroupStatus](#dedicatedhostgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `platformFaultDomainCount` | ***int64***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zones` | ***[]string***| ***(Optional)*** |
## DedicatedHostGroupStatus

Appears on:[DedicatedHostGroup](#dedicatedhostgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DedicatedHostGroupSpec](#dedicatedhostgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DedicatedHostGroupStatus](#dedicatedhostgroupstatus)

---
