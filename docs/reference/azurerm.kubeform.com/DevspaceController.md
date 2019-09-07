---
title: DevspaceController
menu:
  docs_v0.0.1:
    identifier: devspacecontroller-azurerm.kubeform.com
    name: DevspaceController
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DevspaceController
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DevspaceController` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevspaceControllerSpec](#DevspaceControllerSpec)***||
| `status` | ***[DevspaceControllerStatus](#DevspaceControllerStatus)***||
## DevspaceControllerSpec
##### (Appears on:[DevspaceController](#DevspaceController), [DevspaceControllerStatus](#DevspaceControllerStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `dataPlaneFqdn` | ***string***| ***(Optional)*** |
| `hostSuffix` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***[[]DevspaceControllerSpecSku](#DevspaceControllerSpecSku)***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetContainerHostResourceID` | ***string***||
## DevspaceControllerSpecSku
##### (Appears on:[DevspaceControllerSpec](#DevspaceControllerSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `tier` | ***string***||
## DevspaceControllerStatus
##### (Appears on:[DevspaceController](#DevspaceController))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevspaceControllerSpec](#DevspaceControllerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `target_container_host_credentials_base64` | ***string*** ||
