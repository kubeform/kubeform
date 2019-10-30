---
title: KmsAlias
menu:
  docs_{{ .version }}:
    identifier: kmsalias-aws.kubeform.com
    name: KmsAlias
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KmsAlias
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsAlias` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsAliasSpec](#kmsaliasspec)***||
| `status` | ***[KmsAliasStatus](#kmsaliasstatus)***||
## KmsAliasSpec

Appears on:[KmsAlias](#kmsalias), [KmsAliasStatus](#kmsaliasstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `targetKeyArn` | ***string***| ***(Optional)*** |
| `targetKeyID` | ***string***||
## KmsAliasStatus

Appears on:[KmsAlias](#kmsalias)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsAliasSpec](#kmsaliasspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[KmsAliasStatus](#kmsaliasstatus)

---
