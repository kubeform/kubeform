---
title: KubernetesClusterNodePool
menu:
  docs_{{ .version }}:
    identifier: kubernetesclusternodepool-azurerm.kubeform.com
    name: KubernetesClusterNodePool
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KubernetesClusterNodePool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KubernetesClusterNodePool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KubernetesClusterNodePoolSpec](#kubernetesclusternodepoolspec)***||
| `status` | ***[KubernetesClusterNodePoolStatus](#kubernetesclusternodepoolstatus)***||
## KubernetesClusterNodePoolSpec

Appears on:[KubernetesClusterNodePool](#kubernetesclusternodepool), [KubernetesClusterNodePoolStatus](#kubernetesclusternodepoolstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `availabilityZones` | ***[]string***| ***(Optional)*** |
| `enableAutoScaling` | ***bool***| ***(Optional)*** |
| `enableNodePublicIP` | ***bool***| ***(Optional)*** |
| `kubernetesClusterID` | ***string***||
| `maxCount` | ***int64***| ***(Optional)*** |
| `maxPods` | ***int64***| ***(Optional)*** |
| `minCount` | ***int64***| ***(Optional)*** |
| `name` | ***string***||
| `nodeCount` | ***int64***| ***(Optional)*** |
| `nodeTaints` | ***[]string***| ***(Optional)*** |
| `osDiskSizeGb` | ***int64***| ***(Optional)*** |
| `osType` | ***string***| ***(Optional)*** |
| `vmSize` | ***string***||
| `vnetSubnetID` | ***string***| ***(Optional)*** |
## KubernetesClusterNodePoolStatus

Appears on:[KubernetesClusterNodePool](#kubernetesclusternodepool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KubernetesClusterNodePoolSpec](#kubernetesclusternodepoolspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KubernetesClusterNodePoolStatus](#kubernetesclusternodepoolstatus)

---
