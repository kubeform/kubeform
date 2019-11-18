---
title: SharedImageVersion
menu:
  docs_{{ .version }}:
    identifier: sharedimageversion-azurerm.kubeform.com
    name: SharedImageVersion
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SharedImageVersion
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SharedImageVersion` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SharedImageVersionSpec](#sharedimageversionspec)***||
| `status` | ***[SharedImageVersionStatus](#sharedimageversionstatus)***||
## Phase(`string` alias)

Appears on:[SharedImageVersionStatus](#sharedimageversionstatus)

## SharedImageVersionSpec

Appears on:[SharedImageVersion](#sharedimageversion), [SharedImageVersionStatus](#sharedimageversionstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `excludeFromLatest` | ***bool***| ***(Optional)*** |
| `galleryName` | ***string***||
| `imageName` | ***string***||
| `location` | ***string***||
| `managedImageID` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetRegion` | ***[[]SharedImageVersionSpecTargetRegion](#sharedimageversionspectargetregion)***||
## SharedImageVersionSpecTargetRegion

Appears on:[SharedImageVersionSpec](#sharedimageversionspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `regionalReplicaCount` | ***int64***||
## SharedImageVersionStatus

Appears on:[SharedImageVersion](#sharedimageversion)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SharedImageVersionSpec](#sharedimageversionspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
