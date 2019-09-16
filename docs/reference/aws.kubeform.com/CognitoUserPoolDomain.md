---
title: CognitoUserPoolDomain
menu:
  docs_{{ .version }}:
    identifier: cognitouserpooldomain-aws.kubeform.com
    name: CognitoUserPoolDomain
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## CognitoUserPoolDomain
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `CognitoUserPoolDomain` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[CognitoUserPoolDomainSpec](#CognitoUserPoolDomainSpec)***||
| `status` | ***[CognitoUserPoolDomainStatus](#CognitoUserPoolDomainStatus)***||
## CognitoUserPoolDomainSpec

Appears on:[CognitoUserPoolDomain](#CognitoUserPoolDomain), [CognitoUserPoolDomainStatus](#CognitoUserPoolDomainStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `awsAccountID` | ***string***| ***(Optional)*** |
| `certificateArn` | ***string***| ***(Optional)*** |
| `cloudfrontDistributionArn` | ***string***| ***(Optional)*** |
| `domain` | ***string***||
| `s3Bucket` | ***string***| ***(Optional)*** |
| `userPoolID` | ***string***||
| `version` | ***string***| ***(Optional)*** |
## CognitoUserPoolDomainStatus

Appears on:[CognitoUserPoolDomain](#CognitoUserPoolDomain)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[CognitoUserPoolDomainSpec](#CognitoUserPoolDomainSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
