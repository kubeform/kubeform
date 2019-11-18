---
title: AvailabilitySet
menu:
  docs_{{ .version }}:
    identifier: availabilityset-azurerm.kubeform.com
    name: AvailabilitySet
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AvailabilitySet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AvailabilitySet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AvailabilitySetSpec](#availabilitysetspec)***||
| `status` | ***[AvailabilitySetStatus](#availabilitysetstatus)***||
## AvailabilitySetSpec

Appears on:[AvailabilitySet](#availabilityset), [AvailabilitySetStatus](#availabilitysetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `managed` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `platformFaultDomainCount` | ***int64***| ***(Optional)*** |
| `platformUpdateDomainCount` | ***int64***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## AvailabilitySetStatus

Appears on:[AvailabilitySet](#availabilityset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AvailabilitySetSpec](#availabilitysetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AvailabilitySetStatus](#availabilitysetstatus)

---
