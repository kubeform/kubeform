---
title: CloudfunctionsFunctionIamPolicy
menu:
  docs_{{ .version }}:
    identifier: cloudfunctionsfunctioniampolicy-google.kubeform.com
    name: CloudfunctionsFunctionIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CloudfunctionsFunctionIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `CloudfunctionsFunctionIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CloudfunctionsFunctionIamPolicySpec](#cloudfunctionsfunctioniampolicyspec)***||
| `status` | ***[CloudfunctionsFunctionIamPolicyStatus](#cloudfunctionsfunctioniampolicystatus)***||
## CloudfunctionsFunctionIamPolicySpec

Appears on:[CloudfunctionsFunctionIamPolicy](#cloudfunctionsfunctioniampolicy), [CloudfunctionsFunctionIamPolicyStatus](#cloudfunctionsfunctioniampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cloudFunction` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
## CloudfunctionsFunctionIamPolicyStatus

Appears on:[CloudfunctionsFunctionIamPolicy](#cloudfunctionsfunctioniampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CloudfunctionsFunctionIamPolicySpec](#cloudfunctionsfunctioniampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[CloudfunctionsFunctionIamPolicyStatus](#cloudfunctionsfunctioniampolicystatus)

---
