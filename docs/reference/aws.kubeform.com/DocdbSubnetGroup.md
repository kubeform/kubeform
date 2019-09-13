---
title: DocdbSubnetGroup
menu:
  docs_{{ .version }}:
    identifier: docdbsubnetgroup-aws.kubeform.com
    name: DocdbSubnetGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DocdbSubnetGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DocdbSubnetGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DocdbSubnetGroupSpec](#DocdbSubnetGroupSpec)***||
| `status` | ***[DocdbSubnetGroupStatus](#DocdbSubnetGroupStatus)***||
## DocdbSubnetGroupSpec
##### (Appears on:[DocdbSubnetGroup](#DocdbSubnetGroup), [DocdbSubnetGroupStatus](#DocdbSubnetGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
## DocdbSubnetGroupStatus
##### (Appears on:[DocdbSubnetGroup](#DocdbSubnetGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DocdbSubnetGroupSpec](#DocdbSubnetGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
