---
title: ElasticsearchDomainPolicy
menu:
  docs_{{ .version }}:
    identifier: elasticsearchdomainpolicy-aws.kubeform.com
    name: ElasticsearchDomainPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ElasticsearchDomainPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticsearchDomainPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticsearchDomainPolicySpec](#elasticsearchdomainpolicyspec)***||
| `status` | ***[ElasticsearchDomainPolicyStatus](#elasticsearchdomainpolicystatus)***||
## ElasticsearchDomainPolicySpec

Appears on:[ElasticsearchDomainPolicy](#elasticsearchdomainpolicy), [ElasticsearchDomainPolicyStatus](#elasticsearchdomainpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accessPolicies` | ***string***||
| `domainName` | ***string***||
## ElasticsearchDomainPolicyStatus

Appears on:[ElasticsearchDomainPolicy](#elasticsearchdomainpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticsearchDomainPolicySpec](#elasticsearchdomainpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ElasticsearchDomainPolicyStatus](#elasticsearchdomainpolicystatus)

---
