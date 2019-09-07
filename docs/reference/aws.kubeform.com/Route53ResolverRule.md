---
title: Route53ResolverRule
menu:
  docs_v0.0.1:
    identifier: route53resolverrule-aws.kubeform.com
    name: Route53ResolverRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## Route53ResolverRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Route53ResolverRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53ResolverRuleSpec](#Route53ResolverRuleSpec)***||
| `status` | ***[Route53ResolverRuleStatus](#Route53ResolverRuleStatus)***||
## Route53ResolverRuleSpec
##### (Appears on:[Route53ResolverRule](#Route53ResolverRule), [Route53ResolverRuleStatus](#Route53ResolverRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `domainName` | ***string***||
| `name` | ***string***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `resolverEndpointID` | ***string***| ***(Optional)*** |
| `ruleType` | ***string***||
| `shareStatus` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `targetIP` | ***[[]Route53ResolverRuleSpecTargetIP](#Route53ResolverRuleSpecTargetIP)***| ***(Optional)*** |
## Route53ResolverRuleSpecTargetIP
##### (Appears on:[Route53ResolverRuleSpec](#Route53ResolverRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ip` | ***string***||
| `port` | ***int***| ***(Optional)*** |
## Route53ResolverRuleStatus
##### (Appears on:[Route53ResolverRule](#Route53ResolverRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53ResolverRuleSpec](#Route53ResolverRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
