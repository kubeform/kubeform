---
title: Route53DelegationSet
menu:
  docs_{{ .version }}:
    identifier: route53delegationset-aws.kubeform.com
    name: Route53DelegationSet
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Route53DelegationSet
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Route53DelegationSet` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53DelegationSetSpec](#Route53DelegationSetSpec)***||
| `status` | ***[Route53DelegationSetStatus](#Route53DelegationSetStatus)***||
## Route53DelegationSetSpec

Appears on:[Route53DelegationSet](#Route53DelegationSet), [Route53DelegationSetStatus](#Route53DelegationSetStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `nameServers` | ***[]string***| ***(Optional)*** |
| `referenceName` | ***string***| ***(Optional)*** |
## Route53DelegationSetStatus

Appears on:[Route53DelegationSet](#Route53DelegationSet)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53DelegationSetSpec](#Route53DelegationSetSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
