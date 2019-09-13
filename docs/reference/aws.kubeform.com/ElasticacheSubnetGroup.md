---
title: ElasticacheSubnetGroup
menu:
  docs_{{ .version }}:
    identifier: elasticachesubnetgroup-aws.kubeform.com
    name: ElasticacheSubnetGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ElasticacheSubnetGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `ElasticacheSubnetGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticacheSubnetGroupSpec](#ElasticacheSubnetGroupSpec)***||
| `status` | ***[ElasticacheSubnetGroupStatus](#ElasticacheSubnetGroupStatus)***||
## ElasticacheSubnetGroupSpec
##### (Appears on:[ElasticacheSubnetGroup](#ElasticacheSubnetGroup), [ElasticacheSubnetGroupStatus](#ElasticacheSubnetGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `subnetIDS` | ***[]string***||
## ElasticacheSubnetGroupStatus
##### (Appears on:[ElasticacheSubnetGroup](#ElasticacheSubnetGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticacheSubnetGroupSpec](#ElasticacheSubnetGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
