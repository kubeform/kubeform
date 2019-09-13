---
title: SsmPatchBaseline
menu:
  docs_{{ .version }}:
    identifier: ssmpatchbaseline-aws.kubeform.com
    name: SsmPatchBaseline
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## SsmPatchBaseline
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `SsmPatchBaseline` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[SsmPatchBaselineSpec](#SsmPatchBaselineSpec)***||
| `status` | ***[SsmPatchBaselineStatus](#SsmPatchBaselineStatus)***||
## SsmPatchBaselineSpec
##### (Appears on:[SsmPatchBaseline](#SsmPatchBaseline), [SsmPatchBaselineStatus](#SsmPatchBaselineStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `approvalRule` | ***[[]SsmPatchBaselineSpecApprovalRule](#SsmPatchBaselineSpecApprovalRule)***| ***(Optional)*** |
| `approvedPatches` | ***[]string***| ***(Optional)*** |
| `approvedPatchesComplianceLevel` | ***string***| ***(Optional)*** |
| `description` | ***string***| ***(Optional)*** |
| `globalFilter` | ***[[]SsmPatchBaselineSpecGlobalFilter](#SsmPatchBaselineSpecGlobalFilter)***| ***(Optional)*** |
| `name` | ***string***||
| `operatingSystem` | ***string***| ***(Optional)*** |
| `rejectedPatches` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
## SsmPatchBaselineSpecApprovalRule
##### (Appears on:[SsmPatchBaselineSpec](#SsmPatchBaselineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `approveAfterDays` | ***int***||
| `complianceLevel` | ***string***| ***(Optional)*** |
| `enableNonSecurity` | ***bool***| ***(Optional)*** |
| `patchFilter` | ***[[]SsmPatchBaselineSpecApprovalRulePatchFilter](#SsmPatchBaselineSpecApprovalRulePatchFilter)***||
## SsmPatchBaselineSpecApprovalRulePatchFilter
##### (Appears on:[SsmPatchBaselineSpecApprovalRule](#SsmPatchBaselineSpecApprovalRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `values` | ***[]string***||
## SsmPatchBaselineSpecGlobalFilter
##### (Appears on:[SsmPatchBaselineSpec](#SsmPatchBaselineSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `key` | ***string***||
| `values` | ***[]string***||
## SsmPatchBaselineStatus
##### (Appears on:[SsmPatchBaseline](#SsmPatchBaseline))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[SsmPatchBaselineSpec](#SsmPatchBaselineSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
