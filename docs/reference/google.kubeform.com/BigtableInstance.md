---
title: BigtableInstance
menu:
  docs_{{ .version }}:
    identifier: bigtableinstance-google.kubeform.com
    name: BigtableInstance
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BigtableInstance
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BigtableInstance` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BigtableInstanceSpec](#bigtableinstancespec)***||
| `status` | ***[BigtableInstanceStatus](#bigtableinstancestatus)***||
## BigtableInstanceSpec

Appears on:[BigtableInstance](#bigtableinstance), [BigtableInstanceStatus](#bigtableinstancestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cluster` | ***[[]BigtableInstanceSpecCluster](#bigtableinstancespeccluster)***| ***(Optional)*** |
| `clusterID` | ***string***| ***(Optional)*** Deprecated|
| `displayName` | ***string***| ***(Optional)*** |
| `instanceType` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `numNodes` | ***int64***| ***(Optional)*** Deprecated|
| `project` | ***string***| ***(Optional)*** |
| `storageType` | ***string***| ***(Optional)*** Deprecated|
| `zone` | ***string***| ***(Optional)*** Deprecated|
## BigtableInstanceSpecCluster

Appears on:[BigtableInstanceSpec](#bigtableinstancespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clusterID` | ***string***| ***(Optional)*** |
| `numNodes` | ***int64***| ***(Optional)*** |
| `storageType` | ***string***| ***(Optional)*** |
| `zone` | ***string***| ***(Optional)*** |
## BigtableInstanceStatus

Appears on:[BigtableInstance](#bigtableinstance)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BigtableInstanceSpec](#bigtableinstancespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BigtableInstanceStatus](#bigtableinstancestatus)

---
