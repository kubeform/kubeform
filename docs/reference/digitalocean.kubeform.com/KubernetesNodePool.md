---
title: KubernetesNodePool
menu:
  docs_{{ .version }}:
    identifier: kubernetesnodepool-digitalocean.kubeform.com
    name: KubernetesNodePool
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KubernetesNodePool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `KubernetesNodePool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KubernetesNodePoolSpec](#KubernetesNodePoolSpec)***||
| `status` | ***[KubernetesNodePoolStatus](#KubernetesNodePoolStatus)***||
## KubernetesNodePoolSpec
##### (Appears on:[KubernetesNodePool](#KubernetesNodePool), [KubernetesNodePoolStatus](#KubernetesNodePoolStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `clusterID` | ***string***||
| `name` | ***string***||
| `nodeCount` | ***int***||
| `nodes` | ***[[]KubernetesNodePoolSpecNodes](#KubernetesNodePoolSpecNodes)***| ***(Optional)*** |
| `size` | ***string***||
| `tags` | ***[]string***| ***(Optional)*** |
## KubernetesNodePoolSpecNodes
##### (Appears on:[KubernetesNodePoolSpec](#KubernetesNodePoolSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `createdAt` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `updatedAt` | ***string***| ***(Optional)*** |
## KubernetesNodePoolStatus
##### (Appears on:[KubernetesNodePool](#KubernetesNodePool))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KubernetesNodePoolSpec](#KubernetesNodePoolSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
