---
title: DataprocJobIamPolicy
menu:
  docs_{{ .version }}:
    identifier: dataprocjobiampolicy-google.kubeform.com
    name: DataprocJobIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DataprocJobIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `DataprocJobIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DataprocJobIamPolicySpec](#dataprocjobiampolicyspec)***||
| `status` | ***[DataprocJobIamPolicyStatus](#dataprocjobiampolicystatus)***||
## DataprocJobIamPolicySpec

Appears on:[DataprocJobIamPolicy](#dataprocjobiampolicy), [DataprocJobIamPolicyStatus](#dataprocjobiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `jobID` | ***string***||
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `region` | ***string***| ***(Optional)*** |
## DataprocJobIamPolicyStatus

Appears on:[DataprocJobIamPolicy](#dataprocjobiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DataprocJobIamPolicySpec](#dataprocjobiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DataprocJobIamPolicyStatus](#dataprocjobiampolicystatus)

---
