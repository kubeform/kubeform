---
title: DatabaseReplica
menu:
  docs_{{ .version }}:
    identifier: databasereplica-digitalocean.kubeform.com
    name: DatabaseReplica
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DatabaseReplica
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `DatabaseReplica` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatabaseReplicaSpec](#databasereplicaspec)***||
| `status` | ***[DatabaseReplicaStatus](#databasereplicastatus)***||
## DatabaseReplicaSpec

Appears on:[DatabaseReplica](#databasereplica), [DatabaseReplicaStatus](#databasereplicastatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `clusterID` | ***string***||
| `database` | ***string***| ***(Optional)*** |
| `host` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `port` | ***int64***| ***(Optional)*** |
| `privateHost` | ***string***| ***(Optional)*** |
| `privateNetworkUUID` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `size` | ***string***| ***(Optional)*** |
| `tags` | ***[]string***| ***(Optional)*** |
| `user` | ***string***| ***(Optional)*** |
## DatabaseReplicaStatus

Appears on:[DatabaseReplica](#databasereplica)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatabaseReplicaSpec](#databasereplicaspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DatabaseReplicaStatus](#databasereplicastatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
| `private_uri` | ***string*** ||
| `uri` | ***string*** ||
