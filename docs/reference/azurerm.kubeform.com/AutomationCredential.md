---
title: AutomationCredential
menu:
  docs_v0.0.1:
    identifier: automationcredential-azurerm.kubeform.com
    name: AutomationCredential
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## AutomationCredential
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AutomationCredential` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AutomationCredentialSpec](#AutomationCredentialSpec)***||
| `status` | ***[AutomationCredentialStatus](#AutomationCredentialStatus)***||
## AutomationCredentialSpec
##### (Appears on:[AutomationCredential](#AutomationCredential), [AutomationCredentialStatus](#AutomationCredentialStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `accountName` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `username` | ***string***||
## AutomationCredentialStatus
##### (Appears on:[AutomationCredential](#AutomationCredential))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AutomationCredentialSpec](#AutomationCredentialSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `password` | ***string*** ||
