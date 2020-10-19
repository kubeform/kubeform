---
title: SecurityGroupRule
menu:
  docs_{{ .version }}:
    identifier: securitygrouprule-aws.kubeform.com
    name: SecurityGroupRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SecurityGroupRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SecurityGroupRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SecurityGroupRuleSpec](#securitygrouprulespec)***||
| `status` | ***[SecurityGroupRuleStatus](#securitygrouprulestatus)***||
## Phase(`string` alias)

Appears on:[SecurityGroupRuleStatus](#securitygrouprulestatus)

## SecurityGroupRuleSpec

Appears on:[SecurityGroupRule](#securitygrouprule), [SecurityGroupRuleStatus](#securitygrouprulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cidrBlocks` | ***[]string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `fromPort` | ***int64***||
| `ipv6CIDRBlocks` | ***[]string***| ***(Optional)*** |
| `prefixListIDS` | ***[]string***| ***(Optional)*** |
| `protocol` | ***string***||
| `securityGroupID` | ***string***||
| `self` | ***bool***| ***(Optional)*** |
| `sourceSecurityGroupID` | ***string***| ***(Optional)*** |
| `toPort` | ***int64***||
| `type` | ***string***|Type of rule, ingress (inbound) or egress (outbound).|
## SecurityGroupRuleStatus

Appears on:[SecurityGroupRule](#securitygrouprule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SecurityGroupRuleSpec](#securitygrouprulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
