### LbSslNegotiationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbSslNegotiationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbSslNegotiationPolicySpec](#LbSslNegotiationPolicySpec)***||
| `status` | ***[LbSslNegotiationPolicyStatus](#LbSslNegotiationPolicyStatus)***||
### LbSslNegotiationPolicySpec
(<em>Appears on:</em>
[LbSslNegotiationPolicy](#LbSslNegotiationPolicy), 
[LbSslNegotiationPolicyStatus](#LbSslNegotiationPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attribute` | ***[[]LbSslNegotiationPolicySpecAttribute](#LbSslNegotiationPolicySpecAttribute)***| ***(Optional)*** |
| `lbPort` | ***int***||
| `loadBalancer` | ***string***||
| `name` | ***string***||
### LbSslNegotiationPolicySpecAttribute
(<em>Appears on:</em>
[LbSslNegotiationPolicySpec](#LbSslNegotiationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
### LbSslNegotiationPolicyStatus
(<em>Appears on:</em>
[LbSslNegotiationPolicy](#LbSslNegotiationPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbSslNegotiationPolicySpec](#LbSslNegotiationPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
