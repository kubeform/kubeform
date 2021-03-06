---
title: ServiceAccountIamBinding
menu:
  docs_{{ .version }}:
    identifier: serviceaccountiambinding-google.kubeform.com
    name: ServiceAccountIamBinding
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServiceAccountIamBinding
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceAccountIamBinding` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceAccountIamBindingSpec](#serviceaccountiambindingspec)***||
| `status` | ***[ServiceAccountIamBindingStatus](#serviceaccountiambindingstatus)***||
## Phase(`string` alias)

Appears on:[ServiceAccountIamBindingStatus](#serviceaccountiambindingstatus)

## ServiceAccountIamBindingSpec

Appears on:[ServiceAccountIamBinding](#serviceaccountiambinding), [ServiceAccountIamBindingStatus](#serviceaccountiambindingstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `members` | ***[]string***||
| `role` | ***string***||
| `serviceAccountID` | ***string***||
## ServiceAccountIamBindingStatus

Appears on:[ServiceAccountIamBinding](#serviceaccountiambinding)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceAccountIamBindingSpec](#serviceaccountiambindingspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
