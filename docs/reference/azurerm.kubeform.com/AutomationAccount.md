---
title: AutomationAccount
menu:
  docs_{{ .version }}:
    identifier: automationaccount-azurerm.kubeform.com
    name: AutomationAccount
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AutomationAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationAccountSpec](#AutomationAccountSpec)***||
| `status` | ***[AutomationAccountStatus](#AutomationAccountStatus)***||
## AutomationAccountSpec

Appears on:[AutomationAccount](#AutomationAccount), [AutomationAccountStatus](#AutomationAccountStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dscPrimaryAccessKey` | ***string***| ***(Optional)*** |
| `dscSecondaryAccessKey` | ***string***| ***(Optional)*** |
| `dscServerEndpoint` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***[[]AutomationAccountSpecSku](#AutomationAccountSpecSku)***| ***(Optional)*** Deprecated|
| `skuName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## AutomationAccountSpecSku

Appears on:[AutomationAccountSpec](#AutomationAccountSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
## AutomationAccountStatus

Appears on:[AutomationAccount](#AutomationAccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationAccountSpec](#AutomationAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
