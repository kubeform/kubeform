---
title: AlbListener
menu:
  docs_{{ .version }}:
    identifier: alblistener-aws.kubeform.com
    name: AlbListener
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## AlbListener
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `AlbListener` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[AlbListenerSpec](#alblistenerspec)***||
| `status` | ***[AlbListenerStatus](#alblistenerstatus)***||
## AlbListenerSpec

Appears on:[AlbListener](#alblistener), [AlbListenerStatus](#alblistenerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `certificateArn` | ***string***| ***(Optional)*** |
| `defaultAction` | ***[[]AlbListenerSpecDefaultAction](#alblistenerspecdefaultaction)***||
| `loadBalancerArn` | ***string***||
| `port` | ***int***||
| `protocol` | ***string***| ***(Optional)*** |
| `sslPolicy` | ***string***| ***(Optional)*** |
## AlbListenerSpecDefaultAction

Appears on:[AlbListenerSpec](#alblistenerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticateCognito` | ***[[]AlbListenerSpecDefaultActionAuthenticateCognito](#alblistenerspecdefaultactionauthenticatecognito)***| ***(Optional)*** |
| `authenticateOidc` | ***[[]AlbListenerSpecDefaultActionAuthenticateOidc](#alblistenerspecdefaultactionauthenticateoidc)***| ***(Optional)*** |
| `fixedResponse` | ***[[]AlbListenerSpecDefaultActionFixedResponse](#alblistenerspecdefaultactionfixedresponse)***| ***(Optional)*** |
| `order` | ***int***| ***(Optional)*** |
| `redirect` | ***[[]AlbListenerSpecDefaultActionRedirect](#alblistenerspecdefaultactionredirect)***| ***(Optional)*** |
| `targetGroupArn` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## AlbListenerSpecDefaultActionAuthenticateCognito

Appears on:[AlbListenerSpecDefaultAction](#alblistenerspecdefaultaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationRequestExtraParams` | ***map[string]string***| ***(Optional)*** |
| `onUnauthenticatedRequest` | ***string***| ***(Optional)*** |
| `scope` | ***string***| ***(Optional)*** |
| `sessionCookieName` | ***string***| ***(Optional)*** |
| `sessionTimeout` | ***int***| ***(Optional)*** |
| `userPoolArn` | ***string***||
| `userPoolClientID` | ***string***||
| `userPoolDomain` | ***string***||
## AlbListenerSpecDefaultActionAuthenticateOidc

Appears on:[AlbListenerSpecDefaultAction](#alblistenerspecdefaultaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationRequestExtraParams` | ***map[string]string***| ***(Optional)*** |
| `authorizationEndpoint` | ***string***||
| `clientID` | ***string***||
| `issuer` | ***string***||
| `onUnauthenticatedRequest` | ***string***| ***(Optional)*** |
| `scope` | ***string***| ***(Optional)*** |
| `sessionCookieName` | ***string***| ***(Optional)*** |
| `sessionTimeout` | ***int***| ***(Optional)*** |
| `tokenEndpoint` | ***string***||
| `userInfoEndpoint` | ***string***||
## AlbListenerSpecDefaultActionFixedResponse

Appears on:[AlbListenerSpecDefaultAction](#alblistenerspecdefaultaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***||
| `messageBody` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***| ***(Optional)*** |
## AlbListenerSpecDefaultActionRedirect

Appears on:[AlbListenerSpecDefaultAction](#alblistenerspecdefaultaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `query` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## AlbListenerStatus

Appears on:[AlbListener](#alblistener)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[AlbListenerSpec](#alblistenerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `default_action.<index>.authenticate_oidc.<index>.client_secret` | ***string*** ||
