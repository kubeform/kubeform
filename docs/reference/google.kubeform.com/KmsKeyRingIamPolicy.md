## KmsKeyRingIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsKeyRingIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsKeyRingIamPolicySpec](#KmsKeyRingIamPolicySpec)***||
| `status` | ***[KmsKeyRingIamPolicyStatus](#KmsKeyRingIamPolicyStatus)***||
## KmsKeyRingIamPolicySpec
##### (Appears on:[KmsKeyRingIamPolicy](#KmsKeyRingIamPolicy), [KmsKeyRingIamPolicyStatus](#KmsKeyRingIamPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `keyRingID` | ***string***||
| `policyData` | ***string***||
## KmsKeyRingIamPolicyStatus
##### (Appears on:[KmsKeyRingIamPolicy](#KmsKeyRingIamPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsKeyRingIamPolicySpec](#KmsKeyRingIamPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
