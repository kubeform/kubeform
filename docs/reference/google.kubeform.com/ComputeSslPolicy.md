### ComputeSslPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeSslPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeSslPolicySpec](#ComputeSslPolicySpec)***||
| `status` | ***[ComputeSslPolicyStatus](#ComputeSslPolicyStatus)***||
### ComputeSslPolicySpec
(<em>Appears on:</em>
[ComputeSslPolicy](#ComputeSslPolicy), 
[ComputeSslPolicyStatus](#ComputeSslPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `customFeatures` | ***[]string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `enabledFeatures` | ***[]string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `minTlsVersion` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `profile` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
### ComputeSslPolicyStatus
(<em>Appears on:</em>
[ComputeSslPolicy](#ComputeSslPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeSslPolicySpec](#ComputeSslPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
