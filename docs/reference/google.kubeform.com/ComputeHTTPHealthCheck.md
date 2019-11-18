---
title: ComputeHTTPHealthCheck
menu:
  docs_{{ .version }}:
    identifier: computehttphealthcheck-google.kubeform.com
    name: ComputeHTTPHealthCheck
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeHTTPHealthCheck
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeHTTPHealthCheck` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeHTTPHealthCheckSpec](#computehttphealthcheckspec)***||
| `status` | ***[ComputeHTTPHealthCheckStatus](#computehttphealthcheckstatus)***||
## ComputeHTTPHealthCheckSpec

Appears on:[ComputeHTTPHealthCheck](#computehttphealthcheck), [ComputeHTTPHealthCheckStatus](#computehttphealthcheckstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `checkIntervalSec` | ***int64***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `healthyThreshold` | ***int64***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `port` | ***int64***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `requestPath` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `timeoutSec` | ***int64***| ***(Optional)*** |
| `unhealthyThreshold` | ***int64***| ***(Optional)*** |
## ComputeHTTPHealthCheckStatus

Appears on:[ComputeHTTPHealthCheck](#computehttphealthcheck)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeHTTPHealthCheckSpec](#computehttphealthcheckspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeHTTPHealthCheckStatus](#computehttphealthcheckstatus)

---
