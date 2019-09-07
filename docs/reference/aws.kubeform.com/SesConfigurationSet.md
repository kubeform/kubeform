---
title: SesConfigurationSet
menu:
  docs_v0.0.1:
    identifier: sesconfigurationset-aws.kubeform.com
    name: SesConfigurationSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## SesConfigurationSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesConfigurationSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesConfigurationSetSpec](#SesConfigurationSetSpec)***||
| `status` | ***[SesConfigurationSetStatus](#SesConfigurationSetStatus)***||
## SesConfigurationSetSpec
##### (Appears on:[SesConfigurationSet](#SesConfigurationSet), [SesConfigurationSetStatus](#SesConfigurationSetStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `name` | ***string***||
## SesConfigurationSetStatus
##### (Appears on:[SesConfigurationSet](#SesConfigurationSet))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesConfigurationSetSpec](#SesConfigurationSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
