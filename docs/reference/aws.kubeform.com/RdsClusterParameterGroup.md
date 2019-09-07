---
title: RdsClusterParameterGroup
menu:
  docs_v0.0.1:
    identifier: rdsclusterparametergroup-aws.kubeform.com
    name: RdsClusterParameterGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## RdsClusterParameterGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RdsClusterParameterGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RdsClusterParameterGroupSpec](#RdsClusterParameterGroupSpec)***||
| `status` | ***[RdsClusterParameterGroupStatus](#RdsClusterParameterGroupStatus)***||
## RdsClusterParameterGroupSpec
##### (Appears on:[RdsClusterParameterGroup](#RdsClusterParameterGroup), [RdsClusterParameterGroupStatus](#RdsClusterParameterGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `family` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `parameter` | ***[[]RdsClusterParameterGroupSpecParameter](#RdsClusterParameterGroupSpecParameter)***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## RdsClusterParameterGroupSpecParameter
##### (Appears on:[RdsClusterParameterGroupSpec](#RdsClusterParameterGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `applyMethod` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `value` | ***string***||
## RdsClusterParameterGroupStatus
##### (Appears on:[RdsClusterParameterGroup](#RdsClusterParameterGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RdsClusterParameterGroupSpec](#RdsClusterParameterGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
