---
title: DevicefarmProject
menu:
  docs_{{ .version }}:
    identifier: devicefarmproject-aws.kubeform.com
    name: DevicefarmProject
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DevicefarmProject
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DevicefarmProject` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevicefarmProjectSpec](#devicefarmprojectspec)***||
| `status` | ***[DevicefarmProjectStatus](#devicefarmprojectstatus)***||
## DevicefarmProjectSpec

Appears on:[DevicefarmProject](#devicefarmproject), [DevicefarmProjectStatus](#devicefarmprojectstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## DevicefarmProjectStatus

Appears on:[DevicefarmProject](#devicefarmproject)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevicefarmProjectSpec](#devicefarmprojectspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DevicefarmProjectStatus](#devicefarmprojectstatus)

---
