---
title: SharedImageGallery
menu:
  docs_{{ .version }}:
    identifier: sharedimagegallery-azurerm.kubeform.com
    name: SharedImageGallery
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SharedImageGallery
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SharedImageGallery` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SharedImageGallerySpec](#sharedimagegalleryspec)***||
| `status` | ***[SharedImageGalleryStatus](#sharedimagegallerystatus)***||
## Phase(`string` alias)

Appears on:[SharedImageGalleryStatus](#sharedimagegallerystatus)

## SharedImageGallerySpec

Appears on:[SharedImageGallery](#sharedimagegallery), [SharedImageGalleryStatus](#sharedimagegallerystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `uniqueName` | ***string***| ***(Optional)*** |
## SharedImageGalleryStatus

Appears on:[SharedImageGallery](#sharedimagegallery)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SharedImageGallerySpec](#sharedimagegalleryspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
