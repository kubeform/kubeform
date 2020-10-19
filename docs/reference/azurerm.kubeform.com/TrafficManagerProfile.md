---
title: TrafficManagerProfile
menu:
  docs_{{ .version }}:
    identifier: trafficmanagerprofile-azurerm.kubeform.com
    name: TrafficManagerProfile
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## TrafficManagerProfile
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `TrafficManagerProfile` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[TrafficManagerProfileSpec](#trafficmanagerprofilespec)***||
| `status` | ***[TrafficManagerProfileStatus](#trafficmanagerprofilestatus)***||
## Phase(`string` alias)

Appears on:[TrafficManagerProfileStatus](#trafficmanagerprofilestatus)

## TrafficManagerProfileSpec

Appears on:[TrafficManagerProfile](#trafficmanagerprofile), [TrafficManagerProfileStatus](#trafficmanagerprofilestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dnsConfig` | ***[[]TrafficManagerProfileSpecDnsConfig](#trafficmanagerprofilespecdnsconfig)***||
| `fqdn` | ***string***| ***(Optional)*** |
| `monitorConfig` | ***[[]TrafficManagerProfileSpecMonitorConfig](#trafficmanagerprofilespecmonitorconfig)***||
| `name` | ***string***||
| `profileStatus` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `trafficRoutingMethod` | ***string***||
## TrafficManagerProfileSpecDnsConfig

Appears on:[TrafficManagerProfileSpec](#trafficmanagerprofilespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `relativeName` | ***string***||
| `ttl` | ***int64***||
## TrafficManagerProfileSpecMonitorConfig

Appears on:[TrafficManagerProfileSpec](#trafficmanagerprofilespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `expectedStatusCodeRanges` | ***[]string***| ***(Optional)*** |
| `intervalInSeconds` | ***int64***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***int64***||
| `protocol` | ***string***||
| `timeoutInSeconds` | ***int64***| ***(Optional)*** |
| `toleratedNumberOfFailures` | ***int64***| ***(Optional)*** |
## TrafficManagerProfileStatus

Appears on:[TrafficManagerProfile](#trafficmanagerprofile)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[TrafficManagerProfileSpec](#trafficmanagerprofilespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
