## RedisInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `RedisInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RedisInstanceSpec](#RedisInstanceSpec)***||
| `status` | ***[RedisInstanceStatus](#RedisInstanceStatus)***||
## RedisInstanceSpec
##### (Appears on:[RedisInstance](#RedisInstance), [RedisInstanceStatus](#RedisInstanceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alternativeLocationID` | ***string***| ***(Optional)*** |
| `authorizedNetwork` | ***string***| ***(Optional)*** |
| `createTime` | ***string***| ***(Optional)*** |
| `currentLocationID` | ***string***| ***(Optional)*** |
| `displayName` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `locationID` | ***string***| ***(Optional)*** |
| `memorySizeGb` | ***int***||
| `name` | ***string***||
| `port` | ***int***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `redisConfigs` | ***map[string]string***| ***(Optional)*** |
| `redisVersion` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `reservedIPRange` | ***string***| ***(Optional)*** |
| `tier` | ***string***| ***(Optional)*** |
## RedisInstanceStatus
##### (Appears on:[RedisInstance](#RedisInstance))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RedisInstanceSpec](#RedisInstanceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
