## ServiceDiscoveryPublicDNSNamespace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceDiscoveryPublicDNSNamespace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceDiscoveryPublicDNSNamespaceSpec](#ServiceDiscoveryPublicDNSNamespaceSpec)***||
| `status` | ***[ServiceDiscoveryPublicDNSNamespaceStatus](#ServiceDiscoveryPublicDNSNamespaceStatus)***||
## ServiceDiscoveryPublicDNSNamespaceSpec
##### (Appears on:[ServiceDiscoveryPublicDNSNamespace](#ServiceDiscoveryPublicDNSNamespace), [ServiceDiscoveryPublicDNSNamespaceStatus](#ServiceDiscoveryPublicDNSNamespaceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `hostedZone` | ***string***| ***(Optional)*** |
| `name` | ***string***||
## ServiceDiscoveryPublicDNSNamespaceStatus
##### (Appears on:[ServiceDiscoveryPublicDNSNamespace](#ServiceDiscoveryPublicDNSNamespace))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceDiscoveryPublicDNSNamespaceSpec](#ServiceDiscoveryPublicDNSNamespaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
