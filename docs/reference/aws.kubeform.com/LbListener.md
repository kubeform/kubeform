---
title: LbListener
menu:
  docs_{{ .version }}:
    identifier: lblistener-aws.kubeform.com
    name: LbListener
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LbListener
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbListener` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbListenerSpec](#LbListenerSpec)***||
| `status` | ***[LbListenerStatus](#LbListenerStatus)***||
## LbListenerSpec

Appears on:[LbListener](#LbListener), [LbListenerStatus](#LbListenerStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `certificateArn` | ***string***| ***(Optional)*** |
| `defaultAction` | ***[[]LbListenerSpecDefaultAction](#LbListenerSpecDefaultAction)***||
| `loadBalancerArn` | ***string***||
| `port` | ***int***||
| `protocol` | ***string***| ***(Optional)*** |
| `sslPolicy` | ***string***| ***(Optional)*** |
## LbListenerSpecDefaultAction

Appears on:[LbListenerSpec](#LbListenerSpec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticateCognito` | ***[[]LbListenerSpecDefaultActionAuthenticateCognito](#LbListenerSpecDefaultActionAuthenticateCognito)***| ***(Optional)*** |
| `authenticateOidc` | ***[[]LbListenerSpecDefaultActionAuthenticateOidc](#LbListenerSpecDefaultActionAuthenticateOidc)***| ***(Optional)*** |
| `fixedResponse` | ***[[]LbListenerSpecDefaultActionFixedResponse](#LbListenerSpecDefaultActionFixedResponse)***| ***(Optional)*** |
| `order` | ***int***| ***(Optional)*** |
| `redirect` | ***[[]LbListenerSpecDefaultActionRedirect](#LbListenerSpecDefaultActionRedirect)***| ***(Optional)*** |
| `targetGroupArn` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## LbListenerSpecDefaultActionAuthenticateCognito

Appears on:[LbListenerSpecDefaultAction](#LbListenerSpecDefaultAction)

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
## LbListenerSpecDefaultActionAuthenticateOidc

Appears on:[LbListenerSpecDefaultAction](#LbListenerSpecDefaultAction)

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
## LbListenerSpecDefaultActionFixedResponse

Appears on:[LbListenerSpecDefaultAction](#LbListenerSpecDefaultAction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***||
| `messageBody` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***| ***(Optional)*** |
## LbListenerSpecDefaultActionRedirect

Appears on:[LbListenerSpecDefaultAction](#LbListenerSpecDefaultAction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `query` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## LbListenerStatus

Appears on:[LbListener](#LbListener)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbListenerSpec](#LbListenerSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `default_action.<index>.authenticate_oidc.<index>.client_secret` | ***string*** ||
