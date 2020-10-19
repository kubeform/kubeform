---
title: BigqueryDataTransferConfig
menu:
  docs_{{ .version }}:
    identifier: bigquerydatatransferconfig-google.kubeform.com
    name: BigqueryDataTransferConfig
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BigqueryDataTransferConfig
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BigqueryDataTransferConfig` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BigqueryDataTransferConfigSpec](#bigquerydatatransferconfigspec)***||
| `status` | ***[BigqueryDataTransferConfigStatus](#bigquerydatatransferconfigstatus)***||
## BigqueryDataTransferConfigSpec

Appears on:[BigqueryDataTransferConfig](#bigquerydatatransferconfig), [BigqueryDataTransferConfigStatus](#bigquerydatatransferconfigstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dataRefreshWindowDays` | ***int64***| ***(Optional)*** |
| `dataSourceID` | ***string***||
| `destinationDatasetID` | ***string***||
| `disabled` | ***bool***| ***(Optional)*** |
| `displayName` | ***string***||
| `location` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `params` | ***map[string]string***||
| `project` | ***string***| ***(Optional)*** |
| `schedule` | ***string***| ***(Optional)*** |
## BigqueryDataTransferConfigStatus

Appears on:[BigqueryDataTransferConfig](#bigquerydatatransferconfig)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BigqueryDataTransferConfigSpec](#bigquerydatatransferconfigspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BigqueryDataTransferConfigStatus](#bigquerydatatransferconfigstatus)

---
