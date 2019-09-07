---
title: RedshiftSecurityGroup
menu:
  docs_v0.0.1:
    identifier: redshiftsecuritygroup-aws.kubeform.com
    name: RedshiftSecurityGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## RedshiftSecurityGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `RedshiftSecurityGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RedshiftSecurityGroupSpec](#RedshiftSecurityGroupSpec)***||
| `status` | ***[RedshiftSecurityGroupStatus](#RedshiftSecurityGroupStatus)***||
## RedshiftSecurityGroupSpec
##### (Appears on:[RedshiftSecurityGroup](#RedshiftSecurityGroup), [RedshiftSecurityGroupStatus](#RedshiftSecurityGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `ingress` | ***[[]RedshiftSecurityGroupSpecIngress](#RedshiftSecurityGroupSpecIngress)***||
| `name` | ***string***||
## RedshiftSecurityGroupSpecIngress
##### (Appears on:[RedshiftSecurityGroupSpec](#RedshiftSecurityGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidr` | ***string***| ***(Optional)*** |
| `securityGroupName` | ***string***| ***(Optional)*** |
| `securityGroupOwnerID` | ***string***| ***(Optional)*** |
## RedshiftSecurityGroupStatus
##### (Appears on:[RedshiftSecurityGroup](#RedshiftSecurityGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RedshiftSecurityGroupSpec](#RedshiftSecurityGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
