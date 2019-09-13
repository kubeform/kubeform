---
title: LightsailDomain
menu:
  docs_{{ .version }}:
    identifier: lightsaildomain-aws.kubeform.com
    name: LightsailDomain
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## LightsailDomain
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `LightsailDomain` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[LightsailDomainSpec](#LightsailDomainSpec)***||
| `status` | ***[LightsailDomainStatus](#LightsailDomainStatus)***||
## LightsailDomainSpec
##### (Appears on:[LightsailDomain](#LightsailDomain), [LightsailDomainStatus](#LightsailDomainStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `domainName` | ***string***||
## LightsailDomainStatus
##### (Appears on:[LightsailDomain](#LightsailDomain))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[LightsailDomainSpec](#LightsailDomainSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
