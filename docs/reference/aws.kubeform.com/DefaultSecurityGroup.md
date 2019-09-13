---
title: DefaultSecurityGroup
menu:
  docs_{{ .version }}:
    identifier: defaultsecuritygroup-aws.kubeform.com
    name: DefaultSecurityGroup
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DefaultSecurityGroup
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DefaultSecurityGroup` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DefaultSecurityGroupSpec](#DefaultSecurityGroupSpec)***||
| `status` | ***[DefaultSecurityGroupStatus](#DefaultSecurityGroupStatus)***||
## DefaultSecurityGroupSpec
##### (Appears on:[DefaultSecurityGroup](#DefaultSecurityGroup), [DefaultSecurityGroupStatus](#DefaultSecurityGroupStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `egress` | ***[[]DefaultSecurityGroupSpecEgress](#DefaultSecurityGroupSpecEgress)***| ***(Optional)*** |
| `ingress` | ***[[]DefaultSecurityGroupSpecIngress](#DefaultSecurityGroupSpecIngress)***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `revokeRulesOnDelete` | ***bool***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## DefaultSecurityGroupSpecEgress
##### (Appears on:[DefaultSecurityGroupSpec](#DefaultSecurityGroupSpec))
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
## DefaultSecurityGroupSpecIngress
##### (Appears on:[DefaultSecurityGroupSpec](#DefaultSecurityGroupSpec))
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
## DefaultSecurityGroupStatus
##### (Appears on:[DefaultSecurityGroup](#DefaultSecurityGroup))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DefaultSecurityGroupSpec](#DefaultSecurityGroupSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
