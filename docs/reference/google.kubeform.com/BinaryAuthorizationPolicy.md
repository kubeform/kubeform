---
title: BinaryAuthorizationPolicy
menu:
  docs_{{ .version }}:
    identifier: binaryauthorizationpolicy-google.kubeform.com
    name: BinaryAuthorizationPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BinaryAuthorizationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BinaryAuthorizationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BinaryAuthorizationPolicySpec](#BinaryAuthorizationPolicySpec)***||
| `status` | ***[BinaryAuthorizationPolicyStatus](#BinaryAuthorizationPolicyStatus)***||
## BinaryAuthorizationPolicySpec
##### (Appears on:[BinaryAuthorizationPolicy](#BinaryAuthorizationPolicy), [BinaryAuthorizationPolicyStatus](#BinaryAuthorizationPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `admissionWhitelistPatterns` | ***[[]BinaryAuthorizationPolicySpecAdmissionWhitelistPatterns](#BinaryAuthorizationPolicySpecAdmissionWhitelistPatterns)***| ***(Optional)*** |
| `clusterAdmissionRules` | ***[[]BinaryAuthorizationPolicySpecClusterAdmissionRules](#BinaryAuthorizationPolicySpecClusterAdmissionRules)***| ***(Optional)*** |
| `defaultAdmissionRule` | ***[[]BinaryAuthorizationPolicySpecDefaultAdmissionRule](#BinaryAuthorizationPolicySpecDefaultAdmissionRule)***||
| `description` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
## BinaryAuthorizationPolicySpecAdmissionWhitelistPatterns
##### (Appears on:[BinaryAuthorizationPolicySpec](#BinaryAuthorizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `namePattern` | ***string***| ***(Optional)*** |
## BinaryAuthorizationPolicySpecClusterAdmissionRules
##### (Appears on:[BinaryAuthorizationPolicySpec](#BinaryAuthorizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `cluster` | ***string***||
| `enforcementMode` | ***string***| ***(Optional)*** |
| `evaluationMode` | ***string***| ***(Optional)*** |
| `requireAttestationsBy` | ***[]string***| ***(Optional)*** |
## BinaryAuthorizationPolicySpecDefaultAdmissionRule
##### (Appears on:[BinaryAuthorizationPolicySpec](#BinaryAuthorizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enforcementMode` | ***string***||
| `evaluationMode` | ***string***||
| `requireAttestationsBy` | ***[]string***| ***(Optional)*** |
## BinaryAuthorizationPolicyStatus
##### (Appears on:[BinaryAuthorizationPolicy](#BinaryAuthorizationPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BinaryAuthorizationPolicySpec](#BinaryAuthorizationPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
