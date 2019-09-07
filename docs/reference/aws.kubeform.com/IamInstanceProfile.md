---
title: IamInstanceProfile
menu:
  docs_v0.0.1:
    identifier: iaminstanceprofile-aws.kubeform.com
    name: IamInstanceProfile
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IamInstanceProfile
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamInstanceProfile` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamInstanceProfileSpec](#IamInstanceProfileSpec)***||
| `status` | ***[IamInstanceProfileStatus](#IamInstanceProfileStatus)***||
## IamInstanceProfileSpec
##### (Appears on:[IamInstanceProfile](#IamInstanceProfile), [IamInstanceProfileStatus](#IamInstanceProfileStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `createDate` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `role` | ***string***| ***(Optional)*** |
| `roles` | ***[]string***| ***(Optional)*** Deprecated|
| `uniqueID` | ***string***| ***(Optional)*** |
## IamInstanceProfileStatus
##### (Appears on:[IamInstanceProfile](#IamInstanceProfile))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamInstanceProfileSpec](#IamInstanceProfileSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
