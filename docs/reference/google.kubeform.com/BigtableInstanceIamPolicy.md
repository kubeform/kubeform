---
title: BigtableInstanceIamPolicy
menu:
  docs_{{ .version }}:
    identifier: bigtableinstanceiampolicy-google.kubeform.com
    name: BigtableInstanceIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## BigtableInstanceIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `BigtableInstanceIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[BigtableInstanceIamPolicySpec](#bigtableinstanceiampolicyspec)***||
| `status` | ***[BigtableInstanceIamPolicyStatus](#bigtableinstanceiampolicystatus)***||
## BigtableInstanceIamPolicySpec

Appears on:[BigtableInstanceIamPolicy](#bigtableinstanceiampolicy), [BigtableInstanceIamPolicyStatus](#bigtableinstanceiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `instance` | ***string***||
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
## BigtableInstanceIamPolicyStatus

Appears on:[BigtableInstanceIamPolicy](#bigtableinstanceiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[BigtableInstanceIamPolicySpec](#bigtableinstanceiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[BigtableInstanceIamPolicyStatus](#bigtableinstanceiampolicystatus)

---
