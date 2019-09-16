---
title: ServiceDiscoveryService
menu:
  docs_{{ .version }}:
    identifier: servicediscoveryservice-aws.kubeform.com
    name: ServiceDiscoveryService
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServiceDiscoveryService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceDiscoveryService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceDiscoveryServiceSpec](#servicediscoveryservicespec)***||
| `status` | ***[ServiceDiscoveryServiceStatus](#servicediscoveryservicestatus)***||
## ServiceDiscoveryServiceSpec

Appears on:[ServiceDiscoveryService](#servicediscoveryservice), [ServiceDiscoveryServiceStatus](#servicediscoveryservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `dnsConfig` | ***[[]ServiceDiscoveryServiceSpecDnsConfig](#servicediscoveryservicespecdnsconfig)***||
| `healthCheckConfig` | ***[[]ServiceDiscoveryServiceSpecHealthCheckConfig](#servicediscoveryservicespechealthcheckconfig)***| ***(Optional)*** |
| `healthCheckCustomConfig` | ***[[]ServiceDiscoveryServiceSpecHealthCheckCustomConfig](#servicediscoveryservicespechealthcheckcustomconfig)***| ***(Optional)*** |
| `name` | ***string***||
## ServiceDiscoveryServiceSpecDnsConfig

Appears on:[ServiceDiscoveryServiceSpec](#servicediscoveryservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsRecords` | ***[[]ServiceDiscoveryServiceSpecDnsConfigDnsRecords](#servicediscoveryservicespecdnsconfigdnsrecords)***||
| `namespaceID` | ***string***||
| `routingPolicy` | ***string***| ***(Optional)*** |
## ServiceDiscoveryServiceSpecDnsConfigDnsRecords

Appears on:[ServiceDiscoveryServiceSpecDnsConfig](#servicediscoveryservicespecdnsconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ttl` | ***int***||
| `type` | ***string***||
## ServiceDiscoveryServiceSpecHealthCheckConfig

Appears on:[ServiceDiscoveryServiceSpec](#servicediscoveryservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `failureThreshold` | ***int***| ***(Optional)*** |
| `resourcePath` | ***string***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## ServiceDiscoveryServiceSpecHealthCheckCustomConfig

Appears on:[ServiceDiscoveryServiceSpec](#servicediscoveryservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `failureThreshold` | ***int***| ***(Optional)*** |
## ServiceDiscoveryServiceStatus

Appears on:[ServiceDiscoveryService](#servicediscoveryservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceDiscoveryServiceSpec](#servicediscoveryservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
