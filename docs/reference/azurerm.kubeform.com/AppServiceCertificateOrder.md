---
title: AppServiceCertificateOrder
menu:
  docs_{{ .version }}:
    identifier: appservicecertificateorder-azurerm.kubeform.com
    name: AppServiceCertificateOrder
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppServiceCertificateOrder
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AppServiceCertificateOrder` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppServiceCertificateOrderSpec](#appservicecertificateorderspec)***||
| `status` | ***[AppServiceCertificateOrderStatus](#appservicecertificateorderstatus)***||
## AppServiceCertificateOrderSpec

Appears on:[AppServiceCertificateOrder](#appservicecertificateorder), [AppServiceCertificateOrderStatus](#appservicecertificateorderstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appServiceCertificateNotRenewableReasons` | ***[]string***| ***(Optional)*** |
| `autoRenew` | ***bool***| ***(Optional)*** |
| `certificates` | ***[[]AppServiceCertificateOrderSpecCertificates](#appservicecertificateorderspeccertificates)***| ***(Optional)*** |
| `csr` | ***string***| ***(Optional)*** |
| `distinguishedName` | ***string***| ***(Optional)*** |
| `domainVerificationToken` | ***string***| ***(Optional)*** |
| `expirationTime` | ***string***| ***(Optional)*** |
| `intermediateThumbprint` | ***string***| ***(Optional)*** |
| `isPrivateKeyExternal` | ***bool***| ***(Optional)*** |
| `keySize` | ***int64***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `productType` | ***string***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `rootThumbprint` | ***string***| ***(Optional)*** |
| `signedCertificateThumbprint` | ***string***| ***(Optional)*** |
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `validityInYears` | ***int64***| ***(Optional)*** |
## AppServiceCertificateOrderSpecCertificates

Appears on:[AppServiceCertificateOrderSpec](#appservicecertificateorderspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `certificateName` | ***string***| ***(Optional)*** |
| `keyVaultID` | ***string***| ***(Optional)*** |
| `keyVaultSecretName` | ***string***| ***(Optional)*** |
| `provisioningState` | ***string***| ***(Optional)*** |
## AppServiceCertificateOrderStatus

Appears on:[AppServiceCertificateOrder](#appservicecertificateorder)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppServiceCertificateOrderSpec](#appservicecertificateorderspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AppServiceCertificateOrderStatus](#appservicecertificateorderstatus)

---
