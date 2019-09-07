---
title: AthenaDatabase
menu:
  docs_v0.0.1:
    identifier: athenadatabase-aws.kubeform.com
    name: AthenaDatabase
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AthenaDatabase
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AthenaDatabase` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AthenaDatabaseSpec](#AthenaDatabaseSpec)***||
| `status` | ***[AthenaDatabaseStatus](#AthenaDatabaseStatus)***||
## AthenaDatabaseSpec
##### (Appears on:[AthenaDatabase](#AthenaDatabase), [AthenaDatabaseStatus](#AthenaDatabaseStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bucket` | ***string***||
| `encryptionConfiguration` | ***[[]AthenaDatabaseSpecEncryptionConfiguration](#AthenaDatabaseSpecEncryptionConfiguration)***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
## AthenaDatabaseSpecEncryptionConfiguration
##### (Appears on:[AthenaDatabaseSpec](#AthenaDatabaseSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `encryptionOption` | ***string***||
| `kmsKey` | ***string***| ***(Optional)*** |
## AthenaDatabaseStatus
##### (Appears on:[AthenaDatabase](#AthenaDatabase))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AthenaDatabaseSpec](#AthenaDatabaseSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
