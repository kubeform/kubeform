---
title: SesIdentityPolicy
menu:
  docs_{{ .version }}:
    identifier: sesidentitypolicy-aws.kubeform.com
    name: SesIdentityPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SesIdentityPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SesIdentityPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SesIdentityPolicySpec](#sesidentitypolicyspec)***||
| `status` | ***[SesIdentityPolicyStatus](#sesidentitypolicystatus)***||
## SesIdentityPolicySpec


Appears on:[SesIdentityPolicy](#sesidentitypolicy), [SesIdentityPolicyStatus](#sesidentitypolicystatus)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `identity` | ***string***||
| `name` | ***string***||
| `policy` | ***string***||
## SesIdentityPolicyStatus


Appears on:[SesIdentityPolicy](#sesidentitypolicy)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SesIdentityPolicySpec](#sesidentitypolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
