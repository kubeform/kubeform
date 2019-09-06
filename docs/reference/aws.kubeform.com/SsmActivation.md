## SsmActivation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmActivation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmActivationSpec](#SsmActivationSpec)***||
| `status` | ***[SsmActivationStatus](#SsmActivationStatus)***||
## SsmActivationSpec
##### (Appears on:[SsmActivation](#SsmActivation), [SsmActivationStatus](#SsmActivationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activationCode` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `expirationDate` | ***string***| ***(Optional)*** |
| `expired` | ***string***| ***(Optional)*** |
| `iamRole` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `registrationCount` | ***int***| ***(Optional)*** |
| `registrationLimit` | ***int***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SsmActivationStatus
##### (Appears on:[SsmActivation](#SsmActivation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmActivationSpec](#SsmActivationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
