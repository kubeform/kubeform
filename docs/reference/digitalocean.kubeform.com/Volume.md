## Volume
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Volume` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VolumeSpec](#VolumeSpec)***||
| `status` | ***[VolumeStatus](#VolumeStatus)***||
## VolumeSpec
##### (Appears on:[Volume](#Volume), [VolumeStatus](#VolumeStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `dropletIDS` | ***[]int64***| ***(Optional)*** |
| `filesystemLabel` | ***string***| ***(Optional)*** |
| `filesystemType` | ***string***| ***(Optional)*** Deprecated|
| `initialFilesystemLabel` | ***string***| ***(Optional)*** |
| `initialFilesystemType` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `region` | ***string***||
| `size` | ***int***||
| `snapshotID` | ***string***| ***(Optional)*** |
| `urn` | ***string***| ***(Optional)*** the uniform resource name for the volume.|
## VolumeStatus
##### (Appears on:[Volume](#Volume))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VolumeSpec](#VolumeSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
