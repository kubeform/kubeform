---
title: ProjectOrganizationPolicy
menu:
  docs_v0.0.1:
    identifier: projectorganizationpolicy-google.kubeform.com
    name: ProjectOrganizationPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## ProjectOrganizationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ProjectOrganizationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectOrganizationPolicySpec](#ProjectOrganizationPolicySpec)***||
| `status` | ***[ProjectOrganizationPolicyStatus](#ProjectOrganizationPolicyStatus)***||
## ProjectOrganizationPolicySpec
##### (Appears on:[ProjectOrganizationPolicy](#ProjectOrganizationPolicy), [ProjectOrganizationPolicyStatus](#ProjectOrganizationPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `booleanPolicy` | ***[[]ProjectOrganizationPolicySpecBooleanPolicy](#ProjectOrganizationPolicySpecBooleanPolicy)***| ***(Optional)*** |
| `constraint` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `listPolicy` | ***[[]ProjectOrganizationPolicySpecListPolicy](#ProjectOrganizationPolicySpecListPolicy)***| ***(Optional)*** |
| `project` | ***string***||
| `restorePolicy` | ***[[]ProjectOrganizationPolicySpecRestorePolicy](#ProjectOrganizationPolicySpecRestorePolicy)***| ***(Optional)*** |
| `updateTime` | ***string***| ***(Optional)*** |
| `version` | ***int***| ***(Optional)*** |
## ProjectOrganizationPolicySpecBooleanPolicy
##### (Appears on:[ProjectOrganizationPolicySpec](#ProjectOrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enforced` | ***bool***||
## ProjectOrganizationPolicySpecListPolicy
##### (Appears on:[ProjectOrganizationPolicySpec](#ProjectOrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allow` | ***[[]ProjectOrganizationPolicySpecListPolicyAllow](#ProjectOrganizationPolicySpecListPolicyAllow)***| ***(Optional)*** |
| `deny` | ***[[]ProjectOrganizationPolicySpecListPolicyDeny](#ProjectOrganizationPolicySpecListPolicyDeny)***| ***(Optional)*** |
| `suggestedValue` | ***string***| ***(Optional)*** |
## ProjectOrganizationPolicySpecListPolicyAllow
##### (Appears on:[ProjectOrganizationPolicySpecListPolicy](#ProjectOrganizationPolicySpecListPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `all` | ***bool***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## ProjectOrganizationPolicySpecListPolicyDeny
##### (Appears on:[ProjectOrganizationPolicySpecListPolicy](#ProjectOrganizationPolicySpecListPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `all` | ***bool***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## ProjectOrganizationPolicySpecRestorePolicy
##### (Appears on:[ProjectOrganizationPolicySpec](#ProjectOrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `default` | ***bool***||
## ProjectOrganizationPolicyStatus
##### (Appears on:[ProjectOrganizationPolicy](#ProjectOrganizationPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectOrganizationPolicySpec](#ProjectOrganizationPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
