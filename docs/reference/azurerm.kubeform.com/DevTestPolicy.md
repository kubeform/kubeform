---
title: DevTestPolicy
menu:
  docs_{{ .version }}:
    identifier: devtestpolicy-azurerm.kubeform.com
    name: DevTestPolicy
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DevTestPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DevTestPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DevTestPolicySpec](#devtestpolicyspec)***||
| `status` | ***[DevTestPolicyStatus](#devtestpolicystatus)***||
## DevTestPolicySpec

Appears on:[DevTestPolicy](#devtestpolicy), [DevTestPolicyStatus](#devtestpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `description` | ***string***| ***(Optional)*** |
| `evaluatorType` | ***string***||
| `factData` | ***string***| ***(Optional)*** |
| `labName` | ***string***||
| `name` | ***string***||
| `policySetName` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `threshold` | ***string***||
## DevTestPolicyStatus

Appears on:[DevTestPolicy](#devtestpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DevTestPolicySpec](#devtestpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DevTestPolicyStatus](#devtestpolicystatus)

---
