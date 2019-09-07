---
title: CognitoIdentityPoolRolesAttachment
menu:
  docs_v0.0.1:
    identifier: cognitoidentitypoolrolesattachment-aws.kubeform.com
    name: CognitoIdentityPoolRolesAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## CognitoIdentityPoolRolesAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitoIdentityPoolRolesAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitoIdentityPoolRolesAttachmentSpec](#CognitoIdentityPoolRolesAttachmentSpec)***||
| `status` | ***[CognitoIdentityPoolRolesAttachmentStatus](#CognitoIdentityPoolRolesAttachmentStatus)***||
## CognitoIdentityPoolRolesAttachmentSpec
##### (Appears on:[CognitoIdentityPoolRolesAttachment](#CognitoIdentityPoolRolesAttachment), [CognitoIdentityPoolRolesAttachmentStatus](#CognitoIdentityPoolRolesAttachmentStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `identityPoolID` | ***string***||
| `roleMapping` | ***[[]CognitoIdentityPoolRolesAttachmentSpecRoleMapping](#CognitoIdentityPoolRolesAttachmentSpecRoleMapping)***| ***(Optional)*** |
| `roles` | ***map[string]kubeform.dev/kubeform/apis/aws/v1alpha1.CognitoIdentityPoolRolesAttachmentSpecRoles***||
## CognitoIdentityPoolRolesAttachmentSpecRoleMapping
##### (Appears on:[CognitoIdentityPoolRolesAttachmentSpec](#CognitoIdentityPoolRolesAttachmentSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `ambiguousRoleResolution` | ***string***| ***(Optional)*** |
| `identityProvider` | ***string***||
| `mappingRule` | ***[[]CognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule](#CognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule)***| ***(Optional)*** |
| `type` | ***string***||
## CognitoIdentityPoolRolesAttachmentSpecRoleMappingMappingRule
##### (Appears on:[CognitoIdentityPoolRolesAttachmentSpecRoleMapping](#CognitoIdentityPoolRolesAttachmentSpecRoleMapping))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `claim` | ***string***||
| `matchType` | ***string***||
| `roleArn` | ***string***||
| `value` | ***string***||
## CognitoIdentityPoolRolesAttachmentStatus
##### (Appears on:[CognitoIdentityPoolRolesAttachment](#CognitoIdentityPoolRolesAttachment))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitoIdentityPoolRolesAttachmentSpec](#CognitoIdentityPoolRolesAttachmentSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
