---
title: ServiceAccount
menu:
  docs_{{ .version }}:
    identifier: serviceaccount-google.kubeform.com
    name: ServiceAccount
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ServiceAccount
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ServiceAccount` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ServiceAccountSpec](#ServiceAccountSpec)***||
| `status` | ***[ServiceAccountStatus](#ServiceAccountStatus)***||
## ServiceAccountSpec

Appears on:[ServiceAccount](#ServiceAccount), [ServiceAccountStatus](#ServiceAccountStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `accountID` | ***string***||
| `displayName` | ***string***| ***(Optional)*** |
| `email` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `policyData` | ***string***| ***(Optional)*** Deprecated|
| `project` | ***string***| ***(Optional)*** |
| `uniqueID` | ***string***| ***(Optional)*** |
## ServiceAccountStatus

Appears on:[ServiceAccount](#ServiceAccount)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ServiceAccountSpec](#ServiceAccountSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
