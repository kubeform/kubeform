---
title: KubernetesCluster
menu:
  docs_{{ .version }}:
    identifier: kubernetescluster-digitalocean.kubeform.com
    name: KubernetesCluster
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KubernetesCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `KubernetesCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KubernetesClusterSpec](#kubernetesclusterspec)***||
| `status` | ***[KubernetesClusterStatus](#kubernetesclusterstatus)***||
## KubernetesClusterSpec

Appears on:[KubernetesCluster](#kubernetescluster), [KubernetesClusterStatus](#kubernetesclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `clusterSubnet` | ***string***| ***(Optional)*** |
| `createdAt` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `ipv4Address` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `nodePool` | ***[[]KubernetesClusterSpecNodePool](#kubernetesclusterspecnodepool)***||
| `region` | ***string***||
| `serviceSubnet` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `updatedAt` | ***string***| ***(Optional)*** |
| `version` | ***string***||
| `vpcUUID` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecNodePool

Appears on:[KubernetesClusterSpec](#kubernetesclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `actualNodeCount` | ***int64***| ***(Optional)*** |
| `autoScale` | ***bool***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `maxNodes` | ***int64***| ***(Optional)*** |
| `minNodes` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `nodeCount` | ***int64***| ***(Optional)*** |
| `nodes` | ***[[]KubernetesClusterSpecNodePoolNodes](#kubernetesclusterspecnodepoolnodes)***| ***(Optional)*** |
| `size` | ***string***||
| `tags` | ***[]string***| ***(Optional)*** |
## KubernetesClusterSpecNodePoolNodes

Appears on:[KubernetesClusterSpecNodePool](#kubernetesclusterspecnodepool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `createdAt` | ***string***| ***(Optional)*** |
| `dropletID` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `updatedAt` | ***string***| ***(Optional)*** |
## KubernetesClusterStatus

Appears on:[KubernetesCluster](#kubernetescluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KubernetesClusterSpec](#kubernetesclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KubernetesClusterStatus](#kubernetesclusterstatus)

---
