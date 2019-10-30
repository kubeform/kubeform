---
title: ComputeRegionBackendService
menu:
  docs_{{ .version }}:
    identifier: computeregionbackendservice-google.kubeform.com
    name: ComputeRegionBackendService
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeRegionBackendService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeRegionBackendService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeRegionBackendServiceSpec](#computeregionbackendservicespec)***||
| `status` | ***[ComputeRegionBackendServiceStatus](#computeregionbackendservicestatus)***||
## ComputeRegionBackendServiceSpec

Appears on:[ComputeRegionBackendService](#computeregionbackendservice), [ComputeRegionBackendServiceStatus](#computeregionbackendservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backend` | ***[[]ComputeRegionBackendServiceSpecBackend](#computeregionbackendservicespecbackend)***| ***(Optional)*** |
| `connectionDrainingTimeoutSec` | ***int***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `healthChecks` | ***[]string***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `sessionAffinity` | ***string***| ***(Optional)*** |
| `timeoutSec` | ***int***| ***(Optional)*** |
## ComputeRegionBackendServiceSpecBackend

Appears on:[ComputeRegionBackendServiceSpec](#computeregionbackendservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `description` | ***string***| ***(Optional)*** |
| `group` | ***string***| ***(Optional)*** |
## ComputeRegionBackendServiceStatus

Appears on:[ComputeRegionBackendService](#computeregionbackendservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeRegionBackendServiceSpec](#computeregionbackendservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeRegionBackendServiceStatus](#computeregionbackendservicestatus)

---
