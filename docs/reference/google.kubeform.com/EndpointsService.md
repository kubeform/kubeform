---
title: EndpointsService
menu:
  docs_v0.0.1:
    identifier: endpointsservice-google.kubeform.com
    name: EndpointsService
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## EndpointsService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `EndpointsService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[EndpointsServiceSpec](#EndpointsServiceSpec)***||
| `status` | ***[EndpointsServiceStatus](#EndpointsServiceStatus)***||
## EndpointsServiceSpec
##### (Appears on:[EndpointsService](#EndpointsService), [EndpointsServiceStatus](#EndpointsServiceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `apis` | ***[[]EndpointsServiceSpecApis](#EndpointsServiceSpecApis)***| ***(Optional)*** |
| `configID` | ***string***| ***(Optional)*** |
| `dnsAddress` | ***string***| ***(Optional)*** |
| `endpoints` | ***[[]EndpointsServiceSpecEndpoints](#EndpointsServiceSpecEndpoints)***| ***(Optional)*** |
| `grpcConfig` | ***string***| ***(Optional)*** |
| `openapiConfig` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `protocOutput` | ***string***| ***(Optional)*** Deprecated|
| `protocOutputBase64` | ***string***| ***(Optional)*** |
| `serviceName` | ***string***||
## EndpointsServiceSpecApis
##### (Appears on:[EndpointsServiceSpec](#EndpointsServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `methods` | ***[[]EndpointsServiceSpecApisMethods](#EndpointsServiceSpecApisMethods)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `syntax` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## EndpointsServiceSpecApisMethods
##### (Appears on:[EndpointsServiceSpecApis](#EndpointsServiceSpecApis))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `requestType` | ***string***| ***(Optional)*** |
| `responseType` | ***string***| ***(Optional)*** |
| `syntax` | ***string***| ***(Optional)*** |
## EndpointsServiceSpecEndpoints
##### (Appears on:[EndpointsServiceSpec](#EndpointsServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `address` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
## EndpointsServiceStatus
##### (Appears on:[EndpointsService](#EndpointsService))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[EndpointsServiceSpec](#EndpointsServiceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
