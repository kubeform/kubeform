---
title: IamSamlProvider
menu:
  docs_v0.0.1:
    identifier: iamsamlprovider-aws.kubeform.com
    name: IamSamlProvider
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IamSamlProvider
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamSamlProvider` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamSamlProviderSpec](#IamSamlProviderSpec)***||
| `status` | ***[IamSamlProviderStatus](#IamSamlProviderStatus)***||
## IamSamlProviderSpec
##### (Appears on:[IamSamlProvider](#IamSamlProvider), [IamSamlProviderStatus](#IamSamlProviderStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `samlMetadataDocument` | ***string***||
| `validUntil` | ***string***| ***(Optional)*** |
## IamSamlProviderStatus
##### (Appears on:[IamSamlProvider](#IamSamlProvider))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamSamlProviderSpec](#IamSamlProviderSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
