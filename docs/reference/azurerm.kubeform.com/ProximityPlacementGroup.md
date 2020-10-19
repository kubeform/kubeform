---
title: ProximityPlacementGroup
menu:
  docs_{{ .version }}:
    identifier: proximityplacementgroup-azurerm.kubeform.com
    name: ProximityPlacementGroup
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ProximityPlacementGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ProximityPlacementGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProximityPlacementGroupSpec](#proximityplacementgroupspec)***||
| `status` | ***[ProximityPlacementGroupStatus](#proximityplacementgroupstatus)***||
## Phase(`string` alias)

Appears on:[ProximityPlacementGroupStatus](#proximityplacementgroupstatus)

## ProximityPlacementGroupSpec

Appears on:[ProximityPlacementGroup](#proximityplacementgroup), [ProximityPlacementGroupStatus](#proximityplacementgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## ProximityPlacementGroupStatus

Appears on:[ProximityPlacementGroup](#proximityplacementgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProximityPlacementGroupSpec](#proximityplacementgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
