---
title: IotThing
menu:
  docs_v0.0.1:
    identifier: iotthing-aws.kubeform.com
    name: IotThing
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IotThing
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IotThing` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IotThingSpec](#IotThingSpec)***||
| `status` | ***[IotThingStatus](#IotThingStatus)***||
## IotThingSpec
##### (Appears on:[IotThing](#IotThing), [IotThingStatus](#IotThingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `attributes` | ***map[string]string***| ***(Optional)*** |
| `defaultClientID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `thingTypeName` | ***string***| ***(Optional)*** |
| `version` | ***int***| ***(Optional)*** |
## IotThingStatus
##### (Appears on:[IotThing](#IotThing))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IotThingSpec](#IotThingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
