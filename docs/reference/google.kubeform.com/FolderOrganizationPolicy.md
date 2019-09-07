---
title: FolderOrganizationPolicy
menu:
  docs_v0.0.1:
    identifier: folderorganizationpolicy-google.kubeform.com
    name: FolderOrganizationPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## FolderOrganizationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `FolderOrganizationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[FolderOrganizationPolicySpec](#FolderOrganizationPolicySpec)***||
| `status` | ***[FolderOrganizationPolicyStatus](#FolderOrganizationPolicyStatus)***||
## FolderOrganizationPolicySpec
##### (Appears on:[FolderOrganizationPolicy](#FolderOrganizationPolicy), [FolderOrganizationPolicyStatus](#FolderOrganizationPolicyStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `booleanPolicy` | ***[[]FolderOrganizationPolicySpecBooleanPolicy](#FolderOrganizationPolicySpecBooleanPolicy)***| ***(Optional)*** |
| `constraint` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `folder` | ***string***||
| `listPolicy` | ***[[]FolderOrganizationPolicySpecListPolicy](#FolderOrganizationPolicySpecListPolicy)***| ***(Optional)*** |
| `restorePolicy` | ***[[]FolderOrganizationPolicySpecRestorePolicy](#FolderOrganizationPolicySpecRestorePolicy)***| ***(Optional)*** |
| `updateTime` | ***string***| ***(Optional)*** |
| `version` | ***int***| ***(Optional)*** |
## FolderOrganizationPolicySpecBooleanPolicy
##### (Appears on:[FolderOrganizationPolicySpec](#FolderOrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `enforced` | ***bool***||
## FolderOrganizationPolicySpecListPolicy
##### (Appears on:[FolderOrganizationPolicySpec](#FolderOrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `allow` | ***[[]FolderOrganizationPolicySpecListPolicyAllow](#FolderOrganizationPolicySpecListPolicyAllow)***| ***(Optional)*** |
| `deny` | ***[[]FolderOrganizationPolicySpecListPolicyDeny](#FolderOrganizationPolicySpecListPolicyDeny)***| ***(Optional)*** |
| `suggestedValue` | ***string***| ***(Optional)*** |
## FolderOrganizationPolicySpecListPolicyAllow
##### (Appears on:[FolderOrganizationPolicySpecListPolicy](#FolderOrganizationPolicySpecListPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `all` | ***bool***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## FolderOrganizationPolicySpecListPolicyDeny
##### (Appears on:[FolderOrganizationPolicySpecListPolicy](#FolderOrganizationPolicySpecListPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `all` | ***bool***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## FolderOrganizationPolicySpecRestorePolicy
##### (Appears on:[FolderOrganizationPolicySpec](#FolderOrganizationPolicySpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `default` | ***bool***||
## FolderOrganizationPolicyStatus
##### (Appears on:[FolderOrganizationPolicy](#FolderOrganizationPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[FolderOrganizationPolicySpec](#FolderOrganizationPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
