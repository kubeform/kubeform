---
title: ApiManagementIdentityProviderTwitter
menu:
  docs_{{ .version }}:
    identifier: apimanagementidentityprovidertwitter-azurerm.kubeform.com
    name: ApiManagementIdentityProviderTwitter
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ApiManagementIdentityProviderTwitter
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ApiManagementIdentityProviderTwitter` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ApiManagementIdentityProviderTwitterSpec](#apimanagementidentityprovidertwitterspec)***||
| `status` | ***[ApiManagementIdentityProviderTwitterStatus](#apimanagementidentityprovidertwitterstatus)***||
## ApiManagementIdentityProviderTwitterSpec

Appears on:[ApiManagementIdentityProviderTwitter](#apimanagementidentityprovidertwitter), [ApiManagementIdentityProviderTwitterStatus](#apimanagementidentityprovidertwitterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `apiKey` | ***string***||
| `apiManagementName` | ***string***||
| `resourceGroupName` | ***string***||
## ApiManagementIdentityProviderTwitterStatus

Appears on:[ApiManagementIdentityProviderTwitter](#apimanagementidentityprovidertwitter)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ApiManagementIdentityProviderTwitterSpec](#apimanagementidentityprovidertwitterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ApiManagementIdentityProviderTwitterStatus](#apimanagementidentityprovidertwitterstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `api_secret_key` | ***string*** ||
