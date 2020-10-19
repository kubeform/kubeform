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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KubernetesNodePoolSpec](#kubernetesnodepoolspec)***||
| `status` | ***[KubernetesNodePoolStatus](#kubernetesnodepoolstatus)***||
## KubernetesNodePoolSpec

Appears on:[KubernetesNodePool](#kubernetesnodepool), [KubernetesNodePoolStatus](#kubernetesnodepoolstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `actualNodeCount` | ***int64***| ***(Optional)*** |
| `autoScale` | ***bool***| ***(Optional)*** |
| `clusterID` | ***string***||
| `labels` | ***map[string]string***| ***(Optional)*** |
| `maxNodes` | ***int64***| ***(Optional)*** |
| `minNodes` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `nodeCount` | ***int64***| ***(Optional)*** |
| `nodes` | ***[[]KubernetesNodePoolSpecNodes](#kubernetesnodepoolspecnodes)***| ***(Optional)*** |
| `size` | ***string***||
| `tags` | ***[]string***| ***(Optional)*** |
## KubernetesNodePoolSpecNodes

Appears on:[KubernetesNodePoolSpec](#kubernetesnodepoolspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `createdAt` | ***string***| ***(Optional)*** |
| `dropletID` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `updatedAt` | ***string***| ***(Optional)*** |
## KubernetesNodePoolStatus

Appears on:[KubernetesNodePool](#kubernetesnodepool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KubernetesNodePoolSpec](#kubernetesnodepoolspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KubernetesNodePoolStatus](#kubernetesnodepoolstatus)

---
