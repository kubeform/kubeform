---
title: LoadBalancerPolicy
menu:
  docs_{{ .version }}:
    identifier: loadbalancerpolicy-aws.kubeform.com
    name: LoadBalancerPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LoadBalancerPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LoadBalancerPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LoadBalancerPolicySpec](#LoadBalancerPolicySpec)***||
| `status` | ***[LoadBalancerPolicyStatus](#LoadBalancerPolicyStatus)***||
## LoadBalancerPolicySpec
##### (Appears on:[LoadBalancerPolicy](#LoadBalancerPolicy), [LoadBalancerPolicyStatus](#LoadBalancerPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `loadBalancerName` | ***string***||
| `policyAttribute` | ***[[]LoadBalancerPolicySpecPolicyAttribute](#LoadBalancerPolicySpecPolicyAttribute)***| ***(Optional)*** |
| `policyName` | ***string***||
| `policyTypeName` | ***string***||
## LoadBalancerPolicySpecPolicyAttribute
##### (Appears on:[LoadBalancerPolicySpec](#LoadBalancerPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `name` | ***string***| ***(Optional)*** |
| `value` | ***string***| ***(Optional)*** |
## LoadBalancerPolicyStatus
##### (Appears on:[LoadBalancerPolicy](#LoadBalancerPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LoadBalancerPolicySpec](#LoadBalancerPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
