---
title: DbSubnetGroup
menu:
  docs_v0.0.1:
    identifier: dbsubnetgroup-aws.kubeform.com
    name: DbSubnetGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## DbSubnetGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DbSubnetGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DbSubnetGroupSpec](#DbSubnetGroupSpec)***||
| `status` | ***[DbSubnetGroupStatus](#DbSubnetGroupStatus)***||
## DbSubnetGroupSpec
##### (Appears on:[DbSubnetGroup](#DbSubnetGroup), [DbSubnetGroupStatus](#DbSubnetGroupStatus))
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
## DbSubnetGroupStatus
##### (Appears on:[DbSubnetGroup](#DbSubnetGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DbSubnetGroupSpec](#DbSubnetGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
