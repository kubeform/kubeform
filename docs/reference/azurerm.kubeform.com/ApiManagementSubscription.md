---
title: ApiManagementSubscription
menu:
  docs_v0.0.1:
    identifier: apimanagementsubscription-azurerm.kubeform.com
    name: ApiManagementSubscription
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ApiManagementSubscription
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementSubscription` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementSubscriptionSpec](#ApiManagementSubscriptionSpec)***||
| `status` | ***[ApiManagementSubscriptionStatus](#ApiManagementSubscriptionStatus)***||
## ApiManagementSubscriptionSpec
##### (Appears on:[ApiManagementSubscription](#ApiManagementSubscription), [ApiManagementSubscriptionStatus](#ApiManagementSubscriptionStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `apiManagementName` | ***string***||
| `displayName` | ***string***||
| `productID` | ***string***||
| `resourceGroupName` | ***string***||
| `state` | ***string***| ***(Optional)*** |
| `subscriptionID` | ***string***| ***(Optional)*** |
| `userID` | ***string***||
## ApiManagementSubscriptionStatus
##### (Appears on:[ApiManagementSubscription](#ApiManagementSubscription))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementSubscriptionSpec](#ApiManagementSubscriptionSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_key` | ***string*** ||
| `secondary_key` | ***string*** ||
