---
title: BinaryAuthorizationAttestorIamPolicy
menu:
  docs_{{ .version }}:
    identifier: binaryauthorizationattestoriampolicy-google.kubeform.com
    name: BinaryAuthorizationAttestorIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BinaryAuthorizationAttestorIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BinaryAuthorizationAttestorIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BinaryAuthorizationAttestorIamPolicySpec](#binaryauthorizationattestoriampolicyspec)***||
| `status` | ***[BinaryAuthorizationAttestorIamPolicyStatus](#binaryauthorizationattestoriampolicystatus)***||
## BinaryAuthorizationAttestorIamPolicySpec

Appears on:[BinaryAuthorizationAttestorIamPolicy](#binaryauthorizationattestoriampolicy), [BinaryAuthorizationAttestorIamPolicyStatus](#binaryauthorizationattestoriampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attestor` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## BinaryAuthorizationAttestorIamPolicyStatus

Appears on:[BinaryAuthorizationAttestorIamPolicy](#binaryauthorizationattestoriampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BinaryAuthorizationAttestorIamPolicySpec](#binaryauthorizationattestoriampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BinaryAuthorizationAttestorIamPolicyStatus](#binaryauthorizationattestoriampolicystatus)

---
