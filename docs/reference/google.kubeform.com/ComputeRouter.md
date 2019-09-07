---
title: ComputeRouter
menu:
  docs_v0.0.1:
    identifier: computerouter-google.kubeform.com
    name: ComputeRouter
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ComputeRouter
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRouter` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRouterSpec](#ComputeRouterSpec)***||
| `status` | ***[ComputeRouterStatus](#ComputeRouterStatus)***||
## ComputeRouterSpec
##### (Appears on:[ComputeRouter](#ComputeRouter), [ComputeRouterStatus](#ComputeRouterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bgp` | ***[[]ComputeRouterSpecBgp](#ComputeRouterSpecBgp)***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `network` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeRouterSpecBgp
##### (Appears on:[ComputeRouterSpec](#ComputeRouterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `advertiseMode` | ***string***| ***(Optional)*** |
| `advertisedGroups` | ***[]string***| ***(Optional)*** |
| `advertisedIPRanges` | ***[[]ComputeRouterSpecBgpAdvertisedIPRanges](#ComputeRouterSpecBgpAdvertisedIPRanges)***| ***(Optional)*** |
| `asn` | ***int***||
## ComputeRouterSpecBgpAdvertisedIPRanges
##### (Appears on:[ComputeRouterSpecBgp](#ComputeRouterSpecBgp))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `range` | ***string***| ***(Optional)*** |
## ComputeRouterStatus
##### (Appears on:[ComputeRouter](#ComputeRouter))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRouterSpec](#ComputeRouterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
