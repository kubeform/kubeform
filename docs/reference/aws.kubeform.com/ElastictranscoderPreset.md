---
title: ElastictranscoderPreset
menu:
  docs_{{ .version }}:
    identifier: elastictranscoderpreset-aws.kubeform.com
    name: ElastictranscoderPreset
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ElastictranscoderPreset
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElastictranscoderPreset` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElastictranscoderPresetSpec](#elastictranscoderpresetspec)***||
| `status` | ***[ElastictranscoderPresetStatus](#elastictranscoderpresetstatus)***||
## ElastictranscoderPresetSpec

Appears on:[ElastictranscoderPreset](#elastictranscoderpreset), [ElastictranscoderPresetStatus](#elastictranscoderpresetstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `audio` | ***[[]ElastictranscoderPresetSpecAudio](#elastictranscoderpresetspecaudio)***| ***(Optional)*** |
| `audioCodecOptions` | ***[[]ElastictranscoderPresetSpecAudioCodecOptions](#elastictranscoderpresetspecaudiocodecoptions)***| ***(Optional)*** |
| `container` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `thumbnails` | ***[[]ElastictranscoderPresetSpecThumbnails](#elastictranscoderpresetspecthumbnails)***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `video` | ***[[]ElastictranscoderPresetSpecVideo](#elastictranscoderpresetspecvideo)***| ***(Optional)*** |
| `videoCodecOptions` | ***map[string]string***| ***(Optional)*** |
| `videoWatermarks` | ***[[]ElastictranscoderPresetSpecVideoWatermarks](#elastictranscoderpresetspecvideowatermarks)***| ***(Optional)*** |
## ElastictranscoderPresetSpecAudio

Appears on:[ElastictranscoderPresetSpec](#elastictranscoderpresetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `audioPackingMode` | ***string***| ***(Optional)*** |
| `bitRate` | ***string***| ***(Optional)*** |
| `channels` | ***string***| ***(Optional)*** |
| `codec` | ***string***| ***(Optional)*** |
| `sampleRate` | ***string***| ***(Optional)*** |
## ElastictranscoderPresetSpecAudioCodecOptions

Appears on:[ElastictranscoderPresetSpec](#elastictranscoderpresetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `bitDepth` | ***string***| ***(Optional)*** |
| `bitOrder` | ***string***| ***(Optional)*** |
| `profile` | ***string***| ***(Optional)*** |
| `signed` | ***string***| ***(Optional)*** |
## ElastictranscoderPresetSpecThumbnails

Appears on:[ElastictranscoderPresetSpec](#elastictranscoderpresetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `aspectRatio` | ***string***| ***(Optional)*** |
| `format` | ***string***| ***(Optional)*** |
| `interval` | ***string***| ***(Optional)*** |
| `maxHeight` | ***string***| ***(Optional)*** |
| `maxWidth` | ***string***| ***(Optional)*** |
| `paddingPolicy` | ***string***| ***(Optional)*** |
| `resolution` | ***string***| ***(Optional)*** |
| `sizingPolicy` | ***string***| ***(Optional)*** |
## ElastictranscoderPresetSpecVideo

Appears on:[ElastictranscoderPresetSpec](#elastictranscoderpresetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `aspectRatio` | ***string***| ***(Optional)*** |
| `bitRate` | ***string***| ***(Optional)*** |
| `codec` | ***string***| ***(Optional)*** |
| `displayAspectRatio` | ***string***| ***(Optional)*** |
| `fixedGop` | ***string***| ***(Optional)*** |
| `frameRate` | ***string***| ***(Optional)*** |
| `keyframesMaxDist` | ***string***| ***(Optional)*** |
| `maxFrameRate` | ***string***| ***(Optional)*** |
| `maxHeight` | ***string***| ***(Optional)*** |
| `maxWidth` | ***string***| ***(Optional)*** |
| `paddingPolicy` | ***string***| ***(Optional)*** |
| `resolution` | ***string***| ***(Optional)*** |
| `sizingPolicy` | ***string***| ***(Optional)*** |
## ElastictranscoderPresetSpecVideoWatermarks

Appears on:[ElastictranscoderPresetSpec](#elastictranscoderpresetspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `horizontalAlign` | ***string***| ***(Optional)*** |
| `horizontalOffset` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `maxHeight` | ***string***| ***(Optional)*** |
| `maxWidth` | ***string***| ***(Optional)*** |
| `opacity` | ***string***| ***(Optional)*** |
| `sizingPolicy` | ***string***| ***(Optional)*** |
| `target` | ***string***| ***(Optional)*** |
| `verticalAlign` | ***string***| ***(Optional)*** |
| `verticalOffset` | ***string***| ***(Optional)*** |
## ElastictranscoderPresetStatus

Appears on:[ElastictranscoderPreset](#elastictranscoderpreset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElastictranscoderPresetSpec](#elastictranscoderpresetspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ElastictranscoderPresetStatus](#elastictranscoderpresetstatus)

---
