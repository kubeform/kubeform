---
title: DaxSubnetGroup
menu:
  docs_{{ .version }}:
    identifier: daxsubnetgroup-aws.kubeform.com
    name: DaxSubnetGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DaxSubnetGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DaxSubnetGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DaxSubnetGroupSpec](#DaxSubnetGroupSpec)***||
| `status` | ***[DaxSubnetGroupStatus](#DaxSubnetGroupStatus)***||
## DaxSubnetGroupSpec
##### (Appears on:[DaxSubnetGroup](#DaxSubnetGroup), [DaxSubnetGroupStatus](#DaxSubnetGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `subnetIDS` | ***[]string***||
| `vpcID` | ***string***| ***(Optional)*** |
## DaxSubnetGroupStatus
##### (Appears on:[DaxSubnetGroup](#DaxSubnetGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DaxSubnetGroupSpec](#DaxSubnetGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
