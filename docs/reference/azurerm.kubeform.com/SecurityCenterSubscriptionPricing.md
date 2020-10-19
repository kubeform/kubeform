---
title: SecurityCenterSubscriptionPricing
menu:
  docs_{{ .version }}:
    identifier: securitycentersubscriptionpricing-azurerm.kubeform.com
    name: SecurityCenterSubscriptionPricing
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SecurityCenterSubscriptionPricing
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SecurityCenterSubscriptionPricing` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityCenterSubscriptionPricingSpec](#securitycentersubscriptionpricingspec)***||
| `status` | ***[SecurityCenterSubscriptionPricingStatus](#securitycentersubscriptionpricingstatus)***||
## Phase(`string` alias)

Appears on:[SecurityCenterSubscriptionPricingStatus](#securitycentersubscriptionpricingstatus)

## SecurityCenterSubscriptionPricingSpec

Appears on:[SecurityCenterSubscriptionPricing](#securitycentersubscriptionpricing), [SecurityCenterSubscriptionPricingStatus](#securitycentersubscriptionpricingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `tier` | ***string***||
## SecurityCenterSubscriptionPricingStatus

Appears on:[SecurityCenterSubscriptionPricing](#securitycentersubscriptionpricing)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityCenterSubscriptionPricingSpec](#securitycentersubscriptionpricingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
