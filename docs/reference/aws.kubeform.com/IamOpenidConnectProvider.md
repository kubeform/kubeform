---
title: IamOpenidConnectProvider
menu:
  docs_v0.0.1:
    identifier: iamopenidconnectprovider-aws.kubeform.com
    name: IamOpenidConnectProvider
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_v0.0.1
section_menu_id: reference
---

## IamOpenidConnectProvider
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamOpenidConnectProvider` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamOpenidConnectProviderSpec](#IamOpenidConnectProviderSpec)***||
| `status` | ***[IamOpenidConnectProviderStatus](#IamOpenidConnectProviderStatus)***||
## IamOpenidConnectProviderSpec
##### (Appears on:[IamOpenidConnectProvider](#IamOpenidConnectProvider), [IamOpenidConnectProviderStatus](#IamOpenidConnectProviderStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `clientIDList` | ***[]string***||
| `thumbprintList` | ***[]string***||
| `url` | ***string***||
## IamOpenidConnectProviderStatus
##### (Appears on:[IamOpenidConnectProvider](#IamOpenidConnectProvider))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamOpenidConnectProviderSpec](#IamOpenidConnectProviderSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
