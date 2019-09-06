### ComputeTargetSslProxy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeTargetSslProxy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeTargetSslProxySpec](#ComputeTargetSslProxySpec)***||
| `status` | ***[ComputeTargetSslProxyStatus](#ComputeTargetSslProxyStatus)***||
### ComputeTargetSslProxySpec
(<em>Appears on:</em>
[ComputeTargetSslProxy](#ComputeTargetSslProxy), 
[ComputeTargetSslProxyStatus](#ComputeTargetSslProxyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backendService` | ***string***||
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `proxyHeader` | ***string***| ***(Optional)*** |
| `proxyID` | ***int***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sslCertificates` | ***[]string***||
| `sslPolicy` | ***string***| ***(Optional)*** |
### ComputeTargetSslProxyStatus
(<em>Appears on:</em>
[ComputeTargetSslProxy](#ComputeTargetSslProxy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeTargetSslProxySpec](#ComputeTargetSslProxySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
