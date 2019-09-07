---
title: IamGroup
menu:
  docs_v0.0.1:
    identifier: iamgroup-aws.kubeform.com
    name: IamGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IamGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamGroupSpec](#IamGroupSpec)***||
| `status` | ***[IamGroupStatus](#IamGroupStatus)***||
## IamGroupSpec
##### (Appears on:[IamGroup](#IamGroup), [IamGroupStatus](#IamGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `path` | ***string***| ***(Optional)*** |
| `uniqueID` | ***string***| ***(Optional)*** |
## IamGroupStatus
##### (Appears on:[IamGroup](#IamGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamGroupSpec](#IamGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
