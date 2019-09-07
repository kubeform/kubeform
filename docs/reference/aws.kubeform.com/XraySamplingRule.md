---
title: XraySamplingRule
menu:
  docs_v0.0.1:
    identifier: xraysamplingrule-aws.kubeform.com
    name: XraySamplingRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## XraySamplingRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `XraySamplingRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[XraySamplingRuleSpec](#XraySamplingRuleSpec)***||
| `status` | ***[XraySamplingRuleStatus](#XraySamplingRuleStatus)***||
## XraySamplingRuleSpec
##### (Appears on:[XraySamplingRule](#XraySamplingRule), [XraySamplingRuleStatus](#XraySamplingRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `attributes` | ***map[string]string***| ***(Optional)*** |
| `fixedRate` | ***encoding/json.Number***||
| `host` | ***string***||
| `httpMethod` | ***string***||
| `priority` | ***int***||
| `reservoirSize` | ***int***||
| `resourceArn` | ***string***||
| `ruleName` | ***string***| ***(Optional)*** |
| `serviceName` | ***string***||
| `serviceType` | ***string***||
| `urlPath` | ***string***||
| `version` | ***int***||
## XraySamplingRuleStatus
##### (Appears on:[XraySamplingRule](#XraySamplingRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[XraySamplingRuleSpec](#XraySamplingRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
