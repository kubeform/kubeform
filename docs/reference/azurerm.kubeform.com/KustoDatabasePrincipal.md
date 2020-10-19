---
title: KustoDatabasePrincipal
menu:
  docs_{{ .version }}:
    identifier: kustodatabaseprincipal-azurerm.kubeform.com
    name: KustoDatabasePrincipal
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KustoDatabasePrincipal
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `KustoDatabasePrincipal` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KustoDatabasePrincipalSpec](#kustodatabaseprincipalspec)***||
| `status` | ***[KustoDatabasePrincipalStatus](#kustodatabaseprincipalstatus)***||
## KustoDatabasePrincipalSpec

Appears on:[KustoDatabasePrincipal](#kustodatabaseprincipal), [KustoDatabasePrincipalStatus](#kustodatabaseprincipalstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appID` | ***string***| ***(Optional)*** |
| `clientID` | ***string***||
| `clusterName` | ***string***||
| `databaseName` | ***string***||
| `email` | ***string***| ***(Optional)*** |
| `fullyQualifiedName` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `objectID` | ***string***||
| `resourceGroupName` | ***string***||
| `role` | ***string***||
| `type` | ***string***||
## KustoDatabasePrincipalStatus

Appears on:[KustoDatabasePrincipal](#kustodatabaseprincipal)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KustoDatabasePrincipalSpec](#kustodatabaseprincipalspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KustoDatabasePrincipalStatus](#kustodatabaseprincipalstatus)

---
