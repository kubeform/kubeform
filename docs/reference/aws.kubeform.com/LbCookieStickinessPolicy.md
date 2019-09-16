---
title: LbCookieStickinessPolicy
menu:
  docs_{{ .version }}:
    identifier: lbcookiestickinesspolicy-aws.kubeform.com
    name: LbCookieStickinessPolicy
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LbCookieStickinessPolicy
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LbCookieStickinessPolicy` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LbCookieStickinessPolicySpec](#LbCookieStickinessPolicySpec)***||
| `status` | ***[LbCookieStickinessPolicyStatus](#LbCookieStickinessPolicyStatus)***||
## LbCookieStickinessPolicySpec

Appears on:[LbCookieStickinessPolicy](#LbCookieStickinessPolicy), [LbCookieStickinessPolicyStatus](#LbCookieStickinessPolicyStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cookieExpirationPeriod` | ***int***| ***(Optional)*** |
| `lbPort` | ***int***||
| `loadBalancer` | ***string***||
| `name` | ***string***||
## LbCookieStickinessPolicyStatus

Appears on:[LbCookieStickinessPolicy](#LbCookieStickinessPolicy)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LbCookieStickinessPolicySpec](#LbCookieStickinessPolicySpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
