## ComputeForwardingRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeForwardingRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeForwardingRuleSpec](#ComputeForwardingRuleSpec)***||
| `status` | ***[ComputeForwardingRuleStatus](#ComputeForwardingRuleStatus)***||
## ComputeForwardingRuleSpec
##### (Appears on:[ComputeForwardingRule](#ComputeForwardingRule), [ComputeForwardingRuleStatus](#ComputeForwardingRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `backendService` | ***string***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `ipProtocol` | ***string***| ***(Optional)*** |
| `ipVersion` | ***string***| ***(Optional)*** |
| `labelFingerprint` | ***string***| ***(Optional)*** |
| `labels` | ***map[string]string***| ***(Optional)*** Deprecated|
| `loadBalancingScheme` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `network` | ***string***| ***(Optional)*** |
| `networkTier` | ***string***| ***(Optional)*** |
| `portRange` | ***string***| ***(Optional)*** |
| `ports` | ***[]string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `serviceLabel` | ***string***| ***(Optional)*** Deprecated|
| `serviceName` | ***string***| ***(Optional)*** Deprecated|
| `subnetwork` | ***string***| ***(Optional)*** |
| `target` | ***string***| ***(Optional)*** |
## ComputeForwardingRuleStatus
##### (Appears on:[ComputeForwardingRule](#ComputeForwardingRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeForwardingRuleSpec](#ComputeForwardingRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
