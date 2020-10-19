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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ElasticacheSubnetGroupSpec](#elasticachesubnetgroupspec)***||
| `status` | ***[ElasticacheSubnetGroupStatus](#elasticachesubnetgroupstatus)***||
## ElasticacheSubnetGroupSpec

Appears on:[ElasticacheSubnetGroup](#elasticachesubnetgroup), [ElasticacheSubnetGroupStatus](#elasticachesubnetgroupstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `subnetIDS` | ***[]string***||
## ElasticacheSubnetGroupStatus

Appears on:[ElasticacheSubnetGroup](#elasticachesubnetgroup)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ElasticacheSubnetGroupSpec](#elasticachesubnetgroupspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ElasticacheSubnetGroupStatus](#elasticachesubnetgroupstatus)

---
