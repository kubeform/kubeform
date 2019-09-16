---
title: CognitoIdentityProvider
menu:
  docs_{{ .version }}:
    identifier: cognitoidentityprovider-aws.kubeform.com
    name: CognitoIdentityProvider
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CognitoIdentityProvider
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitoIdentityProvider` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitoIdentityProviderSpec](#CognitoIdentityProviderSpec)***||
| `status` | ***[CognitoIdentityProviderStatus](#CognitoIdentityProviderStatus)***||
## CognitoIdentityProviderSpec

Appears on:[CognitoIdentityProvider](#CognitoIdentityProvider), [CognitoIdentityProviderStatus](#CognitoIdentityProviderStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `attributeMapping` | ***map[string]string***| ***(Optional)*** |
| `idpIdentifiers` | ***[]string***| ***(Optional)*** |
| `providerDetails` | ***map[string]string***||
| `providerName` | ***string***||
| `providerType` | ***string***||
| `userPoolID` | ***string***||
## CognitoIdentityProviderStatus

Appears on:[CognitoIdentityProvider](#CognitoIdentityProvider)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitoIdentityProviderSpec](#CognitoIdentityProviderSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
