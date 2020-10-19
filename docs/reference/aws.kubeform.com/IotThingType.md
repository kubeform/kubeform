---
title: IotThingType
menu:
  docs_{{ .version }}:
    identifier: iotthingtype-aws.kubeform.com
    name: IotThingType
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IotThingType
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IotThingType` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IotThingTypeSpec](#iotthingtypespec)***||
| `status` | ***[IotThingTypeStatus](#iotthingtypestatus)***||
## IotThingTypeSpec

Appears on:[IotThingType](#iotthingtype), [IotThingTypeStatus](#iotthingtypestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `deprecated` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `properties` | ***[[]IotThingTypeSpecProperties](#iotthingtypespecproperties)***| ***(Optional)*** |
## IotThingTypeSpecProperties

Appears on:[IotThingTypeSpec](#iotthingtypespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `searchableAttributes` | ***[]string***| ***(Optional)*** |
## IotThingTypeStatus

Appears on:[IotThingType](#iotthingtype)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IotThingTypeSpec](#iotthingtypespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IotThingTypeStatus](#iotthingtypestatus)

---
