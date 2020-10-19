---
title: ContainerRegistryWebhook
menu:
  docs_{{ .version }}:
    identifier: containerregistrywebhook-azurerm.kubeform.com
    name: ContainerRegistryWebhook
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ContainerRegistryWebhook
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `ContainerRegistryWebhook` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ContainerRegistryWebhookSpec](#containerregistrywebhookspec)***||
| `status` | ***[ContainerRegistryWebhookStatus](#containerregistrywebhookstatus)***||
## ContainerRegistryWebhookSpec

Appears on:[ContainerRegistryWebhook](#containerregistrywebhook), [ContainerRegistryWebhookStatus](#containerregistrywebhookstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `actions` | ***[]string***||
| `customHeaders` | ***map[string]string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `registryName` | ***string***||
| `resourceGroupName` | ***string***||
| `scope` | ***string***| ***(Optional)*** |
| `serviceURI` | ***string***||
| `status` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## ContainerRegistryWebhookStatus

Appears on:[ContainerRegistryWebhook](#containerregistrywebhook)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ContainerRegistryWebhookSpec](#containerregistrywebhookspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ContainerRegistryWebhookStatus](#containerregistrywebhookstatus)

---
