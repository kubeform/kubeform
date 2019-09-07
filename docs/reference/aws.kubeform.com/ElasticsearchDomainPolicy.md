---
title: ElasticsearchDomainPolicy
menu:
  docs_v0.0.1:
    identifier: elasticsearchdomainpolicy-aws.kubeform.com
    name: ElasticsearchDomainPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ElasticsearchDomainPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticsearchDomainPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticsearchDomainPolicySpec](#ElasticsearchDomainPolicySpec)***||
| `status` | ***[ElasticsearchDomainPolicyStatus](#ElasticsearchDomainPolicyStatus)***||
## ElasticsearchDomainPolicySpec
##### (Appears on:[ElasticsearchDomainPolicy](#ElasticsearchDomainPolicy), [ElasticsearchDomainPolicyStatus](#ElasticsearchDomainPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessPolicies` | ***string***||
| `domainName` | ***string***||
## ElasticsearchDomainPolicyStatus
##### (Appears on:[ElasticsearchDomainPolicy](#ElasticsearchDomainPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticsearchDomainPolicySpec](#ElasticsearchDomainPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
