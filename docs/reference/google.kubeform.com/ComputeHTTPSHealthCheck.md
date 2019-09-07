---
title: ComputeHTTPSHealthCheck
menu:
  docs_v0.0.1:
    identifier: computehttpshealthcheck-google.kubeform.com
    name: ComputeHTTPSHealthCheck
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ComputeHTTPSHealthCheck
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeHTTPSHealthCheck` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeHTTPSHealthCheckSpec](#ComputeHTTPSHealthCheckSpec)***||
| `status` | ***[ComputeHTTPSHealthCheckStatus](#ComputeHTTPSHealthCheckStatus)***||
## ComputeHTTPSHealthCheckSpec
##### (Appears on:[ComputeHTTPSHealthCheck](#ComputeHTTPSHealthCheck), [ComputeHTTPSHealthCheckStatus](#ComputeHTTPSHealthCheckStatus))
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
## ComputeHTTPSHealthCheckStatus
##### (Appears on:[ComputeHTTPSHealthCheck](#ComputeHTTPSHealthCheck))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeHTTPSHealthCheckSpec](#ComputeHTTPSHealthCheckSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
