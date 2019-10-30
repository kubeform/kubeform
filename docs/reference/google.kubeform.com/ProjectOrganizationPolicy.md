---
title: ProjectOrganizationPolicy
menu:
  docs_{{ .version }}:
    identifier: projectorganizationpolicy-google.kubeform.com
    name: ProjectOrganizationPolicy
    parent: google.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## ProjectOrganizationPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `google.kubeform.com/v1alpha1` |
|    `kind` | string | `ProjectOrganizationPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[ProjectOrganizationPolicySpec](#projectorganizationpolicyspec)***||
| `status` | ***[ProjectOrganizationPolicyStatus](#projectorganizationpolicystatus)***||
## Phase(`string` alias)

Appears on:[ProjectOrganizationPolicyStatus](#projectorganizationpolicystatus)

## ProjectOrganizationPolicySpec

Appears on:[ProjectOrganizationPolicy](#projectorganizationpolicy), [ProjectOrganizationPolicyStatus](#projectorganizationpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `booleanPolicy` | ***[[]ProjectOrganizationPolicySpecBooleanPolicy](#projectorganizationpolicyspecbooleanpolicy)***| ***(Optional)*** |
| `constraint` | ***string***||
| `etag` | ***string***| ***(Optional)*** |
| `listPolicy` | ***[[]ProjectOrganizationPolicySpecListPolicy](#projectorganizationpolicyspeclistpolicy)***| ***(Optional)*** |
| `project` | ***string***||
| `restorePolicy` | ***[[]ProjectOrganizationPolicySpecRestorePolicy](#projectorganizationpolicyspecrestorepolicy)***| ***(Optional)*** |
| `updateTime` | ***string***| ***(Optional)*** |
| `version` | ***int***| ***(Optional)*** |
## ProjectOrganizationPolicySpecBooleanPolicy

Appears on:[ProjectOrganizationPolicySpec](#projectorganizationpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `enforced` | ***bool***||
## ProjectOrganizationPolicySpecListPolicy

Appears on:[ProjectOrganizationPolicySpec](#projectorganizationpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `allow` | ***[[]ProjectOrganizationPolicySpecListPolicyAllow](#projectorganizationpolicyspeclistpolicyallow)***| ***(Optional)*** |
| `deny` | ***[[]ProjectOrganizationPolicySpecListPolicyDeny](#projectorganizationpolicyspeclistpolicydeny)***| ***(Optional)*** |
| `suggestedValue` | ***string***| ***(Optional)*** |
## ProjectOrganizationPolicySpecListPolicyAllow

Appears on:[ProjectOrganizationPolicySpecListPolicy](#projectorganizationpolicyspeclistpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `all` | ***bool***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## ProjectOrganizationPolicySpecListPolicyDeny

Appears on:[ProjectOrganizationPolicySpecListPolicy](#projectorganizationpolicyspeclistpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `all` | ***bool***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## ProjectOrganizationPolicySpecRestorePolicy

Appears on:[ProjectOrganizationPolicySpec](#projectorganizationpolicyspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `default` | ***bool***||
## ProjectOrganizationPolicyStatus

Appears on:[ProjectOrganizationPolicy](#projectorganizationpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[ProjectOrganizationPolicySpec](#projectorganizationpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
