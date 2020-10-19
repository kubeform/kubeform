---
title: IamAccountAlias
menu:
  docs_{{ .version }}:
    identifier: iamaccountalias-aws.kubeform.com
    name: IamAccountAlias
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamAccountAlias
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamAccountAlias` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamAccountAliasSpec](#iamaccountaliasspec)***||
| `status` | ***[IamAccountAliasStatus](#iamaccountaliasstatus)***||
## IamAccountAliasSpec

Appears on:[IamAccountAlias](#iamaccountalias), [IamAccountAliasStatus](#iamaccountaliasstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountAlias` | ***string***||
## IamAccountAliasStatus

Appears on:[IamAccountAlias](#iamaccountalias)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamAccountAliasSpec](#iamaccountaliasspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IamAccountAliasStatus](#iamaccountaliasstatus)

---
