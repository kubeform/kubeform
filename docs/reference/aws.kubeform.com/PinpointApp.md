---
title: PinpointApp
menu:
  docs_v0.0.1:
    identifier: pinpointapp-aws.kubeform.com
    name: PinpointApp
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## PinpointApp
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointApp` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointAppSpec](#PinpointAppSpec)***||
| `status` | ***[PinpointAppStatus](#PinpointAppStatus)***||
## PinpointAppSpec
##### (Appears on:[PinpointApp](#PinpointApp), [PinpointAppStatus](#PinpointAppStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationID` | ***string***| ***(Optional)*** |
| `campaignHook` | ***[[]PinpointAppSpecCampaignHook](#PinpointAppSpecCampaignHook)***| ***(Optional)*** |
| `limits` | ***[[]PinpointAppSpecLimits](#PinpointAppSpecLimits)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `quietTime` | ***[[]PinpointAppSpecQuietTime](#PinpointAppSpecQuietTime)***| ***(Optional)*** |
## PinpointAppSpecCampaignHook
##### (Appears on:[PinpointAppSpec](#PinpointAppSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `lambdaFunctionName` | ***string***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
| `webURL` | ***string***| ***(Optional)*** |
## PinpointAppSpecLimits
##### (Appears on:[PinpointAppSpec](#PinpointAppSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `daily` | ***int***| ***(Optional)*** |
| `maximumDuration` | ***int***| ***(Optional)*** |
| `messagesPerSecond` | ***int***| ***(Optional)*** |
| `total` | ***int***| ***(Optional)*** |
## PinpointAppSpecQuietTime
##### (Appears on:[PinpointAppSpec](#PinpointAppSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `end` | ***string***| ***(Optional)*** |
| `start` | ***string***| ***(Optional)*** |
## PinpointAppStatus
##### (Appears on:[PinpointApp](#PinpointApp))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointAppSpec](#PinpointAppSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
