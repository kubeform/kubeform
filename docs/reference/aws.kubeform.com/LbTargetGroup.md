## LbTargetGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbTargetGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbTargetGroupSpec](#LbTargetGroupSpec)***||
| `status` | ***[LbTargetGroupStatus](#LbTargetGroupStatus)***||
## LbTargetGroupSpec
##### (Appears on:[LbTargetGroup](#LbTargetGroup), [LbTargetGroupStatus](#LbTargetGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `arnSuffix` | ***string***| ***(Optional)*** |
| `deregistrationDelay` | ***int***| ***(Optional)*** |
| `healthCheck` | ***[[]LbTargetGroupSpecHealthCheck](#LbTargetGroupSpecHealthCheck)***| ***(Optional)*** |
| `lambdaMultiValueHeadersEnabled` | ***bool***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `proxyProtocolV2` | ***bool***| ***(Optional)*** |
| `slowStart` | ***int***| ***(Optional)*** |
| `stickiness` | ***[[]LbTargetGroupSpecStickiness](#LbTargetGroupSpecStickiness)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetType` | ***string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## LbTargetGroupSpecHealthCheck
##### (Appears on:[LbTargetGroupSpec](#LbTargetGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enabled` | ***bool***| ***(Optional)*** |
| `healthyThreshold` | ***int***| ***(Optional)*** |
| `interval` | ***int***| ***(Optional)*** |
| `matcher` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `timeout` | ***int***| ***(Optional)*** |
| `unhealthyThreshold` | ***int***| ***(Optional)*** |
## LbTargetGroupSpecStickiness
##### (Appears on:[LbTargetGroupSpec](#LbTargetGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cookieDuration` | ***int***| ***(Optional)*** |
| `enabled` | ***bool***| ***(Optional)*** |
| `type` | ***string***||
## LbTargetGroupStatus
##### (Appears on:[LbTargetGroup](#LbTargetGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbTargetGroupSpec](#LbTargetGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
