---
title: ContainerRegistry
menu:
  docs_{{ .version }}:
    identifier: containerregistry-digitalocean.kubeform.com
    name: ContainerRegistry
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ContainerRegistry
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerRegistry` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerRegistrySpec](#containerregistryspec)***||
| `status` | ***[ContainerRegistryStatus](#containerregistrystatus)***||
## ContainerRegistrySpec

Appears on:[ContainerRegistry](#containerregistry), [ContainerRegistryStatus](#containerregistrystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `endpoint` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `serverURL` | ***string***| ***(Optional)*** |
## ContainerRegistryStatus

Appears on:[ContainerRegistry](#containerregistry)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerRegistrySpec](#containerregistryspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ContainerRegistryStatus](#containerregistrystatus)

---
