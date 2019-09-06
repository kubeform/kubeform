## ConfigConfigurationRecorder
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ConfigConfigurationRecorder` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ConfigConfigurationRecorderSpec](#ConfigConfigurationRecorderSpec)***||
| `status` | ***[ConfigConfigurationRecorderStatus](#ConfigConfigurationRecorderStatus)***||
## ConfigConfigurationRecorderSpec
##### (Appears on:[ConfigConfigurationRecorder](#ConfigConfigurationRecorder), [ConfigConfigurationRecorderStatus](#ConfigConfigurationRecorderStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `recordingGroup` | ***[[]ConfigConfigurationRecorderSpecRecordingGroup](#ConfigConfigurationRecorderSpecRecordingGroup)***| ***(Optional)*** |
| `roleArn` | ***string***||
## ConfigConfigurationRecorderSpecRecordingGroup
##### (Appears on:[ConfigConfigurationRecorderSpec](#ConfigConfigurationRecorderSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allSupported` | ***bool***| ***(Optional)*** |
| `includeGlobalResourceTypes` | ***bool***| ***(Optional)*** |
| `resourceTypes` | ***[]string***| ***(Optional)*** |
## ConfigConfigurationRecorderStatus
##### (Appears on:[ConfigConfigurationRecorder](#ConfigConfigurationRecorder))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ConfigConfigurationRecorderSpec](#ConfigConfigurationRecorderSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
