---
title: SwfDomain
menu:
  docs_{{ .version }}:
    identifier: swfdomain-aws.kubeform.com
    name: SwfDomain
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SwfDomain
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SwfDomain` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SwfDomainSpec](#swfdomainspec)***||
| `status` | ***[SwfDomainStatus](#swfdomainstatus)***||
## Phase(`string` alias)

Appears on:[SwfDomainStatus](#swfdomainstatus)

## SwfDomainSpec

Appears on:[SwfDomain](#swfdomain), [SwfDomainStatus](#swfdomainstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `workflowExecutionRetentionPeriodInDays` | ***string***||
## SwfDomainStatus

Appears on:[SwfDomain](#swfdomain)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SwfDomainSpec](#swfdomainspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
