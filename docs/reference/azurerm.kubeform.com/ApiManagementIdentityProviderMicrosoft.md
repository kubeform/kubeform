---
title: ApiManagementIdentityProviderMicrosoft
menu:
  docs_{{ .version }}:
    identifier: apimanagementidentityprovidermicrosoft-azurerm.kubeform.com
    name: ApiManagementIdentityProviderMicrosoft
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementIdentityProviderMicrosoft
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementIdentityProviderMicrosoft` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementIdentityProviderMicrosoftSpec](#apimanagementidentityprovidermicrosoftspec)***||
| `status` | ***[ApiManagementIdentityProviderMicrosoftStatus](#apimanagementidentityprovidermicrosoftstatus)***||
## ApiManagementIdentityProviderMicrosoftSpec

Appears on:[ApiManagementIdentityProviderMicrosoft](#apimanagementidentityprovidermicrosoft), [ApiManagementIdentityProviderMicrosoftStatus](#apimanagementidentityprovidermicrosoftstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `apiManagementName` | ***string***||
| `clientID` | ***string***||
| `resourceGroupName` | ***string***||
## ApiManagementIdentityProviderMicrosoftStatus

Appears on:[ApiManagementIdentityProviderMicrosoft](#apimanagementidentityprovidermicrosoft)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementIdentityProviderMicrosoftSpec](#apimanagementidentityprovidermicrosoftspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiManagementIdentityProviderMicrosoftStatus](#apimanagementidentityprovidermicrosoftstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `client_secret` | ***string*** ||
