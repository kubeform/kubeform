## LightsailStaticIP
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LightsailStaticIP` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LightsailStaticIPSpec](#LightsailStaticIPSpec)***||
| `status` | ***[LightsailStaticIPStatus](#LightsailStaticIPStatus)***||
## LightsailStaticIPSpec
##### (Appears on:[LightsailStaticIP](#LightsailStaticIP), [LightsailStaticIPStatus](#LightsailStaticIPStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `supportCode` | ***string***| ***(Optional)*** |
## LightsailStaticIPStatus
##### (Appears on:[LightsailStaticIP](#LightsailStaticIP))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LightsailStaticIPSpec](#LightsailStaticIPSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
