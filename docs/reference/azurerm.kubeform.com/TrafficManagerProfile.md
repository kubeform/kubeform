## TrafficManagerProfile
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `TrafficManagerProfile` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[TrafficManagerProfileSpec](#TrafficManagerProfileSpec)***||
| `status` | ***[TrafficManagerProfileStatus](#TrafficManagerProfileStatus)***||
## TrafficManagerProfileSpec
##### (Appears on:[TrafficManagerProfile](#TrafficManagerProfile), [TrafficManagerProfileStatus](#TrafficManagerProfileStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dnsConfig` | ***[[]TrafficManagerProfileSpecDnsConfig](#TrafficManagerProfileSpecDnsConfig)***||
| `fqdn` | ***string***| ***(Optional)*** |
| `monitorConfig` | ***[[]TrafficManagerProfileSpecMonitorConfig](#TrafficManagerProfileSpecMonitorConfig)***||
| `name` | ***string***||
| `profileStatus` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `trafficRoutingMethod` | ***string***||
## TrafficManagerProfileSpecDnsConfig
##### (Appears on:[TrafficManagerProfileSpec](#TrafficManagerProfileSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `relativeName` | ***string***||
| `ttl` | ***int***||
## TrafficManagerProfileSpecMonitorConfig
##### (Appears on:[TrafficManagerProfileSpec](#TrafficManagerProfileSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***int***||
| `protocol` | ***string***||
## TrafficManagerProfileStatus
##### (Appears on:[TrafficManagerProfile](#TrafficManagerProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[TrafficManagerProfileSpec](#TrafficManagerProfileSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
