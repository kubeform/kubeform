---
title: IamAccountPasswordPolicy
menu:
  docs_{{ .version }}:
    identifier: iamaccountpasswordpolicy-aws.kubeform.com
    name: IamAccountPasswordPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamAccountPasswordPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamAccountPasswordPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamAccountPasswordPolicySpec](#iamaccountpasswordpolicyspec)***||
| `status` | ***[IamAccountPasswordPolicyStatus](#iamaccountpasswordpolicystatus)***||
## IamAccountPasswordPolicySpec

Appears on:[IamAccountPasswordPolicy](#iamaccountpasswordpolicy), [IamAccountPasswordPolicyStatus](#iamaccountpasswordpolicystatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowUsersToChangePassword` | ***bool***| ***(Optional)*** |
| `expirePasswords` | ***bool***| ***(Optional)*** |
| `hardExpiry` | ***bool***| ***(Optional)*** |
| `maxPasswordAge` | ***int***| ***(Optional)*** |
| `minimumPasswordLength` | ***int***| ***(Optional)*** |
| `passwordReusePrevention` | ***int***| ***(Optional)*** |
| `requireLowercaseCharacters` | ***bool***| ***(Optional)*** |
| `requireNumbers` | ***bool***| ***(Optional)*** |
| `requireSymbols` | ***bool***| ***(Optional)*** |
| `requireUppercaseCharacters` | ***bool***| ***(Optional)*** |
## IamAccountPasswordPolicyStatus

Appears on:[IamAccountPasswordPolicy](#iamaccountpasswordpolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamAccountPasswordPolicySpec](#iamaccountpasswordpolicyspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IamAccountPasswordPolicyStatus](#iamaccountpasswordpolicystatus)

---
