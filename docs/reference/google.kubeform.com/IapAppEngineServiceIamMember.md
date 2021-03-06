---
title: IapAppEngineServiceIamMember
menu:
  docs_{{ .version }}:
    identifier: iapappengineserviceiammember-google.kubeform.com
    name: IapAppEngineServiceIamMember
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IapAppEngineServiceIamMember
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `IapAppEngineServiceIamMember` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IapAppEngineServiceIamMemberSpec](#iapappengineserviceiammemberspec)***||
| `status` | ***[IapAppEngineServiceIamMemberStatus](#iapappengineserviceiammemberstatus)***||
## IapAppEngineServiceIamMemberSpec

Appears on:[IapAppEngineServiceIamMember](#iapappengineserviceiammember), [IapAppEngineServiceIamMemberStatus](#iapappengineserviceiammemberstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `member` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `service` | ***string***||
## IapAppEngineServiceIamMemberStatus

Appears on:[IapAppEngineServiceIamMember](#iapappengineserviceiammember)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IapAppEngineServiceIamMemberSpec](#iapappengineserviceiammemberspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IapAppEngineServiceIamMemberStatus](#iapappengineserviceiammemberstatus)

---
