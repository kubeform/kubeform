---
title: DxHostedPublicVirtualInterfaceAccepter
menu:
  docs_{{ .version }}:
    identifier: dxhostedpublicvirtualinterfaceaccepter-aws.kubeform.com
    name: DxHostedPublicVirtualInterfaceAccepter
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DxHostedPublicVirtualInterfaceAccepter
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxHostedPublicVirtualInterfaceAccepter` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxHostedPublicVirtualInterfaceAccepterSpec](#dxhostedpublicvirtualinterfaceaccepterspec)***||
| `status` | ***[DxHostedPublicVirtualInterfaceAccepterStatus](#dxhostedpublicvirtualinterfaceaccepterstatus)***||
## DxHostedPublicVirtualInterfaceAccepterSpec

Appears on:[DxHostedPublicVirtualInterfaceAccepter](#dxhostedpublicvirtualinterfaceaccepter), [DxHostedPublicVirtualInterfaceAccepterStatus](#dxhostedpublicvirtualinterfaceaccepterstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `arn` | ***string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `virtualInterfaceID` | ***string***||
## DxHostedPublicVirtualInterfaceAccepterStatus

Appears on:[DxHostedPublicVirtualInterfaceAccepter](#dxhostedpublicvirtualinterfaceaccepter)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxHostedPublicVirtualInterfaceAccepterSpec](#dxhostedpublicvirtualinterfaceaccepterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
