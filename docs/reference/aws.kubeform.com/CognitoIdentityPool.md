---
title: CognitoIdentityPool
menu:
  docs_{{ .version }}:
    identifier: cognitoidentitypool-aws.kubeform.com
    name: CognitoIdentityPool
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CognitoIdentityPool
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitoIdentityPool` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitoIdentityPoolSpec](#CognitoIdentityPoolSpec)***||
| `status` | ***[CognitoIdentityPoolStatus](#CognitoIdentityPoolStatus)***||
## CognitoIdentityPoolSpec

Appears on:[CognitoIdentityPool](#CognitoIdentityPool), [CognitoIdentityPoolStatus](#CognitoIdentityPoolStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowUnauthenticatedIdentities` | ***bool***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `cognitoIdentityProviders` | ***[[]CognitoIdentityPoolSpecCognitoIdentityProviders](#CognitoIdentityPoolSpecCognitoIdentityProviders)***| ***(Optional)*** |
| `developerProviderName` | ***string***| ***(Optional)*** |
| `identityPoolName` | ***string***||
| `openidConnectProviderArns` | ***[]string***| ***(Optional)*** |
| `samlProviderArns` | ***[]string***| ***(Optional)*** |
| `supportedLoginProviders` | ***map[string]string***| ***(Optional)*** |
## CognitoIdentityPoolSpecCognitoIdentityProviders

Appears on:[CognitoIdentityPoolSpec](#CognitoIdentityPoolSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `clientID` | ***string***| ***(Optional)*** |
| `providerName` | ***string***| ***(Optional)*** |
| `serverSideTokenCheck` | ***bool***| ***(Optional)*** |
## CognitoIdentityPoolStatus

Appears on:[CognitoIdentityPool](#CognitoIdentityPool)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitoIdentityPoolSpec](#CognitoIdentityPoolSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
