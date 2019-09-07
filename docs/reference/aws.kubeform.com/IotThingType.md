---
title: IotThingType
menu:
  docs_v0.0.1:
    identifier: iotthingtype-aws.kubeform.com
    name: IotThingType
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IotThingType
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IotThingType` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IotThingTypeSpec](#IotThingTypeSpec)***||
| `status` | ***[IotThingTypeStatus](#IotThingTypeStatus)***||
## IotThingTypeSpec
##### (Appears on:[IotThingType](#IotThingType), [IotThingTypeStatus](#IotThingTypeStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `deprecated` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `properties` | ***[[]IotThingTypeSpecProperties](#IotThingTypeSpecProperties)***| ***(Optional)*** |
## IotThingTypeSpecProperties
##### (Appears on:[IotThingTypeSpec](#IotThingTypeSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `searchableAttributes` | ***[]string***| ***(Optional)*** |
## IotThingTypeStatus
##### (Appears on:[IotThingType](#IotThingType))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IotThingTypeSpec](#IotThingTypeSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
