---
title: PrivateLinkService
menu:
  docs_{{ .version }}:
    identifier: privatelinkservice-azurerm.kubeform.com
    name: PrivateLinkService
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PrivateLinkService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PrivateLinkService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PrivateLinkServiceSpec](#privatelinkservicespec)***||
| `status` | ***[PrivateLinkServiceStatus](#privatelinkservicestatus)***||
## Phase(`string` alias)

Appears on:[PrivateLinkServiceStatus](#privatelinkservicestatus)

## PrivateLinkServiceSpec

Appears on:[PrivateLinkService](#privatelinkservice), [PrivateLinkServiceStatus](#privatelinkservicestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `alias` | ***string***| ***(Optional)*** |
| `autoApprovalSubscriptionIDS` | ***[]string***| ***(Optional)*** |
| `enableProxyProtocol` | ***bool***| ***(Optional)*** |
| `loadBalancerFrontendIPConfigurationIDS` | ***[]string***||
| `location` | ***string***||
| `name` | ***string***||
| `natIPConfiguration` | ***[[]PrivateLinkServiceSpecNatIPConfiguration](#privatelinkservicespecnatipconfiguration)***||
| `networkInterfaceIDS` | ***[]string***| ***(Optional)*** Deprecated|
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `visibilitySubscriptionIDS` | ***[]string***| ***(Optional)*** |
## PrivateLinkServiceSpecNatIPConfiguration

Appears on:[PrivateLinkServiceSpec](#privatelinkservicespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `primary` | ***bool***||
| `privateIPAddress` | ***string***| ***(Optional)*** |
| `privateIPAddressVersion` | ***string***| ***(Optional)*** |
| `subnetID` | ***string***||
## PrivateLinkServiceStatus

Appears on:[PrivateLinkService](#privatelinkservice)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PrivateLinkServiceSpec](#privatelinkservicespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
