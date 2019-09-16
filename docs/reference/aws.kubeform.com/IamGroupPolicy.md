---
title: IamGroupPolicy
menu:
  docs_{{ .version }}:
    identifier: iamgrouppolicy-aws.kubeform.com
    name: IamGroupPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamGroupPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamGroupPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamGroupPolicySpec](#iamgrouppolicyspec)***||
| `status` | ***[IamGroupPolicyStatus](#iamgrouppolicystatus)***||
## IamGroupPolicySpec

Appears on:[IamGroupPolicy](#iamgrouppolicy), [IamGroupPolicyStatus](#iamgrouppolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `group` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `policy` | ***string***||
## IamGroupPolicyStatus

Appears on:[IamGroupPolicy](#iamgrouppolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamGroupPolicySpec](#iamgrouppolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---