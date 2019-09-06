## ComputeSecurityPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeSecurityPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeSecurityPolicySpec](#ComputeSecurityPolicySpec)***||
| `status` | ***[ComputeSecurityPolicyStatus](#ComputeSecurityPolicyStatus)***||
## ComputeSecurityPolicySpec
##### (Appears on:[ComputeSecurityPolicy](#ComputeSecurityPolicy), [ComputeSecurityPolicyStatus](#ComputeSecurityPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `fingerprint` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `rule` | ***[[]ComputeSecurityPolicySpecRule](#ComputeSecurityPolicySpecRule)***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
## ComputeSecurityPolicySpecRule
##### (Appears on:[ComputeSecurityPolicySpec](#ComputeSecurityPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `match` | ***[[]ComputeSecurityPolicySpecRuleMatch](#ComputeSecurityPolicySpecRuleMatch)***||
| `preview` | ***bool***| ***(Optional)*** |
| `priority` | ***int***||
## ComputeSecurityPolicySpecRuleMatch
##### (Appears on:[ComputeSecurityPolicySpecRule](#ComputeSecurityPolicySpecRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `config` | ***[[]ComputeSecurityPolicySpecRuleMatchConfig](#ComputeSecurityPolicySpecRuleMatchConfig)***||
| `versionedExpr` | ***string***||
## ComputeSecurityPolicySpecRuleMatchConfig
##### (Appears on:[ComputeSecurityPolicySpecRuleMatch](#ComputeSecurityPolicySpecRuleMatch))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `srcIPRanges` | ***[]string***||
## ComputeSecurityPolicyStatus
##### (Appears on:[ComputeSecurityPolicy](#ComputeSecurityPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeSecurityPolicySpec](#ComputeSecurityPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
