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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53RecordSpec](#route53recordspec)***||
| `status` | ***[Route53RecordStatus](#route53recordstatus)***||
## Phase(`string` alias)

Appears on:[Route53RecordStatus](#route53recordstatus)

## Route53RecordSpec

Appears on:[Route53Record](#route53record), [Route53RecordStatus](#route53recordstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alias` | ***[[]Route53RecordSpecAlias](#route53recordspecalias)***| ***(Optional)*** |
| `allowOverwrite` | ***bool***| ***(Optional)*** |
| `failoverRoutingPolicy` | ***[[]Route53RecordSpecFailoverRoutingPolicy](#route53recordspecfailoverroutingpolicy)***| ***(Optional)*** |
| `fqdn` | ***string***| ***(Optional)*** |
| `geolocationRoutingPolicy` | ***[[]Route53RecordSpecGeolocationRoutingPolicy](#route53recordspecgeolocationroutingpolicy)***| ***(Optional)*** |
| `healthCheckID` | ***string***| ***(Optional)*** |
| `latencyRoutingPolicy` | ***[[]Route53RecordSpecLatencyRoutingPolicy](#route53recordspeclatencyroutingpolicy)***| ***(Optional)*** |
| `multivalueAnswerRoutingPolicy` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `records` | ***[]string***| ***(Optional)*** |
| `setIdentifier` | ***string***| ***(Optional)*** |
| `ttl` | ***int64***| ***(Optional)*** |
| `type` | ***string***||
| `weightedRoutingPolicy` | ***[[]Route53RecordSpecWeightedRoutingPolicy](#route53recordspecweightedroutingpolicy)***| ***(Optional)*** |
| `zoneID` | ***string***||
## Route53RecordSpecAlias

Appears on:[Route53RecordSpec](#route53recordspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `evaluateTargetHealth` | ***bool***||
| `name` | ***string***||
| `zoneID` | ***string***||
## Route53RecordSpecFailoverRoutingPolicy

Appears on:[Route53RecordSpec](#route53recordspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
## Route53RecordSpecGeolocationRoutingPolicy

Appears on:[Route53RecordSpec](#route53recordspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `continent` | ***string***| ***(Optional)*** |
| `country` | ***string***| ***(Optional)*** |
| `subdivision` | ***string***| ***(Optional)*** |
## Route53RecordSpecLatencyRoutingPolicy

Appears on:[Route53RecordSpec](#route53recordspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `region` | ***string***||
## Route53RecordSpecWeightedRoutingPolicy

Appears on:[Route53RecordSpec](#route53recordspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `weight` | ***int64***||
## Route53RecordStatus

Appears on:[Route53Record](#route53record)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53RecordSpec](#route53recordspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
