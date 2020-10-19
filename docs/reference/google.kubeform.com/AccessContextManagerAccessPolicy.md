---
title: AccessContextManagerAccessPolicy
menu:
  docs_{{ .version }}:
    identifier: accesscontextmanageraccesspolicy-google.kubeform.com
    name: AccessContextManagerAccessPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AccessContextManagerAccessPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `AccessContextManagerAccessPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AccessContextManagerAccessPolicySpec](#accesscontextmanageraccesspolicyspec)***||
| `status` | ***[AccessContextManagerAccessPolicyStatus](#accesscontextmanageraccesspolicystatus)***||
## AccessContextManagerAccessPolicySpec

Appears on:[AccessContextManagerAccessPolicy](#accesscontextmanageraccesspolicy), [AccessContextManagerAccessPolicyStatus](#accesscontextmanageraccesspolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `createTime` | ***string***| ***(Optional)*** |
| `name` | ***string***| ***(Optional)*** |
| `parent` | ***string***||
| `title` | ***string***||
| `updateTime` | ***string***| ***(Optional)*** |
## AccessContextManagerAccessPolicyStatus

Appears on:[AccessContextManagerAccessPolicy](#accesscontextmanageraccesspolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AccessContextManagerAccessPolicySpec](#accesscontextmanageraccesspolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[AccessContextManagerAccessPolicyStatus](#accesscontextmanageraccesspolicystatus)

---
