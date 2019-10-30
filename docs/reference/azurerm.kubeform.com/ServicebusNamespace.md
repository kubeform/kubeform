---
title: ServicebusNamespace
menu:
  docs_{{ .version }}:
    identifier: servicebusnamespace-azurerm.kubeform.com
    name: ServicebusNamespace
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServicebusNamespace
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ServicebusNamespace` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServicebusNamespaceSpec](#servicebusnamespacespec)***||
| `status` | ***[ServicebusNamespaceStatus](#servicebusnamespacestatus)***||
## Phase(`string` alias)

Appears on:[ServicebusNamespaceStatus](#servicebusnamespacestatus)

## ServicebusNamespaceSpec

Appears on:[ServicebusNamespace](#servicebusnamespace), [ServicebusNamespaceStatus](#servicebusnamespacestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `capacity` | ***int***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## ServicebusNamespaceStatus

Appears on:[ServicebusNamespace](#servicebusnamespace)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServicebusNamespaceSpec](#servicebusnamespacespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `default_primary_connection_string` | ***string*** ||
| `default_primary_key` | ***string*** ||
| `default_secondary_connection_string` | ***string*** ||
| `default_secondary_key` | ***string*** ||
