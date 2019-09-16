---
title: ConfigConfigurationRecorder
menu:
  docs_{{ .version }}:
    identifier: configconfigurationrecorder-aws.kubeform.com
    name: ConfigConfigurationRecorder
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ConfigConfigurationRecorder
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigConfigurationRecorder` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigConfigurationRecorderSpec](#configconfigurationrecorderspec)***||
| `status` | ***[ConfigConfigurationRecorderStatus](#configconfigurationrecorderstatus)***||
## ConfigConfigurationRecorderSpec

Appears on:[ConfigConfigurationRecorder](#configconfigurationrecorder), [ConfigConfigurationRecorderStatus](#configconfigurationrecorderstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `recordingGroup` | ***[[]ConfigConfigurationRecorderSpecRecordingGroup](#configconfigurationrecorderspecrecordinggroup)***| ***(Optional)*** |
| `roleArn` | ***string***||
## ConfigConfigurationRecorderSpecRecordingGroup

Appears on:[ConfigConfigurationRecorderSpec](#configconfigurationrecorderspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allSupported` | ***bool***| ***(Optional)*** |
| `includeGlobalResourceTypes` | ***bool***| ***(Optional)*** |
| `resourceTypes` | ***[]string***| ***(Optional)*** |
## ConfigConfigurationRecorderStatus

Appears on:[ConfigConfigurationRecorder](#configconfigurationrecorder)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigConfigurationRecorderSpec](#configconfigurationrecorderspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
