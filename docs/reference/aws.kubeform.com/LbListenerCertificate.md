## LbListenerCertificate
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbListenerCertificate` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbListenerCertificateSpec](#LbListenerCertificateSpec)***||
| `status` | ***[LbListenerCertificateStatus](#LbListenerCertificateStatus)***||
## LbListenerCertificateSpec
##### (Appears on:[LbListenerCertificate](#LbListenerCertificate), [LbListenerCertificateStatus](#LbListenerCertificateStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `certificateArn` | ***string***||
| `listenerArn` | ***string***||
## LbListenerCertificateStatus
##### (Appears on:[LbListenerCertificate](#LbListenerCertificate))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbListenerCertificateSpec](#LbListenerCertificateSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
