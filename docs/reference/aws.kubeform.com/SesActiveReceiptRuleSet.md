---
title: SesActiveReceiptRuleSet
menu:
  docs_v0.0.1:
    identifier: sesactivereceiptruleset-aws.kubeform.com
    name: SesActiveReceiptRuleSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SesActiveReceiptRuleSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesActiveReceiptRuleSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesActiveReceiptRuleSetSpec](#SesActiveReceiptRuleSetSpec)***||
| `status` | ***[SesActiveReceiptRuleSetStatus](#SesActiveReceiptRuleSetStatus)***||
## SesActiveReceiptRuleSetSpec
##### (Appears on:[SesActiveReceiptRuleSet](#SesActiveReceiptRuleSet), [SesActiveReceiptRuleSetStatus](#SesActiveReceiptRuleSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `ruleSetName` | ***string***||
## SesActiveReceiptRuleSetStatus
##### (Appears on:[SesActiveReceiptRuleSet](#SesActiveReceiptRuleSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesActiveReceiptRuleSetSpec](#SesActiveReceiptRuleSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
