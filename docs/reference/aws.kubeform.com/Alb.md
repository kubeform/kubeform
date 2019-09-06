## Alb
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Alb` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AlbSpec](#AlbSpec)***||
| `status` | ***[AlbStatus](#AlbStatus)***||
## AlbSpec
##### (Appears on:[Alb](#Alb), [AlbStatus](#AlbStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessLogs` | ***[[]AlbSpecAccessLogs](#AlbSpecAccessLogs)***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `arnSuffix` | ***string***| ***(Optional)*** |
| `dnsName` | ***string***| ***(Optional)*** |
| `enableCrossZoneLoadBalancing` | ***bool***| ***(Optional)*** |
| `enableDeletionProtection` | ***bool***| ***(Optional)*** |
| `enableHttp2` | ***bool***| ***(Optional)*** |
| `idleTimeout` | ***int***| ***(Optional)*** |
| `internal` | ***bool***| ***(Optional)*** |
| `ipAddressType` | ***string***| ***(Optional)*** |
| `loadBalancerType` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `subnetMapping` | ***[[]AlbSpecSubnetMapping](#AlbSpecSubnetMapping)***| ***(Optional)*** |
| `subnets` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
| `zoneID` | ***string***| ***(Optional)*** |
## AlbSpecAccessLogs
##### (Appears on:[AlbSpec](#AlbSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `bucket` | ***string***||
| `enabled` | ***bool***| ***(Optional)*** |
| `prefix` | ***string***| ***(Optional)*** |
## AlbSpecSubnetMapping
##### (Appears on:[AlbSpec](#AlbSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allocationID` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***||
## AlbStatus
##### (Appears on:[Alb](#Alb))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AlbSpec](#AlbSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
