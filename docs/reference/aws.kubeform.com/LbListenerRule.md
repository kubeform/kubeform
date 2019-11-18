---
title: LbListenerRule
menu:
  docs_{{ .version }}:
    identifier: lblistenerrule-aws.kubeform.com
    name: LbListenerRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LbListenerRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbListenerRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbListenerRuleSpec](#lblistenerrulespec)***||
| `status` | ***[LbListenerRuleStatus](#lblistenerrulestatus)***||
## LbListenerRuleSpec

Appears on:[LbListenerRule](#lblistenerrule), [LbListenerRuleStatus](#lblistenerrulestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `action` | ***[[]LbListenerRuleSpecAction](#lblistenerrulespecaction)***||
| `arn` | ***string***| ***(Optional)*** |
| `condition` | ***[[]LbListenerRuleSpecCondition](#lblistenerrulespeccondition)***||
| `listenerArn` | ***string***||
| `priority` | ***int64***| ***(Optional)*** |
## LbListenerRuleSpecAction

Appears on:[LbListenerRuleSpec](#lblistenerrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticateCognito` | ***[[]LbListenerRuleSpecActionAuthenticateCognito](#lblistenerrulespecactionauthenticatecognito)***| ***(Optional)*** |
| `authenticateOidc` | ***[[]LbListenerRuleSpecActionAuthenticateOidc](#lblistenerrulespecactionauthenticateoidc)***| ***(Optional)*** |
| `fixedResponse` | ***[[]LbListenerRuleSpecActionFixedResponse](#lblistenerrulespecactionfixedresponse)***| ***(Optional)*** |
| `order` | ***int64***| ***(Optional)*** |
| `redirect` | ***[[]LbListenerRuleSpecActionRedirect](#lblistenerrulespecactionredirect)***| ***(Optional)*** |
| `targetGroupArn` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## LbListenerRuleSpecActionAuthenticateCognito

Appears on:[LbListenerRuleSpecAction](#lblistenerrulespecaction)

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
## LbListenerRuleSpecActionAuthenticateOidc

Appears on:[LbListenerRuleSpecAction](#lblistenerrulespecaction)

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
## LbListenerRuleSpecActionFixedResponse

Appears on:[LbListenerRuleSpecAction](#lblistenerrulespecaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***||
| `messageBody` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***| ***(Optional)*** |
## LbListenerRuleSpecActionRedirect

Appears on:[LbListenerRuleSpecAction](#lblistenerrulespecaction)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `query` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## LbListenerRuleSpecCondition

Appears on:[LbListenerRuleSpec](#lblistenerrulespec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `field` | ***string***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## LbListenerRuleStatus

Appears on:[LbListenerRule](#lblistenerrule)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbListenerRuleSpec](#lblistenerrulespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[LbListenerRuleStatus](#lblistenerrulestatus)

---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `action.<index>.authenticate_oidc.<index>.client_secret` | ***string*** ||
