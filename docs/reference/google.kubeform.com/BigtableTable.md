---
title: BigtableTable
menu:
  docs_{{ .version }}:
    identifier: bigtabletable-google.kubeform.com
    name: BigtableTable
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BigtableTable
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BigtableTable` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BigtableTableSpec](#bigtabletablespec)***||
| `status` | ***[BigtableTableStatus](#bigtabletablestatus)***||
## BigtableTableSpec

Appears on:[BigtableTable](#bigtabletable), [BigtableTableStatus](#bigtabletablestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `columnFamily` | ***[[]BigtableTableSpecColumnFamily](#bigtabletablespeccolumnfamily)***| ***(Optional)*** |
| `instanceName` | ***string***||
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `splitKeys` | ***[]string***| ***(Optional)*** |
## BigtableTableSpecColumnFamily

Appears on:[BigtableTableSpec](#bigtabletablespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `family` | ***string***||
## BigtableTableStatus

Appears on:[BigtableTable](#bigtabletable)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BigtableTableSpec](#bigtabletablespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BigtableTableStatus](#bigtabletablestatus)

---
