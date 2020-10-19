---
title: ProjectResources
menu:
  docs_{{ .version }}:
    identifier: projectresources-digitalocean.kubeform.com
    name: ProjectResources
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ProjectResources
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `ProjectResources` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectResourcesSpec](#projectresourcesspec)***||
| `status` | ***[ProjectResourcesStatus](#projectresourcesstatus)***||
## Phase(`string` alias)

Appears on:[ProjectResourcesStatus](#projectresourcesstatus)

## ProjectResourcesSpec

Appears on:[ProjectResources](#projectresources), [ProjectResourcesStatus](#projectresourcesstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `project` | ***string***|project ID|
| `resources` | ***[]string***|the resources associated with the project|
## ProjectResourcesStatus

Appears on:[ProjectResources](#projectresources)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectResourcesSpec](#projectresourcesspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
