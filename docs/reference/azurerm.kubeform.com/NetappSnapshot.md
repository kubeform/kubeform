---
title: NetappSnapshot
menu:
  docs_{{ .version }}:
    identifier: netappsnapshot-azurerm.kubeform.com
    name: NetappSnapshot
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetappSnapshot
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetappSnapshot` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetappSnapshotSpec](#netappsnapshotspec)***||
| `status` | ***[NetappSnapshotStatus](#netappsnapshotstatus)***||
## NetappSnapshotSpec

Appears on:[NetappSnapshot](#netappsnapshot), [NetappSnapshotStatus](#netappsnapshotstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountName` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `poolName` | ***string***||
| `resourceGroupName` | ***string***||
| `volumeName` | ***string***||
## NetappSnapshotStatus

Appears on:[NetappSnapshot](#netappsnapshot)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetappSnapshotSpec](#netappsnapshotspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NetappSnapshotStatus](#netappsnapshotstatus)

---
