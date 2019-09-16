---
title: CognitoUserPoolClient
menu:
  docs_{{ .version }}:
    identifier: cognitouserpoolclient-aws.kubeform.com
    name: CognitoUserPoolClient
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CognitoUserPoolClient
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitoUserPoolClient` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitoUserPoolClientSpec](#CognitoUserPoolClientSpec)***||
| `status` | ***[CognitoUserPoolClientStatus](#CognitoUserPoolClientStatus)***||
## CognitoUserPoolClientSpec

Appears on:[CognitoUserPoolClient](#CognitoUserPoolClient), [CognitoUserPoolClientStatus](#CognitoUserPoolClientStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `allowedOauthFlows` | ***[]string***| ***(Optional)*** |
| `allowedOauthFlowsUserPoolClient` | ***bool***| ***(Optional)*** |
| `allowedOauthScopes` | ***[]string***| ***(Optional)*** |
| `callbackUrls` | ***[]string***| ***(Optional)*** |
| `clientSecret` | ***string***| ***(Optional)*** |
| `defaultRedirectURI` | ***string***| ***(Optional)*** |
| `explicitAuthFlows` | ***[]string***| ***(Optional)*** |
| `generateSecret` | ***bool***| ***(Optional)*** |
| `logoutUrls` | ***[]string***| ***(Optional)*** |
| `name` | ***string***||
| `readAttributes` | ***[]string***| ***(Optional)*** |
| `refreshTokenValidity` | ***int***| ***(Optional)*** |
| `supportedIdentityProviders` | ***[]string***| ***(Optional)*** |
| `userPoolID` | ***string***||
| `writeAttributes` | ***[]string***| ***(Optional)*** |
## CognitoUserPoolClientStatus

Appears on:[CognitoUserPoolClient](#CognitoUserPoolClient)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitoUserPoolClientSpec](#CognitoUserPoolClientSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
