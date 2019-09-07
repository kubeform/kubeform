---
title: ServiceDiscoveryService
menu:
  docs_v0.0.1:
    identifier: servicediscoveryservice-aws.kubeform.com
    name: ServiceDiscoveryService
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ServiceDiscoveryService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceDiscoveryService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceDiscoveryServiceSpec](#ServiceDiscoveryServiceSpec)***||
| `status` | ***[ServiceDiscoveryServiceStatus](#ServiceDiscoveryServiceStatus)***||
## ServiceDiscoveryServiceSpec
##### (Appears on:[ServiceDiscoveryService](#ServiceDiscoveryService), [ServiceDiscoveryServiceStatus](#ServiceDiscoveryServiceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `dnsConfig` | ***[[]ServiceDiscoveryServiceSpecDnsConfig](#ServiceDiscoveryServiceSpecDnsConfig)***||
| `healthCheckConfig` | ***[[]ServiceDiscoveryServiceSpecHealthCheckConfig](#ServiceDiscoveryServiceSpecHealthCheckConfig)***| ***(Optional)*** |
| `healthCheckCustomConfig` | ***[[]ServiceDiscoveryServiceSpecHealthCheckCustomConfig](#ServiceDiscoveryServiceSpecHealthCheckCustomConfig)***| ***(Optional)*** |
| `name` | ***string***||
## ServiceDiscoveryServiceSpecDnsConfig
##### (Appears on:[ServiceDiscoveryServiceSpec](#ServiceDiscoveryServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsRecords` | ***[[]ServiceDiscoveryServiceSpecDnsConfigDnsRecords](#ServiceDiscoveryServiceSpecDnsConfigDnsRecords)***||
| `namespaceID` | ***string***||
| `routingPolicy` | ***string***| ***(Optional)*** |
## ServiceDiscoveryServiceSpecDnsConfigDnsRecords
##### (Appears on:[ServiceDiscoveryServiceSpecDnsConfig](#ServiceDiscoveryServiceSpecDnsConfig))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ttl` | ***int***||
| `type` | ***string***||
## ServiceDiscoveryServiceSpecHealthCheckConfig
##### (Appears on:[ServiceDiscoveryServiceSpec](#ServiceDiscoveryServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `failureThreshold` | ***int***| ***(Optional)*** |
| `resourcePath` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## ServiceDiscoveryServiceSpecHealthCheckCustomConfig
##### (Appears on:[ServiceDiscoveryServiceSpec](#ServiceDiscoveryServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `failureThreshold` | ***int***| ***(Optional)*** |
## ServiceDiscoveryServiceStatus
##### (Appears on:[ServiceDiscoveryService](#ServiceDiscoveryService))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceDiscoveryServiceSpec](#ServiceDiscoveryServiceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
