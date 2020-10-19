---
title: FmsAdminAccount
menu:
  docs_{{ .version }}:
    identifier: fmsadminaccount-aws.kubeform.com
    name: FmsAdminAccount
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## FmsAdminAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `FmsAdminAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FmsAdminAccountSpec](#fmsadminaccountspec)***||
| `status` | ***[FmsAdminAccountStatus](#fmsadminaccountstatus)***||
## FmsAdminAccountSpec

Appears on:[FmsAdminAccount](#fmsadminaccount), [FmsAdminAccountStatus](#fmsadminaccountstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***| ***(Optional)*** |
## FmsAdminAccountStatus

Appears on:[FmsAdminAccount](#fmsadminaccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FmsAdminAccountSpec](#fmsadminaccountspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[FmsAdminAccountStatus](#fmsadminaccountstatus)

---
