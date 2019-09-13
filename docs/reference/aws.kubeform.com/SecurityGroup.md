---
title: SecurityGroup
menu:
  docs_{{ .version }}:
    identifier: securitygroup-aws.kubeform.com
    name: SecurityGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SecurityGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SecurityGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityGroupSpec](#SecurityGroupSpec)***||
| `status` | ***[SecurityGroupStatus](#SecurityGroupStatus)***||
## SecurityGroupSpec
##### (Appears on:[SecurityGroup](#SecurityGroup), [SecurityGroupStatus](#SecurityGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `egress` | ***[[]SecurityGroupSpecEgress](#SecurityGroupSpecEgress)***| ***(Optional)*** |
| `ingress` | ***[[]SecurityGroupSpecIngress](#SecurityGroupSpecIngress)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `namePrefix` | ***string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `revokeRulesOnDelete` | ***bool***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## SecurityGroupSpecEgress
##### (Appears on:[SecurityGroupSpec](#SecurityGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidrBlocks` | ***[]string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `fromPort` | ***int***||
| `ipv6CIDRBlocks` | ***[]string***| ***(Optional)*** |
| `prefixListIDS` | ***[]string***| ***(Optional)*** |
| `protocol` | ***string***||
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `self` | ***bool***| ***(Optional)*** |
| `toPort` | ***int***||
## SecurityGroupSpecIngress
##### (Appears on:[SecurityGroupSpec](#SecurityGroupSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cidrBlocks` | ***[]string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `fromPort` | ***int***||
| `ipv6CIDRBlocks` | ***[]string***| ***(Optional)*** |
| `prefixListIDS` | ***[]string***| ***(Optional)*** |
| `protocol` | ***string***||
| `securityGroups` | ***[]string***| ***(Optional)*** |
| `self` | ***bool***| ***(Optional)*** |
| `toPort` | ***int***||
## SecurityGroupStatus
##### (Appears on:[SecurityGroup](#SecurityGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityGroupSpec](#SecurityGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
