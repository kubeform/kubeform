## NetworkACLRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkACLRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkACLRuleSpec](#NetworkACLRuleSpec)***||
| `status` | ***[NetworkACLRuleStatus](#NetworkACLRuleStatus)***||
## NetworkACLRuleSpec
##### (Appears on:[NetworkACLRule](#NetworkACLRule), [NetworkACLRuleStatus](#NetworkACLRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cidrBlock` | ***string***| ***(Optional)*** |
| `egress` | ***bool***| ***(Optional)*** |
| `fromPort` | ***int***| ***(Optional)*** |
| `icmpCode` | ***string***| ***(Optional)*** |
| `icmpType` | ***string***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `networkACLID` | ***string***||
| `protocol` | ***string***||
| `ruleAction` | ***string***||
| `ruleNumber` | ***int***||
| `toPort` | ***int***| ***(Optional)*** |
## NetworkACLRuleStatus
##### (Appears on:[NetworkACLRule](#NetworkACLRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkACLRuleSpec](#NetworkACLRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
