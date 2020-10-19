---
title: IapAppEngineServiceIamPolicy
menu:
  docs_{{ .version }}:
    identifier: iapappengineserviceiampolicy-google.kubeform.com
    name: IapAppEngineServiceIamPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IapAppEngineServiceIamPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `IapAppEngineServiceIamPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IapAppEngineServiceIamPolicySpec](#iapappengineserviceiampolicyspec)***||
| `status` | ***[IapAppEngineServiceIamPolicyStatus](#iapappengineserviceiampolicystatus)***||
## IapAppEngineServiceIamPolicySpec

Appears on:[IapAppEngineServiceIamPolicy](#iapappengineserviceiampolicy), [IapAppEngineServiceIamPolicyStatus](#iapappengineserviceiampolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `appID` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `policyData` | ***string***||
| `project` | ***string***| ***(Optional)*** |
| `service` | ***string***||
## IapAppEngineServiceIamPolicyStatus

Appears on:[IapAppEngineServiceIamPolicy](#iapappengineserviceiampolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IapAppEngineServiceIamPolicySpec](#iapappengineserviceiampolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IapAppEngineServiceIamPolicyStatus](#iapappengineserviceiampolicystatus)

---
