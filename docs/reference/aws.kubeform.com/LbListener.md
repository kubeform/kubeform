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
| `spec` | ***[LbListenerSpec](#lblistenerspec)***||
| `status` | ***[LbListenerStatus](#lblistenerstatus)***||
## LbListenerSpec

Appears on:[LbListener](#lblistener), [LbListenerStatus](#lblistenerstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `arn` | ***string***| ***(Optional)*** |
| `certificateArn` | ***string***| ***(Optional)*** |
| `defaultAction` | ***[[]LbListenerSpecDefaultAction](#lblistenerspecdefaultaction)***||
| `loadBalancerArn` | ***string***||
| `port` | ***int64***||
| `protocol` | ***string***| ***(Optional)*** |
| `sslPolicy` | ***string***| ***(Optional)*** |
## LbListenerSpecDefaultAction

Appears on:[LbListenerSpec](#lblistenerspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticateCognito` | ***[[]LbListenerSpecDefaultActionAuthenticateCognito](#lblistenerspecdefaultactionauthenticatecognito)***| ***(Optional)*** |
| `authenticateOidc` | ***[[]LbListenerSpecDefaultActionAuthenticateOidc](#lblistenerspecdefaultactionauthenticateoidc)***| ***(Optional)*** |
| `fixedResponse` | ***[[]LbListenerSpecDefaultActionFixedResponse](#lblistenerspecdefaultactionfixedresponse)***| ***(Optional)*** |
| `order` | ***int64***| ***(Optional)*** |
| `redirect` | ***[[]LbListenerSpecDefaultActionRedirect](#lblistenerspecdefaultactionredirect)***| ***(Optional)*** |
| `targetGroupArn` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## LbListenerSpecDefaultActionAuthenticateCognito

Appears on:[LbListenerSpecDefaultAction](#lblistenerspecdefaultaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationRequestExtraParams` | ***map[string]string***| ***(Optional)*** |
| `onUnauthenticatedRequest` | ***string***| ***(Optional)*** |
| `scope` | ***string***| ***(Optional)*** |
| `sessionCookieName` | ***string***| ***(Optional)*** |
| `sessionTimeout` | ***int64***| ***(Optional)*** |
| `userPoolArn` | ***string***||
| `userPoolClientID` | ***string***||
| `userPoolDomain` | ***string***||
## LbListenerSpecDefaultActionAuthenticateOidc

Appears on:[LbListenerSpecDefaultAction](#lblistenerspecdefaultaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticationRequestExtraParams` | ***map[string]string***| ***(Optional)*** |
| `authorizationEndpoint` | ***string***||
| `clientID` | ***string***||
| `issuer` | ***string***||
| `onUnauthenticatedRequest` | ***string***| ***(Optional)*** |
| `scope` | ***string***| ***(Optional)*** |
| `sessionCookieName` | ***string***| ***(Optional)*** |
| `sessionTimeout` | ***int64***| ***(Optional)*** |
| `tokenEndpoint` | ***string***||
| `userInfoEndpoint` | ***string***||
## LbListenerSpecDefaultActionFixedResponse

Appears on:[LbListenerSpecDefaultAction](#lblistenerspecdefaultaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***||
| `messageBody` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***| ***(Optional)*** |
## LbListenerSpecDefaultActionRedirect

Appears on:[LbListenerSpecDefaultAction](#lblistenerspecdefaultaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `query` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## LbListenerStatus

Appears on:[LbListener](#lblistener)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbListenerSpec](#lblistenerspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LbListenerStatus](#lblistenerstatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `default_action.<index>.authenticate_oidc.<index>.client_secret` | ***string*** ||
