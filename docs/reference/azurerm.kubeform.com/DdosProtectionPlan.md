---
title: DdosProtectionPlan
menu:
  docs_{{ .version }}:
    identifier: ddosprotectionplan-azurerm.kubeform.com
    name: DdosProtectionPlan
    parent: azurerm.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DdosProtectionPlan
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `azurerm.kubeform.com/v1alpha1` |
|    `kind` | string | `DdosProtectionPlan` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DdosProtectionPlanSpec](#ddosprotectionplanspec)***||
| `status` | ***[DdosProtectionPlanStatus](#ddosprotectionplanstatus)***||
## DdosProtectionPlanSpec

Appears on:[DdosProtectionPlan](#ddosprotectionplan), [DdosProtectionPlanStatus](#ddosprotectionplanstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `location` | ***string***||
| `name` | ***string***||
| `resourceGroupName` | ***string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `virtualNetworkIDS` | ***[]string***| ***(Optional)*** |
## DdosProtectionPlanStatus

Appears on:[DdosProtectionPlan](#ddosprotectionplan)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DdosProtectionPlanSpec](#ddosprotectionplanspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DdosProtectionPlanStatus](#ddosprotectionplanstatus)

---
