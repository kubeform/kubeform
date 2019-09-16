---
title: BigqueryDataset
menu:
  docs_{{ .version }}:
    identifier: bigquerydataset-google.kubeform.com
    name: BigqueryDataset
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BigqueryDataset
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BigqueryDataset` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BigqueryDatasetSpec](#BigqueryDatasetSpec)***||
| `status` | ***[BigqueryDatasetStatus](#BigqueryDatasetStatus)***||
## BigqueryDatasetSpec

Appears on:[BigqueryDataset](#BigqueryDataset), [BigqueryDatasetStatus](#BigqueryDatasetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `access` | ***[[]BigqueryDatasetSpecAccess](#BigqueryDatasetSpecAccess)***| ***(Optional)*** |
| `creationTime` | ***int***| ***(Optional)*** |
| `datasetID` | ***string***||
| `defaultTableExpirationMs` | ***int***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `etag` | ***string***| ***(Optional)*** |
| `friendlyName` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** |
| `lastModifiedTime` | ***int***| ***(Optional)*** |
| `location` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## BigqueryDatasetSpecAccess

Appears on:[BigqueryDatasetSpec](#BigqueryDatasetSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `domain` | ***string***| ***(Optional)*** |
| `groupByEmail` | ***string***| ***(Optional)*** |
| `role` | ***string***| ***(Optional)*** |
| `specialGroup` | ***string***| ***(Optional)*** |
| `userByEmail` | ***string***| ***(Optional)*** |
| `view` | ***[[]BigqueryDatasetSpecAccessView](#BigqueryDatasetSpecAccessView)***| ***(Optional)*** |
## BigqueryDatasetSpecAccessView

Appears on:[BigqueryDatasetSpecAccess](#BigqueryDatasetSpecAccess)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `datasetID` | ***string***||
| `projectID` | ***string***||
| `tableID` | ***string***||
## BigqueryDatasetStatus

Appears on:[BigqueryDataset](#BigqueryDataset)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BigqueryDatasetSpec](#BigqueryDatasetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
