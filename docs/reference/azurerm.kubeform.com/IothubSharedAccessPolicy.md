---
title: IothubSharedAccessPolicy
menu:
  docs_{{ .version }}:
    identifier: iothubsharedaccesspolicy-azurerm.kubeform.com
    name: IothubSharedAccessPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IothubSharedAccessPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `IothubSharedAccessPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IothubSharedAccessPolicySpec](#iothubsharedaccesspolicyspec)***||
| `status` | ***[IothubSharedAccessPolicyStatus](#iothubsharedaccesspolicystatus)***||
## IothubSharedAccessPolicySpec

Appears on:[IothubSharedAccessPolicy](#iothubsharedaccesspolicy), [IothubSharedAccessPolicyStatus](#iothubsharedaccesspolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `deviceConnect` | ***bool***| ***(Optional)*** |
| `iothubName` | ***string***||
| `name` | ***string***||
| `registryRead` | ***bool***| ***(Optional)*** |
| `registryWrite` | ***bool***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `serviceConnect` | ***bool***| ***(Optional)*** |
## IothubSharedAccessPolicyStatus

Appears on:[IothubSharedAccessPolicy](#iothubsharedaccesspolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IothubSharedAccessPolicySpec](#iothubsharedaccesspolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_connection_string` | ***string*** ||
| `primary_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
| `secondary_key` | ***string*** ||
