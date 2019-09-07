---
title: LbListenerRule
menu:
  docs_v0.0.1:
    identifier: lblistenerrule-aws.kubeform.com
    name: LbListenerRule
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## LbListenerRule
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbListenerRule` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbListenerRuleSpec](#LbListenerRuleSpec)***||
| `status` | ***[LbListenerRuleStatus](#LbListenerRuleStatus)***||
## LbListenerRuleSpec
##### (Appears on:[LbListenerRule](#LbListenerRule), [LbListenerRuleStatus](#LbListenerRuleStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `secretRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `action` | ***[[]LbListenerRuleSpecAction](#LbListenerRuleSpecAction)***||
| `arn` | ***string***| ***(Optional)*** |
| `condition` | ***[[]LbListenerRuleSpecCondition](#LbListenerRuleSpecCondition)***||
| `listenerArn` | ***string***||
| `priority` | ***int***| ***(Optional)*** |
## LbListenerRuleSpecAction
##### (Appears on:[LbListenerRuleSpec](#LbListenerRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `authenticateCognito` | ***[[]LbListenerRuleSpecActionAuthenticateCognito](#LbListenerRuleSpecActionAuthenticateCognito)***| ***(Optional)*** |
| `authenticateOidc` | ***[[]LbListenerRuleSpecActionAuthenticateOidc](#LbListenerRuleSpecActionAuthenticateOidc)***| ***(Optional)*** |
| `fixedResponse` | ***[[]LbListenerRuleSpecActionFixedResponse](#LbListenerRuleSpecActionFixedResponse)***| ***(Optional)*** |
| `order` | ***int***| ***(Optional)*** |
| `redirect` | ***[[]LbListenerRuleSpecActionRedirect](#LbListenerRuleSpecActionRedirect)***| ***(Optional)*** |
| `targetGroupArn` | ***string***| ***(Optional)*** |
| `type` | ***string***||
## LbListenerRuleSpecActionAuthenticateCognito
##### (Appears on:[LbListenerRuleSpecAction](#LbListenerRuleSpecAction))
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
## LbListenerRuleSpecActionAuthenticateOidc
##### (Appears on:[LbListenerRuleSpecAction](#LbListenerRuleSpecAction))
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
## LbListenerRuleSpecActionFixedResponse
##### (Appears on:[LbListenerRuleSpecAction](#LbListenerRuleSpecAction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `contentType` | ***string***||
| `messageBody` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***| ***(Optional)*** |
## LbListenerRuleSpecActionRedirect
##### (Appears on:[LbListenerRuleSpecAction](#LbListenerRuleSpecAction))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `host` | ***string***| ***(Optional)*** |
| `path` | ***string***| ***(Optional)*** |
| `port` | ***string***| ***(Optional)*** |
| `protocol` | ***string***| ***(Optional)*** |
| `query` | ***string***| ***(Optional)*** |
| `statusCode` | ***string***||
## LbListenerRuleSpecCondition
##### (Appears on:[LbListenerRuleSpec](#LbListenerRuleSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `field` | ***string***| ***(Optional)*** |
| `values` | ***[]string***| ***(Optional)*** |
## LbListenerRuleStatus
##### (Appears on:[LbListenerRule](#LbListenerRule))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbListenerRuleSpec](#LbListenerRuleSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
## Sensitive Values
| Name | Type | Description |
|------|------|-------------|
| `action.<index>.authenticate_oidc.<index>.client_secret` | ***string*** ||
