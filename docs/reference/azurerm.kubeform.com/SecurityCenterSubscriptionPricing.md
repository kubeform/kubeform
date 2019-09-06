## SecurityCenterSubscriptionPricing
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SecurityCenterSubscriptionPricing` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityCenterSubscriptionPricingSpec](#SecurityCenterSubscriptionPricingSpec)***||
| `status` | ***[SecurityCenterSubscriptionPricingStatus](#SecurityCenterSubscriptionPricingStatus)***||
## SecurityCenterSubscriptionPricingSpec
##### (Appears on:[SecurityCenterSubscriptionPricing](#SecurityCenterSubscriptionPricing), [SecurityCenterSubscriptionPricingStatus](#SecurityCenterSubscriptionPricingStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `tier` | ***string***||
## SecurityCenterSubscriptionPricingStatus
##### (Appears on:[SecurityCenterSubscriptionPricing](#SecurityCenterSubscriptionPricing))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityCenterSubscriptionPricingSpec](#SecurityCenterSubscriptionPricingSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
