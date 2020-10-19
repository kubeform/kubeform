---
title: LkeCluster
menu:
  docs_{{ .version }}:
    identifier: lkecluster-linode.kubeform.com
    name: LkeCluster
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LkeCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `LkeCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LkeClusterSpec](#lkeclusterspec)***||
| `status` | ***[LkeClusterStatus](#lkeclusterstatus)***||
## LkeClusterSpec

Appears on:[LkeCluster](#lkecluster), [LkeClusterStatus](#lkeclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `apiEndpoints` | ***[]string***| ***(Optional)*** The API endpoints for the cluster.|
| `k8sVersion` | ***string***|The desired Kubernetes version for this Kubernetes cluster in the format of <major>.<minor>. The latest supported patch version will be deployed.|
| `label` | ***string***|The unique label for the cluster.|
| `pool` | ***[[]LkeClusterSpecPool](#lkeclusterspecpool)***|A node pool in the cluster.|
| `region` | ***string***|This cluster's location.|
| `status` | ***string***| ***(Optional)*** The status of the cluster.|
| `tags` | ***[]string***| ***(Optional)*** An array of tags applied to this object. Tags are for organizational purposes only.|
## LkeClusterSpecPool

Appears on:[LkeClusterSpec](#lkeclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `count` | ***int64***|The number of nodes in the Node Pool.|
| `ID` | ***int64***| ***(Optional)*** The ID of the Node Pool.|
| `nodes` | ***[[]LkeClusterSpecPoolNodes](#lkeclusterspecpoolnodes)***| ***(Optional)*** The nodes in the node pool.|
| `type` | ***string***|A Linode Type for all of the nodes in the Node Pool.|
## LkeClusterSpecPoolNodes

Appears on:[LkeClusterSpecPool](#lkeclusterspecpool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `ID` | ***string***| ***(Optional)*** The ID of the node.|
| `instanceID` | ***int64***| ***(Optional)*** The ID of the underlying Linode instance.|
| `status` | ***string***| ***(Optional)*** The status of the node.|
## LkeClusterStatus

Appears on:[LkeCluster](#lkecluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LkeClusterSpec](#lkeclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LkeClusterStatus](#lkeclusterstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `kubeconfig` | ***string*** |The Base64-encoded Kubeconfig for the cluster.|
