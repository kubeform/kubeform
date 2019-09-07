---
title: KubernetesCluster
menu:
  docs_v0.0.1:
    identifier: kubernetescluster-digitalocean.kubeform.com
    name: KubernetesCluster
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## KubernetesCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `KubernetesCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KubernetesClusterSpec](#KubernetesClusterSpec)***||
| `status` | ***[KubernetesClusterStatus](#KubernetesClusterStatus)***||
## KubernetesClusterSpec
##### (Appears on:[KubernetesCluster](#KubernetesCluster), [KubernetesClusterStatus](#KubernetesClusterStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `clusterSubnet` | ***string***| ***(Optional)*** |
| `createdAt` | ***string***| ***(Optional)*** |
| `endpoint` | ***string***| ***(Optional)*** |
| `ipv4Address` | ***string***| ***(Optional)*** |
| `kubeConfig` | ***[[]KubernetesClusterSpecKubeConfig](#KubernetesClusterSpecKubeConfig)***| ***(Optional)*** |
| `name` | ***string***||
| `nodePool` | ***[[]KubernetesClusterSpecNodePool](#KubernetesClusterSpecNodePool)***||
| `region` | ***string***||
| `serviceSubnet` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `updatedAt` | ***string***| ***(Optional)*** |
| `version` | ***string***||
## KubernetesClusterSpecKubeConfig
##### (Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientCertificate` | ***string***| ***(Optional)*** |
| `clientKey` | ***string***| ***(Optional)*** |
| `clusterCaCertificate` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `rawConfig` | ***string***| ***(Optional)*** |
## KubernetesClusterSpecNodePool
##### (Appears on:[KubernetesClusterSpec](#KubernetesClusterSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `nodeCount` | ***int***||
| `nodes` | ***[[]KubernetesClusterSpecNodePoolNodes](#KubernetesClusterSpecNodePoolNodes)***| ***(Optional)*** |
| `size` | ***string***||
| `tags` | ***[]string***| ***(Optional)*** |
## KubernetesClusterSpecNodePoolNodes
##### (Appears on:[KubernetesClusterSpecNodePool](#KubernetesClusterSpecNodePool))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `createdAt` | ***string***| ***(Optional)*** |
| `ID` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `updatedAt` | ***string***| ***(Optional)*** |
## KubernetesClusterStatus
##### (Appears on:[KubernetesCluster](#KubernetesCluster))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KubernetesClusterSpec](#KubernetesClusterSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
