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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevicefarmProjectSpec](#DevicefarmProjectSpec)***||
| `status` | ***[DevicefarmProjectStatus](#DevicefarmProjectStatus)***||
## DevicefarmProjectSpec
##### (Appears on:[DevicefarmProject](#DevicefarmProject), [DevicefarmProjectStatus](#DevicefarmProjectStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## DevicefarmProjectStatus
##### (Appears on:[DevicefarmProject](#DevicefarmProject))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevicefarmProjectSpec](#DevicefarmProjectSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
