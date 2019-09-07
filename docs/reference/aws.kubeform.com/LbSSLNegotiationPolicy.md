---
title: LbSSLNegotiationPolicy
menu:
  docs_v0.0.1:
    identifier: lbsslnegotiationpolicy-aws.kubeform.com
    name: LbSSLNegotiationPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## LbSSLNegotiationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbSSLNegotiationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbSSLNegotiationPolicySpec](#LbSSLNegotiationPolicySpec)***||
| `status` | ***[LbSSLNegotiationPolicyStatus](#LbSSLNegotiationPolicyStatus)***||
## LbSSLNegotiationPolicySpec
##### (Appears on:[LbSSLNegotiationPolicy](#LbSSLNegotiationPolicy), [LbSSLNegotiationPolicyStatus](#LbSSLNegotiationPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attribute` | ***[[]LbSSLNegotiationPolicySpecAttribute](#LbSSLNegotiationPolicySpecAttribute)***| ***(Optional)*** |
| `lbPort` | ***int***||
| `loadBalancer` | ***string***||
| `name` | ***string***||
## LbSSLNegotiationPolicySpecAttribute
##### (Appears on:[LbSSLNegotiationPolicySpec](#LbSSLNegotiationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***||
| `value` | ***string***||
## LbSSLNegotiationPolicyStatus
##### (Appears on:[LbSSLNegotiationPolicy](#LbSSLNegotiationPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbSSLNegotiationPolicySpec](#LbSSLNegotiationPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
