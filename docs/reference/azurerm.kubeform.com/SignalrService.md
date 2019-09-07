---
title: SignalrService
menu:
  docs_v0.0.1:
    identifier: signalrservice-azurerm.kubeform.com
    name: SignalrService
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SignalrService
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `SignalrService` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SignalrServiceSpec](#SignalrServiceSpec)***||
| `status` | ***[SignalrServiceStatus](#SignalrServiceStatus)***||
## SignalrServiceSpec
##### (Appears on:[SignalrService](#SignalrService), [SignalrServiceStatus](#SignalrServiceStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `hostname` | ***string***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `publicPort` | ***int***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `serverPort` | ***int***| ***(Optional)*** |
| `sku` | ***[[]SignalrServiceSpecSku](#SignalrServiceSpecSku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## SignalrServiceSpecSku
##### (Appears on:[SignalrServiceSpec](#SignalrServiceSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `capacity` | ***int***||
| `name` | ***string***||
## SignalrServiceStatus
##### (Appears on:[SignalrService](#SignalrService))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SignalrServiceSpec](#SignalrServiceSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `primary_access_key` | ***string*** ||
| `primary_connection_string` | ***string*** ||
| `secondary_access_key` | ***string*** ||
| `secondary_connection_string` | ***string*** ||
