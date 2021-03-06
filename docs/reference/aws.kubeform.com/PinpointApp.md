---
title: PinpointApp
menu:
  docs_{{ .version }}:
    identifier: pinpointapp-aws.kubeform.com
    name: PinpointApp
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PinpointApp
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `PinpointApp` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PinpointAppSpec](#pinpointappspec)***||
| `status` | ***[PinpointAppStatus](#pinpointappstatus)***||
## Phase(`string` alias)

Appears on:[PinpointAppStatus](#pinpointappstatus)

## PinpointAppSpec

Appears on:[PinpointApp](#pinpointapp), [PinpointAppStatus](#pinpointappstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `applicationID` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `campaignHook` | ***[[]PinpointAppSpecCampaignHook](#pinpointappspeccampaignhook)***| ***(Optional)*** |
| `limits` | ***[[]PinpointAppSpecLimits](#pinpointappspeclimits)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `quietTime` | ***[[]PinpointAppSpecQuietTime](#pinpointappspecquiettime)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## PinpointAppSpecCampaignHook

Appears on:[PinpointAppSpec](#pinpointappspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `lambdaFunctionName` | ***string***| ***(Optional)*** |
| `mode` | ***string***| ***(Optional)*** |
| `webURL` | ***string***| ***(Optional)*** |
## PinpointAppSpecLimits

Appears on:[PinpointAppSpec](#pinpointappspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `daily` | ***int64***| ***(Optional)*** |
| `maximumDuration` | ***int64***| ***(Optional)*** |
| `messagesPerSecond` | ***int64***| ***(Optional)*** |
| `total` | ***int64***| ***(Optional)*** |
## PinpointAppSpecQuietTime

Appears on:[PinpointAppSpec](#pinpointappspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `end` | ***string***| ***(Optional)*** |
| `start` | ***string***| ***(Optional)*** |
## PinpointAppStatus

Appears on:[PinpointApp](#pinpointapp)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PinpointAppSpec](#pinpointappspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
