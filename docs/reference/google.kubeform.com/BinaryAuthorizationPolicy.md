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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BinaryAuthorizationPolicySpec](#binaryauthorizationpolicyspec)***||
| `status` | ***[BinaryAuthorizationPolicyStatus](#binaryauthorizationpolicystatus)***||
## BinaryAuthorizationPolicySpec

Appears on:[BinaryAuthorizationPolicy](#binaryauthorizationpolicy), [BinaryAuthorizationPolicyStatus](#binaryauthorizationpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `admissionWhitelistPatterns` | ***[[]BinaryAuthorizationPolicySpecAdmissionWhitelistPatterns](#binaryauthorizationpolicyspecadmissionwhitelistpatterns)***| ***(Optional)*** |
| `clusterAdmissionRules` | ***[[]BinaryAuthorizationPolicySpecClusterAdmissionRules](#binaryauthorizationpolicyspecclusteradmissionrules)***| ***(Optional)*** |
| `defaultAdmissionRule` | ***[[]BinaryAuthorizationPolicySpecDefaultAdmissionRule](#binaryauthorizationpolicyspecdefaultadmissionrule)***||
| `description` | ***string***| ***(Optional)*** |
| `globalPolicyEvaluationMode` | ***string***| ***(Optional)*** |
| `project` | ***string***| ***(Optional)*** |
## BinaryAuthorizationPolicySpecAdmissionWhitelistPatterns

Appears on:[BinaryAuthorizationPolicySpec](#binaryauthorizationpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `namePattern` | ***string***| ***(Optional)*** |
## BinaryAuthorizationPolicySpecClusterAdmissionRules

Appears on:[BinaryAuthorizationPolicySpec](#binaryauthorizationpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `cluster` | ***string***||
| `enforcementMode` | ***string***| ***(Optional)*** |
| `evaluationMode` | ***string***| ***(Optional)*** |
| `requireAttestationsBy` | ***[]string***| ***(Optional)*** |
## BinaryAuthorizationPolicySpecDefaultAdmissionRule

Appears on:[BinaryAuthorizationPolicySpec](#binaryauthorizationpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enforcementMode` | ***string***||
| `evaluationMode` | ***string***||
| `requireAttestationsBy` | ***[]string***| ***(Optional)*** |
## BinaryAuthorizationPolicyStatus

Appears on:[BinaryAuthorizationPolicy](#binaryauthorizationpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BinaryAuthorizationPolicySpec](#binaryauthorizationpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BinaryAuthorizationPolicyStatus](#binaryauthorizationpolicystatus)

---
