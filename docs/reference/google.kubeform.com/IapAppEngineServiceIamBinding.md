---
title: IapAppEngineServiceIamBinding
menu:
  docs_{{ .version }}:
    identifier: iapappengineserviceiambinding-google.kubeform.com
    name: IapAppEngineServiceIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IapAppEngineServiceIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `IapAppEngineServiceIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IapAppEngineServiceIamBindingSpec](#iapappengineserviceiambindingspec)***||
| `status` | ***[IapAppEngineServiceIamBindingStatus](#iapappengineserviceiambindingstatus)***||
## IapAppEngineServiceIamBindingSpec

Appears on:[IapAppEngineServiceIamBinding](#iapappengineserviceiambinding), [IapAppEngineServiceIamBindingStatus](#iapappengineserviceiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `project` | ***string***| ***(Optional)*** |
| `role` | ***string***||
| `service` | ***string***||
## IapAppEngineServiceIamBindingStatus

Appears on:[IapAppEngineServiceIamBinding](#iapappengineserviceiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IapAppEngineServiceIamBindingSpec](#iapappengineserviceiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IapAppEngineServiceIamBindingStatus](#iapappengineserviceiambindingstatus)

---
