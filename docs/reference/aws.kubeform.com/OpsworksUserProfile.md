## OpsworksUserProfile
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `OpsworksUserProfile` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[OpsworksUserProfileSpec](#OpsworksUserProfileSpec)***||
| `status` | ***[OpsworksUserProfileStatus](#OpsworksUserProfileStatus)***||
## OpsworksUserProfileSpec
##### (Appears on:[OpsworksUserProfile](#OpsworksUserProfile), [OpsworksUserProfileStatus](#OpsworksUserProfileStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowSelfManagement` | ***bool***| ***(Optional)*** |
| `sshPublicKey` | ***string***| ***(Optional)*** |
| `sshUsername` | ***string***||
| `userArn` | ***string***||
## OpsworksUserProfileStatus
##### (Appears on:[OpsworksUserProfile](#OpsworksUserProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[OpsworksUserProfileSpec](#OpsworksUserProfileSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
