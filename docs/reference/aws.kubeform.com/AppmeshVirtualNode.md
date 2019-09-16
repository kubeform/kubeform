---
title: AppmeshVirtualNode
menu:
  docs_{{ .version }}:
    identifier: appmeshvirtualnode-aws.kubeform.com
    name: AppmeshVirtualNode
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppmeshVirtualNode
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AppmeshVirtualNode` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppmeshVirtualNodeSpec](#AppmeshVirtualNodeSpec)***||
| `status` | ***[AppmeshVirtualNodeStatus](#AppmeshVirtualNodeStatus)***||
## AppmeshVirtualNodeSpec

Appears on:[AppmeshVirtualNode](#AppmeshVirtualNode), [AppmeshVirtualNodeStatus](#AppmeshVirtualNodeStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `createdDate` | ***string***| ***(Optional)*** |
| `lastUpdatedDate` | ***string***| ***(Optional)*** |
| `meshName` | ***string***||
| `name` | ***string***||
| `spec` | ***[[]AppmeshVirtualNodeSpecSpec](#AppmeshVirtualNodeSpecSpec)***||
## AppmeshVirtualNodeSpecSpec

Appears on:[AppmeshVirtualNodeSpec](#AppmeshVirtualNodeSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `backend` | ***[[]AppmeshVirtualNodeSpecSpecBackend](#AppmeshVirtualNodeSpecSpecBackend)***| ***(Optional)*** |
| `listener` | ***[[]AppmeshVirtualNodeSpecSpecListener](#AppmeshVirtualNodeSpecSpecListener)***| ***(Optional)*** |
| `logging` | ***[[]AppmeshVirtualNodeSpecSpecLogging](#AppmeshVirtualNodeSpecSpecLogging)***| ***(Optional)*** |
| `serviceDiscovery` | ***[[]AppmeshVirtualNodeSpecSpecServiceDiscovery](#AppmeshVirtualNodeSpecSpecServiceDiscovery)***| ***(Optional)*** |
## AppmeshVirtualNodeSpecSpecBackend

Appears on:[AppmeshVirtualNodeSpecSpec](#AppmeshVirtualNodeSpecSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `virtualService` | ***[[]AppmeshVirtualNodeSpecSpecBackendVirtualService](#AppmeshVirtualNodeSpecSpecBackendVirtualService)***| ***(Optional)*** |
## AppmeshVirtualNodeSpecSpecBackendVirtualService

Appears on:[AppmeshVirtualNodeSpecSpecBackend](#AppmeshVirtualNodeSpecSpecBackend)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `virtualServiceName` | ***string***||
## AppmeshVirtualNodeSpecSpecListener

Appears on:[AppmeshVirtualNodeSpecSpec](#AppmeshVirtualNodeSpecSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `healthCheck` | ***[[]AppmeshVirtualNodeSpecSpecListenerHealthCheck](#AppmeshVirtualNodeSpecSpecListenerHealthCheck)***| ***(Optional)*** |
| `portMapping` | ***[[]AppmeshVirtualNodeSpecSpecListenerPortMapping](#AppmeshVirtualNodeSpecSpecListenerPortMapping)***||
## AppmeshVirtualNodeSpecSpecListenerHealthCheck

Appears on:[AppmeshVirtualNodeSpecSpecListener](#AppmeshVirtualNodeSpecSpecListener)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `healthyThreshold` | ***int***||
| `intervalMillis` | ***int***||
| `path` | ***string***| ***(Optional)*** |
| `port` | ***int***| ***(Optional)*** |
| `protocol` | ***string***||
| `timeoutMillis` | ***int***||
| `unhealthyThreshold` | ***int***||
## AppmeshVirtualNodeSpecSpecListenerPortMapping

Appears on:[AppmeshVirtualNodeSpecSpecListener](#AppmeshVirtualNodeSpecSpecListener)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `port` | ***int***||
| `protocol` | ***string***||
## AppmeshVirtualNodeSpecSpecLogging

Appears on:[AppmeshVirtualNodeSpecSpec](#AppmeshVirtualNodeSpecSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `accessLog` | ***[[]AppmeshVirtualNodeSpecSpecLoggingAccessLog](#AppmeshVirtualNodeSpecSpecLoggingAccessLog)***| ***(Optional)*** |
## AppmeshVirtualNodeSpecSpecLoggingAccessLog

Appears on:[AppmeshVirtualNodeSpecSpecLogging](#AppmeshVirtualNodeSpecSpecLogging)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `file` | ***[[]AppmeshVirtualNodeSpecSpecLoggingAccessLogFile](#AppmeshVirtualNodeSpecSpecLoggingAccessLogFile)***| ***(Optional)*** |
## AppmeshVirtualNodeSpecSpecLoggingAccessLogFile

Appears on:[AppmeshVirtualNodeSpecSpecLoggingAccessLog](#AppmeshVirtualNodeSpecSpecLoggingAccessLog)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `path` | ***string***||
## AppmeshVirtualNodeSpecSpecServiceDiscovery

Appears on:[AppmeshVirtualNodeSpecSpec](#AppmeshVirtualNodeSpecSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dns` | ***[[]AppmeshVirtualNodeSpecSpecServiceDiscoveryDns](#AppmeshVirtualNodeSpecSpecServiceDiscoveryDns)***||
## AppmeshVirtualNodeSpecSpecServiceDiscoveryDns

Appears on:[AppmeshVirtualNodeSpecSpecServiceDiscovery](#AppmeshVirtualNodeSpecSpecServiceDiscovery)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `hostname` | ***string***||
## AppmeshVirtualNodeStatus

Appears on:[AppmeshVirtualNode](#AppmeshVirtualNode)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppmeshVirtualNodeSpec](#AppmeshVirtualNodeSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
