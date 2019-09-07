---
title: ComputeHTTPHealthCheck
menu:
  docs_v0.0.1:
    identifier: computehttphealthcheck-google.kubeform.com
    name: ComputeHTTPHealthCheck
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ComputeHTTPHealthCheck
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeHTTPHealthCheck` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeHTTPHealthCheckSpec](#ComputeHTTPHealthCheckSpec)***||
| `status` | ***[ComputeHTTPHealthCheckStatus](#ComputeHTTPHealthCheckStatus)***||
## ComputeHTTPHealthCheckSpec
##### (Appears on:[ComputeHTTPHealthCheck](#ComputeHTTPHealthCheck), [ComputeHTTPHealthCheckStatus](#ComputeHTTPHealthCheckStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `checkIntervalSec` | ***int***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `healthyThreshold` | ***int***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `port` | ***int***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `requestPath` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `timeoutSec` | ***int***| ***(Optional)*** |
| `unhealthyThreshold` | ***int***| ***(Optional)*** |
## ComputeHTTPHealthCheckStatus
##### (Appears on:[ComputeHTTPHealthCheck](#ComputeHTTPHealthCheck))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeHTTPHealthCheckSpec](#ComputeHTTPHealthCheckSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
