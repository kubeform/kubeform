---
title: IamAccountPasswordPolicy
menu:
  docs_v0.0.1:
    identifier: iamaccountpasswordpolicy-aws.kubeform.com
    name: IamAccountPasswordPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IamAccountPasswordPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamAccountPasswordPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamAccountPasswordPolicySpec](#IamAccountPasswordPolicySpec)***||
| `status` | ***[IamAccountPasswordPolicyStatus](#IamAccountPasswordPolicyStatus)***||
## IamAccountPasswordPolicySpec
##### (Appears on:[IamAccountPasswordPolicy](#IamAccountPasswordPolicy), [IamAccountPasswordPolicyStatus](#IamAccountPasswordPolicyStatus))
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
##### (Appears on:[IamAccountPasswordPolicy](#IamAccountPasswordPolicy))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamAccountPasswordPolicySpec](#IamAccountPasswordPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
