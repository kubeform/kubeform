---
title: DatabaseCluster
menu:
  docs_{{ .version }}:
    identifier: databasecluster-digitalocean.kubeform.com
    name: DatabaseCluster
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DatabaseCluster
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `DatabaseCluster` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatabaseClusterSpec](#databaseclusterspec)***||
| `status` | ***[DatabaseClusterStatus](#databaseclusterstatus)***||
## DatabaseClusterSpec

Appears on:[DatabaseCluster](#databasecluster), [DatabaseClusterStatus](#databaseclusterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `database` | ***string***| ***(Optional)*** |
| `engine` | ***string***||
| `evictionPolicy` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `maintenanceWindow` | ***[[]DatabaseClusterSpecMaintenanceWindow](#databaseclusterspecmaintenancewindow)***| ***(Optional)*** |
| `name` | ***string***||
| `nodeCount` | ***int64***||
| `port` | ***int64***| ***(Optional)*** |
| `privateHost` | ***string***| ***(Optional)*** |
| `privateNetworkUUID` | ***string***| ***(Optional)*** |
| `region` | ***string***||
| `size` | ***string***||
| `sqlMode` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `urn` | ***string***| ***(Optional)*** |
| `user` | ***string***| ***(Optional)*** |
| `version` | ***string***| ***(Optional)*** |
## DatabaseClusterSpecMaintenanceWindow

Appears on:[DatabaseClusterSpec](#databaseclusterspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `day` | ***string***||
| `hour` | ***string***||
## DatabaseClusterStatus

Appears on:[DatabaseCluster](#databasecluster)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatabaseClusterSpec](#databaseclusterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DatabaseClusterStatus](#databaseclusterstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
| `private_uri` | ***string*** ||
| `uri` | ***string*** ||
