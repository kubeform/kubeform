---
title: PublicIPPrefix
menu:
  docs_{{ .version }}:
    identifier: publicipprefix-azurerm.kubeform.com
    name: PublicIPPrefix
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## PublicIPPrefix
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `PublicIPPrefix` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[PublicIPPrefixSpec](#publicipprefixspec)***||
| `status` | ***[PublicIPPrefixStatus](#publicipprefixstatus)***||
## Phase(`string` alias)

Appears on:[PublicIPPrefixStatus](#publicipprefixstatus)

## PublicIPPrefixSpec

Appears on:[PublicIPPrefix](#publicipprefix), [PublicIPPrefixStatus](#publicipprefixstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ipPrefix` | ***string***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `prefixLength` | ***int64***| ***(Optional)*** |
| `resourceGroupName` | ***string***||
| `sku` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `zones` | ***[]string***| ***(Optional)*** |
## PublicIPPrefixStatus

Appears on:[PublicIPPrefix](#publicipprefix)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[PublicIPPrefixSpec](#publicipprefixspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
