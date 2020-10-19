---
title: DevspaceController
menu:
  docs_{{ .version }}:
    identifier: devspacecontroller-azurerm.kubeform.com
    name: DevspaceController
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DevspaceController
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DevspaceController` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevspaceControllerSpec](#devspacecontrollerspec)***||
| `status` | ***[DevspaceControllerStatus](#devspacecontrollerstatus)***||
## DevspaceControllerSpec

Appears on:[DevspaceController](#devspacecontroller), [DevspaceControllerStatus](#devspacecontrollerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `dataPlaneFqdn` | ***string***| ***(Optional)*** |
| `hostSuffix` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `sku` | ***[[]DevspaceControllerSpecSku](#devspacecontrollerspecsku)***| ***(Optional)*** Deprecated|
| `skuName` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetContainerHostResourceID` | ***string***||
## DevspaceControllerSpecSku

Appears on:[DevspaceControllerSpec](#devspacecontrollerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `tier` | ***string***||
## DevspaceControllerStatus

Appears on:[DevspaceController](#devspacecontroller)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevspaceControllerSpec](#devspacecontrollerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DevspaceControllerStatus](#devspacecontrollerstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `target_container_host_credentials_base64` | ***string*** ||
