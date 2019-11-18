---
title: SsmActivation
menu:
  docs_{{ .version }}:
    identifier: ssmactivation-aws.kubeform.com
    name: SsmActivation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SsmActivation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmActivation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmActivationSpec](#ssmactivationspec)***||
| `status` | ***[SsmActivationStatus](#ssmactivationstatus)***||
## Phase(`string` alias)

Appears on:[SsmActivationStatus](#ssmactivationstatus)

## SsmActivationSpec

Appears on:[SsmActivation](#ssmactivation), [SsmActivationStatus](#ssmactivationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `activationCode` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `expirationDate` | ***string***| ***(Optional)*** |
| `expired` | ***string***| ***(Optional)*** |
| `iamRole` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `registrationCount` | ***int64***| ***(Optional)*** |
| `registrationLimit` | ***int64***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SsmActivationStatus

Appears on:[SsmActivation](#ssmactivation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmActivationSpec](#ssmactivationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
