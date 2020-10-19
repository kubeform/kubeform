---
title: ComputeRoute
menu:
  docs_{{ .version }}:
    identifier: computeroute-google.kubeform.com
    name: ComputeRoute
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeRoute
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRoute` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRouteSpec](#computeroutespec)***||
| `status` | ***[ComputeRouteStatus](#computeroutestatus)***||
## ComputeRouteSpec

Appears on:[ComputeRoute](#computeroute), [ComputeRouteStatus](#computeroutestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `destRange` | ***string***||
| `name` | ***string***||
| `network` | ***string***||
| `nextHopGateway` | ***string***| ***(Optional)*** |
| `nextHopInstance` | ***string***| ***(Optional)*** |
| `nextHopInstanceZone` | ***string***| ***(Optional)*** |
| `nextHopIP` | ***string***| ***(Optional)*** |
| `nextHopNetwork` | ***string***| ***(Optional)*** |
| `nextHopVPNTunnel` | ***string***| ***(Optional)*** |
| `priority` | ***int64***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
## ComputeRouteStatus

Appears on:[ComputeRoute](#computeroute)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRouteSpec](#computeroutespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeRouteStatus](#computeroutestatus)

---
