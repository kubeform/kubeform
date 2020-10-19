---
title: FirestoreIndex
menu:
  docs_{{ .version }}:
    identifier: firestoreindex-google.kubeform.com
    name: FirestoreIndex
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FirestoreIndex
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `FirestoreIndex` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FirestoreIndexSpec](#firestoreindexspec)***||
| `status` | ***[FirestoreIndexStatus](#firestoreindexstatus)***||
## FirestoreIndexSpec

Appears on:[FirestoreIndex](#firestoreindex), [FirestoreIndexStatus](#firestoreindexstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `collection` | ***string***||
| `database` | ***string***| ***(Optional)*** |
| `fields` | ***[[]FirestoreIndexSpecFields](#firestoreindexspecfields)***||
| `name` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `queryScope` | ***string***| ***(Optional)*** |
## FirestoreIndexSpecFields

Appears on:[FirestoreIndexSpec](#firestoreindexspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `arrayConfig` | ***string***| ***(Optional)*** |
| `fieldPath` | ***string***| ***(Optional)*** |
| `order` | ***string***| ***(Optional)*** |
## FirestoreIndexStatus

Appears on:[FirestoreIndex](#firestoreindex)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FirestoreIndexSpec](#firestoreindexspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[FirestoreIndexStatus](#firestoreindexstatus)

---
