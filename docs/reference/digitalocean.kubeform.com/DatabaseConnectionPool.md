---
title: DatabaseConnectionPool
menu:
  docs_{{ .version }}:
    identifier: databaseconnectionpool-digitalocean.kubeform.com
    name: DatabaseConnectionPool
    parent: digitalocean.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DatabaseConnectionPool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `digitalocean.kubeform.com/v1alpha1` |
|    `kind` | string | `DatabaseConnectionPool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DatabaseConnectionPoolSpec](#databaseconnectionpoolspec)***||
| `status` | ***[DatabaseConnectionPoolStatus](#databaseconnectionpoolstatus)***||
## DatabaseConnectionPoolSpec

Appears on:[DatabaseConnectionPool](#databaseconnectionpool), [DatabaseConnectionPoolStatus](#databaseconnectionpoolstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `clusterID` | ***string***||
| `dbName` | ***string***||
| `host` | ***string***| ***(Optional)*** |
| `mode` | ***string***||
| `name` | ***string***||
| `port` | ***int64***| ***(Optional)*** |
| `privateHost` | ***string***| ***(Optional)*** |
| `size` | ***int64***||
| `user` | ***string***||
## DatabaseConnectionPoolStatus

Appears on:[DatabaseConnectionPool](#databaseconnectionpool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DatabaseConnectionPoolSpec](#databaseconnectionpoolspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DatabaseConnectionPoolStatus](#databaseconnectionpoolstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
| `private_uri` | ***string*** ||
| `uri` | ***string*** ||
