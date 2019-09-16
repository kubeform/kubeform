---
title: Route53Record
menu:
  docs_{{ .version }}:
    identifier: route53record-aws.kubeform.com
    name: Route53Record
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Route53Record
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Route53Record` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53RecordSpec](#Route53RecordSpec)***||
| `status` | ***[Route53RecordStatus](#Route53RecordStatus)***||
## Route53RecordSpec

Appears on:[Route53Record](#Route53Record), [Route53RecordStatus](#Route53RecordStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alias` | ***[[]Route53RecordSpecAlias](#Route53RecordSpecAlias)***| ***(Optional)*** |
| `allowOverwrite` | ***bool***| ***(Optional)*** |
| `failoverRoutingPolicy` | ***[[]Route53RecordSpecFailoverRoutingPolicy](#Route53RecordSpecFailoverRoutingPolicy)***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** |
| `geolocationRoutingPolicy` | ***[[]Route53RecordSpecGeolocationRoutingPolicy](#Route53RecordSpecGeolocationRoutingPolicy)***| ***(Optional)*** |
| `healthCheckID` | ***string***| ***(Optional)*** |
| `latencyRoutingPolicy` | ***[[]Route53RecordSpecLatencyRoutingPolicy](#Route53RecordSpecLatencyRoutingPolicy)***| ***(Optional)*** |
| `multivalueAnswerRoutingPolicy` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `records` | ***[]string***| ***(Optional)*** |
| `setIdentifier` | ***string***| ***(Optional)*** |
| `ttl` | ***int***| ***(Optional)*** |
| `type` | ***string***||
| `weightedRoutingPolicy` | ***[[]Route53RecordSpecWeightedRoutingPolicy](#Route53RecordSpecWeightedRoutingPolicy)***| ***(Optional)*** |
| `zoneID` | ***string***||
## Route53RecordSpecAlias

Appears on:[Route53RecordSpec](#Route53RecordSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `evaluateTargetHealth` | ***bool***||
| `name` | ***string***||
| `zoneID` | ***string***||
## Route53RecordSpecFailoverRoutingPolicy

Appears on:[Route53RecordSpec](#Route53RecordSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## Route53RecordSpecGeolocationRoutingPolicy

Appears on:[Route53RecordSpec](#Route53RecordSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `continent` | ***string***| ***(Optional)*** |
| `country` | ***string***| ***(Optional)*** |
| `subdivision` | ***string***| ***(Optional)*** |
## Route53RecordSpecLatencyRoutingPolicy

Appears on:[Route53RecordSpec](#Route53RecordSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `region` | ***string***||
## Route53RecordSpecWeightedRoutingPolicy

Appears on:[Route53RecordSpec](#Route53RecordSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `weight` | ***int***||
## Route53RecordStatus

Appears on:[Route53Record](#Route53Record)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53RecordSpec](#Route53RecordSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
