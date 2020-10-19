---
title: IamOpenidConnectProvider
menu:
  docs_{{ .version }}:
    identifier: iamopenidconnectprovider-aws.kubeform.com
    name: IamOpenidConnectProvider
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## IamOpenidConnectProvider
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `IamOpenidConnectProvider` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[IamOpenidConnectProviderSpec](#iamopenidconnectproviderspec)***||
| `status` | ***[IamOpenidConnectProviderStatus](#iamopenidconnectproviderstatus)***||
## IamOpenidConnectProviderSpec

Appears on:[IamOpenidConnectProvider](#iamopenidconnectprovider), [IamOpenidConnectProviderStatus](#iamopenidconnectproviderstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `clientIDList` | ***[]string***||
| `thumbprintList` | ***[]string***||
| `url` | ***string***||
## IamOpenidConnectProviderStatus

Appears on:[IamOpenidConnectProvider](#iamopenidconnectprovider)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[IamOpenidConnectProviderSpec](#iamopenidconnectproviderspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[IamOpenidConnectProviderStatus](#iamopenidconnectproviderstatus)

---
