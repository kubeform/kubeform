---
title: ComputeHealthCheck
menu:
  docs_{{ .version }}:
    identifier: computehealthcheck-google.kubeform.com
    name: ComputeHealthCheck
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeHealthCheck
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeHealthCheck` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeHealthCheckSpec](#ComputeHealthCheckSpec)***||
| `status` | ***[ComputeHealthCheckStatus](#ComputeHealthCheckStatus)***||
## ComputeHealthCheckSpec

Appears on:[ComputeHealthCheck](#ComputeHealthCheck), [ComputeHealthCheckStatus](#ComputeHealthCheckStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `checkIntervalSec` | ***int***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `healthyThreshold` | ***int***| ***(Optional)*** |
| `httpHealthCheck` | ***[[]ComputeHealthCheckSpecHttpHealthCheck](#ComputeHealthCheckSpecHttpHealthCheck)***| ***(Optional)*** |
| `httpsHealthCheck` | ***[[]ComputeHealthCheckSpecHttpsHealthCheck](#ComputeHealthCheckSpecHttpsHealthCheck)***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sslHealthCheck` | ***[[]ComputeHealthCheckSpecSslHealthCheck](#ComputeHealthCheckSpecSslHealthCheck)***| ***(Optional)*** |
| `tcpHealthCheck` | ***[[]ComputeHealthCheckSpecTcpHealthCheck](#ComputeHealthCheckSpecTcpHealthCheck)***| ***(Optional)*** |
| `timeoutSec` | ***int***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
| `unhealthyThreshold` | ***int***| ***(Optional)*** |
## ComputeHealthCheckSpecHttpHealthCheck

Appears on:[ComputeHealthCheckSpec](#ComputeHealthCheckSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `proxyHeader` | ***string***| ***(Optional)*** |
| `requestPath` | ***string***| ***(Optional)*** |
| `response` | ***string***| ***(Optional)*** |
## ComputeHealthCheckSpecHttpsHealthCheck

Appears on:[ComputeHealthCheckSpec](#ComputeHealthCheckSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `proxyHeader` | ***string***| ***(Optional)*** |
| `requestPath` | ***string***| ***(Optional)*** |
| `response` | ***string***| ***(Optional)*** |
## ComputeHealthCheckSpecSslHealthCheck

Appears on:[ComputeHealthCheckSpec](#ComputeHealthCheckSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int***| ***(Optional)*** |
| `proxyHeader` | ***string***| ***(Optional)*** |
| `request` | ***string***| ***(Optional)*** |
| `response` | ***string***| ***(Optional)*** |
## ComputeHealthCheckSpecTcpHealthCheck

Appears on:[ComputeHealthCheckSpec](#ComputeHealthCheckSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int***| ***(Optional)*** |
| `proxyHeader` | ***string***| ***(Optional)*** |
| `request` | ***string***| ***(Optional)*** |
| `response` | ***string***| ***(Optional)*** |
## ComputeHealthCheckStatus

Appears on:[ComputeHealthCheck](#ComputeHealthCheck)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeHealthCheckSpec](#ComputeHealthCheckSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
