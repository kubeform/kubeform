---
title: ProjectService
menu:
  docs_{{ .version }}:
    identifier: projectservice-google.kubeform.com
    name: ProjectService
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ProjectService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ProjectService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectServiceSpec](#projectservicespec)***||
| `status` | ***[ProjectServiceStatus](#projectservicestatus)***||
## ProjectServiceSpec

Appears on:[ProjectService](#projectservice), [ProjectServiceStatus](#projectservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `disableOnDestroy` | ***bool***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `service` | ***string***||
## ProjectServiceStatus

Appears on:[ProjectService](#projectservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectServiceSpec](#projectservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
