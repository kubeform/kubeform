---
title: CognitiveAccount
menu:
  docs_{{ .version }}:
    identifier: cognitiveaccount-azurerm.kubeform.com
    name: CognitiveAccount
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CognitiveAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitiveAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitiveAccountSpec](#cognitiveaccountspec)***||
| `status` | ***[CognitiveAccountStatus](#cognitiveaccountstatus)***||
## CognitiveAccountSpec

Appears on:[CognitiveAccount](#cognitiveaccount), [CognitiveAccountStatus](#cognitiveaccountstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `endpoint` | ***string***| ***(Optional)*** |
| `kind` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***[[]CognitiveAccountSpecSku](#cognitiveaccountspecsku)***| ***(Optional)*** Deprecated|
| `skuName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## CognitiveAccountSpecSku

Appears on:[CognitiveAccountSpec](#cognitiveaccountspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `tier` | ***string***||
## CognitiveAccountStatus

Appears on:[CognitiveAccount](#cognitiveaccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitiveAccountSpec](#cognitiveaccountspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CognitiveAccountStatus](#cognitiveaccountstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_access_key` | ***string*** ||
| `secondary_access_key` | ***string*** ||
