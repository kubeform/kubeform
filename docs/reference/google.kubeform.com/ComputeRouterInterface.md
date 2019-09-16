---
title: ComputeRouterInterface
menu:
  docs_{{ .version }}:
    identifier: computerouterinterface-google.kubeform.com
    name: ComputeRouterInterface
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeRouterInterface
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRouterInterface` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRouterInterfaceSpec](#ComputeRouterInterfaceSpec)***||
| `status` | ***[ComputeRouterInterfaceStatus](#ComputeRouterInterfaceStatus)***||
## ComputeRouterInterfaceSpec

Appears on:[ComputeRouterInterface](#ComputeRouterInterface), [ComputeRouterInterfaceStatus](#ComputeRouterInterfaceStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ipRange` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `router` | ***string***||
| `vpnTunnel` | ***string***||
## ComputeRouterInterfaceStatus

Appears on:[ComputeRouterInterface](#ComputeRouterInterface)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRouterInterfaceSpec](#ComputeRouterInterfaceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
