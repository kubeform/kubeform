---
title: ComputeRouter
menu:
  docs_{{ .version }}:
    identifier: computerouter-google.kubeform.com
    name: ComputeRouter
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeRouter
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRouter` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRouterSpec](#computerouterspec)***||
| `status` | ***[ComputeRouterStatus](#computerouterstatus)***||
## ComputeRouterSpec

Appears on:[ComputeRouter](#computerouter), [ComputeRouterStatus](#computerouterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `bgp` | ***[[]ComputeRouterSpecBgp](#computerouterspecbgp)***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `network` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeRouterSpecBgp

Appears on:[ComputeRouterSpec](#computerouterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `advertiseMode` | ***string***| ***(Optional)*** |
| `advertisedGroups` | ***[]string***| ***(Optional)*** |
| `advertisedIPRanges` | ***[[]ComputeRouterSpecBgpAdvertisedIPRanges](#computerouterspecbgpadvertisedipranges)***| ***(Optional)*** |
| `asn` | ***int64***||
## ComputeRouterSpecBgpAdvertisedIPRanges

Appears on:[ComputeRouterSpecBgp](#computerouterspecbgp)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `range` | ***string***| ***(Optional)*** |
## ComputeRouterStatus

Appears on:[ComputeRouter](#computerouter)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRouterSpec](#computerouterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeRouterStatus](#computerouterstatus)

---
