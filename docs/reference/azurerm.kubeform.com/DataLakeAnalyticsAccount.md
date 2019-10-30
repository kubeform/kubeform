---
title: DataLakeAnalyticsAccount
menu:
  docs_{{ .version }}:
    identifier: datalakeanalyticsaccount-azurerm.kubeform.com
    name: DataLakeAnalyticsAccount
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataLakeAnalyticsAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DataLakeAnalyticsAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataLakeAnalyticsAccountSpec](#datalakeanalyticsaccountspec)***||
| `status` | ***[DataLakeAnalyticsAccountStatus](#datalakeanalyticsaccountstatus)***||
## DataLakeAnalyticsAccountSpec

Appears on:[DataLakeAnalyticsAccount](#datalakeanalyticsaccount), [DataLakeAnalyticsAccountStatus](#datalakeanalyticsaccountstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `defaultStoreAccountName` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `tier` | ***string***| ***(Optional)*** |
## DataLakeAnalyticsAccountStatus

Appears on:[DataLakeAnalyticsAccount](#datalakeanalyticsaccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataLakeAnalyticsAccountSpec](#datalakeanalyticsaccountspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DataLakeAnalyticsAccountStatus](#datalakeanalyticsaccountstatus)

---
