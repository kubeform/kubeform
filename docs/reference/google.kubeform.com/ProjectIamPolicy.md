---
title: ProjectIamPolicy
menu:
  docs_{{ .version }}:
    identifier: projectiampolicy-google.kubeform.com
    name: ProjectIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ProjectIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ProjectIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectIamPolicySpec](#projectiampolicyspec)***||
| `status` | ***[ProjectIamPolicyStatus](#projectiampolicystatus)***||
## ProjectIamPolicySpec

Appears on:[ProjectIamPolicy](#projectiampolicy), [ProjectIamPolicyStatus](#projectiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `authoritative` | ***bool***| ***(Optional)*** Deprecated|
| `disableProject` | ***bool***| ***(Optional)*** Deprecated|
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `restorePolicy` | ***string***| ***(Optional)*** Deprecated|
## ProjectIamPolicyStatus

Appears on:[ProjectIamPolicy](#projectiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectIamPolicySpec](#projectiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
