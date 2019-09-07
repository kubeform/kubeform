---
title: ElasticacheSecurityGroup
menu:
  docs_v0.0.1:
    identifier: elasticachesecuritygroup-aws.kubeform.com
    name: ElasticacheSecurityGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ElasticacheSecurityGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticacheSecurityGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticacheSecurityGroupSpec](#ElasticacheSecurityGroupSpec)***||
| `status` | ***[ElasticacheSecurityGroupStatus](#ElasticacheSecurityGroupStatus)***||
## ElasticacheSecurityGroupSpec
##### (Appears on:[ElasticacheSecurityGroup](#ElasticacheSecurityGroup), [ElasticacheSecurityGroupStatus](#ElasticacheSecurityGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `securityGroupNames` | ***[]string***||
## ElasticacheSecurityGroupStatus
##### (Appears on:[ElasticacheSecurityGroup](#ElasticacheSecurityGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticacheSecurityGroupSpec](#ElasticacheSecurityGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
