---
title: AppServiceCustomHostnameBinding
menu:
  docs_{{ .version }}:
    identifier: appservicecustomhostnamebinding-azurerm.kubeform.com
    name: AppServiceCustomHostnameBinding
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AppServiceCustomHostnameBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `AppServiceCustomHostnameBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AppServiceCustomHostnameBindingSpec](#appservicecustomhostnamebindingspec)***||
| `status` | ***[AppServiceCustomHostnameBindingStatus](#appservicecustomhostnamebindingstatus)***||
## AppServiceCustomHostnameBindingSpec

Appears on:[AppServiceCustomHostnameBinding](#appservicecustomhostnamebinding), [AppServiceCustomHostnameBindingStatus](#appservicecustomhostnamebindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appServiceName` | ***string***||
| `hostname` | ***string***||
| `resourceGroupName` | ***string***||
## AppServiceCustomHostnameBindingStatus

Appears on:[AppServiceCustomHostnameBinding](#appservicecustomhostnamebinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AppServiceCustomHostnameBindingSpec](#appservicecustomhostnamebindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---