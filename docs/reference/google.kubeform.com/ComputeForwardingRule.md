---
title: ComputeForwardingRule
menu:
  docs_{{ .version }}:
    identifier: computeforwardingrule-google.kubeform.com
    name: ComputeForwardingRule
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ComputeForwardingRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ComputeForwardingRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ComputeForwardingRuleSpec](#computeforwardingrulespec)***||
| `status` | ***[ComputeForwardingRuleStatus](#computeforwardingrulestatus)***||
## ComputeForwardingRuleSpec

Appears on:[ComputeForwardingRule](#computeforwardingrule), [ComputeForwardingRuleStatus](#computeforwardingrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allPorts` | ***bool***| ***(Optional)*** |
| `backendService` | ***string***| ***(Optional)*** |
| `creationTimestamp` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `ipAddress` | ***string***| ***(Optional)*** |
| `ipProtocol` | ***string***| ***(Optional)*** |
| `ipVersion` | ***string***| ***(Optional)*** Deprecated|
| `loadBalancingScheme` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `network` | ***string***| ***(Optional)*** |
| `networkTier` | ***string***| ***(Optional)*** |
| `portRange` | ***string***| ***(Optional)*** |
| `ports` | ***[]string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
| `selfLink` | ***string***| ***(Optional)*** |
| `serviceLabel` | ***string***| ***(Optional)*** |
| `serviceName` | ***string***| ***(Optional)*** |
| `subnetwork` | ***string***| ***(Optional)*** |
| `target` | ***string***| ***(Optional)*** |
## ComputeForwardingRuleStatus

Appears on:[ComputeForwardingRule](#computeforwardingrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ComputeForwardingRuleSpec](#computeforwardingrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[ComputeForwardingRuleStatus](#computeforwardingrulestatus)

---
