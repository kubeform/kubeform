---
title: ComputeHTTPSHealthCheck
menu:
  docs_{{ .version }}:
    identifier: computehttpshealthcheck-google.kubeform.com
    name: ComputeHTTPSHealthCheck
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeHTTPSHealthCheck
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeHTTPSHealthCheck` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeHTTPSHealthCheckSpec](#computehttpshealthcheckspec)***||
| `status` | ***[ComputeHTTPSHealthCheckStatus](#computehttpshealthcheckstatus)***||
## ComputeHTTPSHealthCheckSpec

Appears on:[ComputeHTTPSHealthCheck](#computehttpshealthcheck), [ComputeHTTPSHealthCheckStatus](#computehttpshealthcheckstatus)

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

Appears on:[ComputeHTTPSHealthCheck](#computehttpshealthcheck)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeHTTPSHealthCheckSpec](#computehttpshealthcheckspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeHTTPSHealthCheckStatus](#computehttpshealthcheckstatus)

---
