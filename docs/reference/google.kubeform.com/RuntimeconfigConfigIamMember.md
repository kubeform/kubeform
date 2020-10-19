---
title: RuntimeconfigConfigIamMember
menu:
  docs_{{ .version }}:
    identifier: runtimeconfigconfigiammember-google.kubeform.com
    name: RuntimeconfigConfigIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## RuntimeconfigConfigIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `RuntimeconfigConfigIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RuntimeconfigConfigIamMemberSpec](#runtimeconfigconfigiammemberspec)***||
| `status` | ***[RuntimeconfigConfigIamMemberStatus](#runtimeconfigconfigiammemberstatus)***||
## Phase(`string` alias)

Appears on:[RuntimeconfigConfigIamMemberStatus](#runtimeconfigconfigiammemberstatus)

## RuntimeconfigConfigIamMemberSpec

Appears on:[RuntimeconfigConfigIamMember](#runtimeconfigconfigiammember), [RuntimeconfigConfigIamMemberStatus](#runtimeconfigconfigiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `config` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
## RuntimeconfigConfigIamMemberStatus

Appears on:[RuntimeconfigConfigIamMember](#runtimeconfigconfigiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RuntimeconfigConfigIamMemberSpec](#runtimeconfigconfigiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
