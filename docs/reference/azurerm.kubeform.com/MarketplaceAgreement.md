---
title: MarketplaceAgreement
menu:
  docs_{{ .version }}:
    identifier: marketplaceagreement-azurerm.kubeform.com
    name: MarketplaceAgreement
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## MarketplaceAgreement
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `MarketplaceAgreement` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[MarketplaceAgreementSpec](#marketplaceagreementspec)***||
| `status` | ***[MarketplaceAgreementStatus](#marketplaceagreementstatus)***||
## MarketplaceAgreementSpec

Appears on:[MarketplaceAgreement](#marketplaceagreement), [MarketplaceAgreementStatus](#marketplaceagreementstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `licenseTextLink` | ***string***| ***(Optional)*** |
| `offer` | ***string***||
| `plan` | ***string***||
| `privacyPolicyLink` | ***string***| ***(Optional)*** |
| `publisher` | ***string***||
## MarketplaceAgreementStatus

Appears on:[MarketplaceAgreement](#marketplaceagreement)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[MarketplaceAgreementSpec](#marketplaceagreementspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[MarketplaceAgreementStatus](#marketplaceagreementstatus)

---
