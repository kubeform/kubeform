---
title: Loadbalancer
menu:
  docs_{{ .version }}:
    identifier: loadbalancer-digitalocean.kubeform.com
    name: Loadbalancer
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Loadbalancer
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `Loadbalancer` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoadbalancerSpec](#LoadbalancerSpec)***||
| `status` | ***[LoadbalancerStatus](#LoadbalancerStatus)***||
## LoadbalancerSpec
##### (Appears on:[Loadbalancer](#Loadbalancer), [LoadbalancerStatus](#LoadbalancerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `algorithm` | ***string***| ***(Optional)*** |
| `dropletIDS` | ***[]int64***| ***(Optional)*** |
| `dropletTag` | ***string***| ***(Optional)*** |
| `enableProxyProtocol` | ***bool***| ***(Optional)*** |
| `forwardingRule` | ***[[]LoadbalancerSpecForwardingRule](#LoadbalancerSpecForwardingRule)***||
| `healthcheck` | ***[[]LoadbalancerSpecHealthcheck](#LoadbalancerSpecHealthcheck)***| ***(Optional)*** |
| `ip` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `redirectHTTPToHTTPS` | ***bool***| ***(Optional)*** |
| `region` | ***string***||
| `status` | ***string***| ***(Optional)*** |
| `stickySessions` | ***[[]LoadbalancerSpecStickySessions](#LoadbalancerSpecStickySessions)***| ***(Optional)*** |
| `urn` | ***string***| ***(Optional)*** the uniform resource name for the load balancer|
## LoadbalancerSpecForwardingRule
##### (Appears on:[LoadbalancerSpec](#LoadbalancerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateID` | ***string***| ***(Optional)*** |
| `entryPort` | ***int***||
| `entryProtocol` | ***string***||
| `targetPort` | ***int***||
| `targetProtocol` | ***string***||
| `tlsPassthrough` | ***bool***| ***(Optional)*** |
## LoadbalancerSpecHealthcheck
##### (Appears on:[LoadbalancerSpec](#LoadbalancerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `checkIntervalSeconds` | ***int***| ***(Optional)*** |
| `healthyThreshold` | ***int***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***int***||
| `protocol` | ***string***||
| `responseTimeoutSeconds` | ***int***| ***(Optional)*** |
| `unhealthyThreshold` | ***int***| ***(Optional)*** |
## LoadbalancerSpecStickySessions
##### (Appears on:[LoadbalancerSpec](#LoadbalancerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cookieName` | ***string***| ***(Optional)*** |
| `cookieTtlSeconds` | ***int***| ***(Optional)*** |
| `type` | ***string***| ***(Optional)*** |
## LoadbalancerStatus
##### (Appears on:[Loadbalancer](#Loadbalancer))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoadbalancerSpec](#LoadbalancerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
