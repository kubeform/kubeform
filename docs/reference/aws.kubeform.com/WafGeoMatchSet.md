---
title: WafGeoMatchSet
menu:
  docs_{{ .version }}:
    identifier: wafgeomatchset-aws.kubeform.com
    name: WafGeoMatchSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafGeoMatchSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafGeoMatchSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafGeoMatchSetSpec](#WafGeoMatchSetSpec)***||
| `status` | ***[WafGeoMatchSetStatus](#WafGeoMatchSetStatus)***||
## WafGeoMatchSetSpec

Appears on:[WafGeoMatchSet](#WafGeoMatchSet), [WafGeoMatchSetStatus](#WafGeoMatchSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `geoMatchConstraint` | ***[[]WafGeoMatchSetSpecGeoMatchConstraint](#WafGeoMatchSetSpecGeoMatchConstraint)***| ***(Optional)*** |
| `name` | ***string***||
## WafGeoMatchSetSpecGeoMatchConstraint

Appears on:[WafGeoMatchSetSpec](#WafGeoMatchSetSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `type` | ***string***||
| `value` | ***string***||
## WafGeoMatchSetStatus

Appears on:[WafGeoMatchSet](#WafGeoMatchSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafGeoMatchSetSpec](#WafGeoMatchSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
