## Eip
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Eip` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EipSpec](#EipSpec)***||
| `status` | ***[EipStatus](#EipStatus)***||
## EipSpec
##### (Appears on:[Eip](#Eip), [EipStatus](#EipStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allocationID` | ***string***| ***(Optional)*** |
| `associateWithPrivateIP` | ***string***| ***(Optional)*** |
| `associationID` | ***string***| ***(Optional)*** |
| `domain` | ***string***| ***(Optional)*** |
| `instance` | ***string***| ***(Optional)*** |
| `networkInterface` | ***string***| ***(Optional)*** |
| `privateDNS` | ***string***| ***(Optional)*** |
| `privateIP` | ***string***| ***(Optional)*** |
| `publicDNS` | ***string***| ***(Optional)*** |
| `publicIP` | ***string***| ***(Optional)*** |
| `publicIpv4Pool` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpc` | ***bool***| ***(Optional)*** |
## EipStatus
##### (Appears on:[Eip](#Eip))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EipSpec](#EipSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
