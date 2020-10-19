---
title: Route53Zone
menu:
  docs_{{ .version }}:
    identifier: route53zone-aws.kubeform.com
    name: Route53Zone
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Route53Zone
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Route53Zone` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53ZoneSpec](#route53zonespec)***||
| `status` | ***[Route53ZoneStatus](#route53zonestatus)***||
## Phase(`string` alias)

Appears on:[Route53ZoneStatus](#route53zonestatus)

## Route53ZoneSpec

Appears on:[Route53Zone](#route53zone), [Route53ZoneStatus](#route53zonestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `comment` | ***string***| ***(Optional)*** |
| `delegationSetID` | ***string***| ***(Optional)*** |
| `forceDestroy` | ***bool***| ***(Optional)*** |
| `name` | ***string***||
| `nameServers` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpc` | ***[[]Route53ZoneSpecVpc](#route53zonespecvpc)***| ***(Optional)*** |
| `zoneID` | ***string***| ***(Optional)*** |
## Route53ZoneSpecVpc

Appears on:[Route53ZoneSpec](#route53zonespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `vpcID` | ***string***||
| `vpcRegion` | ***string***| ***(Optional)*** |
## Route53ZoneStatus

Appears on:[Route53Zone](#route53zone)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53ZoneSpec](#route53zonespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
