---
title: KmsKeyRingIamMember
menu:
  docs_{{ .version }}:
    identifier: kmskeyringiammember-google.kubeform.com
    name: KmsKeyRingIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## KmsKeyRingIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `KmsKeyRingIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[KmsKeyRingIamMemberSpec](#kmskeyringiammemberspec)***||
| `status` | ***[KmsKeyRingIamMemberStatus](#kmskeyringiammemberstatus)***||
## KmsKeyRingIamMemberSpec

Appears on:[KmsKeyRingIamMember](#kmskeyringiammember), [KmsKeyRingIamMemberStatus](#kmskeyringiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `keyRingID` | ***string***||
| `member` | ***string***||
| `role` | ***string***||
## KmsKeyRingIamMemberStatus

Appears on:[KmsKeyRingIamMember](#kmskeyringiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[KmsKeyRingIamMemberSpec](#kmskeyringiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
