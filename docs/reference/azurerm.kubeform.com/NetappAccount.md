---
title: NetappAccount
menu:
  docs_{{ .version }}:
    identifier: netappaccount-azurerm.kubeform.com
    name: NetappAccount
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetappAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `NetappAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetappAccountSpec](#netappaccountspec)***||
| `status` | ***[NetappAccountStatus](#netappaccountstatus)***||
## NetappAccountSpec

Appears on:[NetappAccount](#netappaccount), [NetappAccountStatus](#netappaccountstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `activeDirectory` | ***[[]NetappAccountSpecActiveDirectory](#netappaccountspecactivedirectory)***| ***(Optional)*** |
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
## NetappAccountSpecActiveDirectory

Appears on:[NetappAccountSpec](#netappaccountspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `dnsServers` | ***[]string***||
| `domain` | ***string***||
| `organizationalUnit` | ***string***| ***(Optional)*** |
| `smbServerName` | ***string***||
| `username` | ***string***||
## NetappAccountStatus

Appears on:[NetappAccount](#netappaccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetappAccountSpec](#netappaccountspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NetappAccountStatus](#netappaccountstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `active_directory.<index>.password` | ***string*** ||
