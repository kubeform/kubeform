---
title: SesEmailIdentity
menu:
  docs_{{ .version }}:
    identifier: sesemailidentity-aws.kubeform.com
    name: SesEmailIdentity
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SesEmailIdentity
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesEmailIdentity` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesEmailIdentitySpec](#sesemailidentityspec)***||
| `status` | ***[SesEmailIdentityStatus](#sesemailidentitystatus)***||
## Phase(`string` alias)

Appears on:[SesEmailIdentityStatus](#sesemailidentitystatus)

## SesEmailIdentitySpec

Appears on:[SesEmailIdentity](#sesemailidentity), [SesEmailIdentityStatus](#sesemailidentitystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `email` | ***string***||
## SesEmailIdentityStatus

Appears on:[SesEmailIdentity](#sesemailidentity)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesEmailIdentitySpec](#sesemailidentityspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
