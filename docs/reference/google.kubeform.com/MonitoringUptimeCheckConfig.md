## MonitoringUptimeCheckConfig
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `MonitoringUptimeCheckConfig` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MonitoringUptimeCheckConfigSpec](#MonitoringUptimeCheckConfigSpec)***||
| `status` | ***[MonitoringUptimeCheckConfigStatus](#MonitoringUptimeCheckConfigStatus)***||
## MonitoringUptimeCheckConfigSpec
##### (Appears on:[MonitoringUptimeCheckConfig](#MonitoringUptimeCheckConfig), [MonitoringUptimeCheckConfigStatus](#MonitoringUptimeCheckConfigStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `contentMatchers` | ***[[]MonitoringUptimeCheckConfigSpecContentMatchers](#MonitoringUptimeCheckConfigSpecContentMatchers)***| ***(Optional)*** |
| `displayName` | ***string***||
| `httpCheck` | ***[[]MonitoringUptimeCheckConfigSpecHttpCheck](#MonitoringUptimeCheckConfigSpecHttpCheck)***| ***(Optional)*** |
| `internalCheckers` | ***[[]MonitoringUptimeCheckConfigSpecInternalCheckers](#MonitoringUptimeCheckConfigSpecInternalCheckers)***| ***(Optional)*** |
| `isInternal` | ***bool***| ***(Optional)*** |
| `monitoredResource` | ***[[]MonitoringUptimeCheckConfigSpecMonitoredResource](#MonitoringUptimeCheckConfigSpecMonitoredResource)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `period` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `resourceGroup` | ***[[]MonitoringUptimeCheckConfigSpecResourceGroup](#MonitoringUptimeCheckConfigSpecResourceGroup)***| ***(Optional)*** |
| `selectedRegions` | ***[]string***| ***(Optional)*** |
| `tcpCheck` | ***[[]MonitoringUptimeCheckConfigSpecTcpCheck](#MonitoringUptimeCheckConfigSpecTcpCheck)***| ***(Optional)*** |
| `timeout` | ***string***||
## MonitoringUptimeCheckConfigSpecContentMatchers
##### (Appears on:[MonitoringUptimeCheckConfigSpec](#MonitoringUptimeCheckConfigSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `content` | ***string***| ***(Optional)*** |
## MonitoringUptimeCheckConfigSpecHttpCheck
##### (Appears on:[MonitoringUptimeCheckConfigSpec](#MonitoringUptimeCheckConfigSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `authInfo` | ***[[]MonitoringUptimeCheckConfigSpecHttpCheckAuthInfo](#MonitoringUptimeCheckConfigSpecHttpCheckAuthInfo)***| ***(Optional)*** |
| `headers` | ***map[string]string***| ***(Optional)*** |
| `maskHeaders` | ***bool***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `useSSL` | ***bool***| ***(Optional)*** |
## MonitoringUptimeCheckConfigSpecHttpCheckAuthInfo
##### (Appears on:[MonitoringUptimeCheckConfigSpecHttpCheck](#MonitoringUptimeCheckConfigSpecHttpCheck))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `username` | ***string***| ***(Optional)*** |
## MonitoringUptimeCheckConfigSpecInternalCheckers
##### (Appears on:[MonitoringUptimeCheckConfigSpec](#MonitoringUptimeCheckConfigSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `displayName` | ***string***| ***(Optional)*** |
| `gcpZone` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `network` | ***string***| ***(Optional)*** |
| `peerProjectID` | ***string***| ***(Optional)*** |
## MonitoringUptimeCheckConfigSpecMonitoredResource
##### (Appears on:[MonitoringUptimeCheckConfigSpec](#MonitoringUptimeCheckConfigSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `labels` | ***map[string]string***||
| `type` | ***string***||
## MonitoringUptimeCheckConfigSpecResourceGroup
##### (Appears on:[MonitoringUptimeCheckConfigSpec](#MonitoringUptimeCheckConfigSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `groupID` | ***string***| ***(Optional)*** |
| `resourceType` | ***string***| ***(Optional)*** |
## MonitoringUptimeCheckConfigSpecTcpCheck
##### (Appears on:[MonitoringUptimeCheckConfigSpec](#MonitoringUptimeCheckConfigSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int***||
## MonitoringUptimeCheckConfigStatus
##### (Appears on:[MonitoringUptimeCheckConfig](#MonitoringUptimeCheckConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MonitoringUptimeCheckConfigSpec](#MonitoringUptimeCheckConfigSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `http_check.<index>.auth_info.<index>.password` | ***string*** ||
