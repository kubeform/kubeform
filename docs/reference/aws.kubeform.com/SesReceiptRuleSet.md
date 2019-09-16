---
title: SesReceiptRuleSet
menu:
  docs_{{ .version }}:
    identifier: sesreceiptruleset-aws.kubeform.com
    name: SesReceiptRuleSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SesReceiptRuleSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesReceiptRuleSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesReceiptRuleSetSpec](#SesReceiptRuleSetSpec)***||
| `status` | ***[SesReceiptRuleSetStatus](#SesReceiptRuleSetStatus)***||
## SesReceiptRuleSetSpec

Appears on:[SesReceiptRuleSet](#SesReceiptRuleSet), [SesReceiptRuleSetStatus](#SesReceiptRuleSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ruleSetName` | ***string***||
## SesReceiptRuleSetStatus

Appears on:[SesReceiptRuleSet](#SesReceiptRuleSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesReceiptRuleSetSpec](#SesReceiptRuleSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
