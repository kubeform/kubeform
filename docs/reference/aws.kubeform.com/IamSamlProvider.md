---
title: IamSamlProvider
menu:
  docs_{{ .version }}:
    identifier: iamsamlprovider-aws.kubeform.com
    name: IamSamlProvider
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamSamlProvider
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamSamlProvider` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamSamlProviderSpec](#iamsamlproviderspec)***||
| `status` | ***[IamSamlProviderStatus](#iamsamlproviderstatus)***||
## IamSamlProviderSpec

Appears on:[IamSamlProvider](#iamsamlprovider), [IamSamlProviderStatus](#iamsamlproviderstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `samlMetadataDocument` | ***string***||
| `validUntil` | ***string***| ***(Optional)*** |
## IamSamlProviderStatus

Appears on:[IamSamlProvider](#iamsamlprovider)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamSamlProviderSpec](#iamsamlproviderspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
