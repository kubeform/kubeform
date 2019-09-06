## Route53Zone
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Route53Zone` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53ZoneSpec](#Route53ZoneSpec)***||
| `status` | ***[Route53ZoneStatus](#Route53ZoneStatus)***||
## Route53ZoneSpec
##### (Appears on:[Route53Zone](#Route53Zone), [Route53ZoneStatus](#Route53ZoneStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `comment` | ***string***| ***(Optional)*** |
| `delegationSetID` | ***string***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `nameServers` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpc` | ***[[]Route53ZoneSpecVpc](#Route53ZoneSpecVpc)***| ***(Optional)*** |
| `zoneID` | ***string***| ***(Optional)*** |
## Route53ZoneSpecVpc
##### (Appears on:[Route53ZoneSpec](#Route53ZoneSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `vpcID` | ***string***||
| `vpcRegion` | ***string***| ***(Optional)*** |
## Route53ZoneStatus
##### (Appears on:[Route53Zone](#Route53Zone))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53ZoneSpec](#Route53ZoneSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
