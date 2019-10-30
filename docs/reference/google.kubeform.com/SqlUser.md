---
title: SqlUser
menu:
  docs_{{ .version }}:
    identifier: sqluser-google.kubeform.com
    name: SqlUser
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SqlUser
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `SqlUser` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SqlUserSpec](#sqluserspec)***||
| `status` | ***[SqlUserStatus](#sqluserstatus)***||
## Phase(`string` alias)

Appears on:[SqlUserStatus](#sqluserstatus)

## SqlUserSpec

Appears on:[SqlUser](#sqluser), [SqlUserStatus](#sqluserstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `host` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## SqlUserStatus

Appears on:[SqlUser](#sqluser)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SqlUserSpec](#sqluserspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
