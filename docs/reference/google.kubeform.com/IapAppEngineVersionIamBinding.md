---
title: IapAppEngineVersionIamBinding
menu:
  docs_{{ .version }}:
    identifier: iapappengineversioniambinding-google.kubeform.com
    name: IapAppEngineVersionIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IapAppEngineVersionIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `IapAppEngineVersionIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IapAppEngineVersionIamBindingSpec](#iapappengineversioniambindingspec)***||
| `status` | ***[IapAppEngineVersionIamBindingStatus](#iapappengineversioniambindingstatus)***||
## IapAppEngineVersionIamBindingSpec

Appears on:[IapAppEngineVersionIamBinding](#iapappengineversioniambinding), [IapAppEngineVersionIamBindingStatus](#iapappengineversioniambindingstatus)

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
| `versionID` | ***string***||
## IapAppEngineVersionIamBindingStatus

Appears on:[IapAppEngineVersionIamBinding](#iapappengineversioniambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IapAppEngineVersionIamBindingSpec](#iapappengineversioniambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IapAppEngineVersionIamBindingStatus](#iapappengineversioniambindingstatus)

---
