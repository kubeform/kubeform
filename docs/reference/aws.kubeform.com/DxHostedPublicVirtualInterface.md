---
title: DxHostedPublicVirtualInterface
menu:
  docs_{{ .version }}:
    identifier: dxhostedpublicvirtualinterface-aws.kubeform.com
    name: DxHostedPublicVirtualInterface
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DxHostedPublicVirtualInterface
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxHostedPublicVirtualInterface` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxHostedPublicVirtualInterfaceSpec](#dxhostedpublicvirtualinterfacespec)***||
| `status` | ***[DxHostedPublicVirtualInterfaceStatus](#dxhostedpublicvirtualinterfacestatus)***||
## DxHostedPublicVirtualInterfaceSpec

Appears on:[DxHostedPublicVirtualInterface](#dxhostedpublicvirtualinterface), [DxHostedPublicVirtualInterfaceStatus](#dxhostedpublicvirtualinterfacestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `addressFamily` | ***string***||
| `amazonAddress` | ***string***| ***(Optional)*** |
| `arn` | ***string***| ***(Optional)*** |
| `awsDevice` | ***string***| ***(Optional)*** |
| `bgpAsn` | ***int64***||
| `bgpAuthKey` | ***string***| ***(Optional)*** |
| `connectionID` | ***string***||
| `customerAddress` | ***string***| ***(Optional)*** |
| `name` | ***string***||
| `ownerAccountID` | ***string***||
| `routeFilterPrefixes` | ***[]string***||
| `vlan` | ***int64***||
## DxHostedPublicVirtualInterfaceStatus

Appears on:[DxHostedPublicVirtualInterface](#dxhostedpublicvirtualinterface)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxHostedPublicVirtualInterfaceSpec](#dxhostedpublicvirtualinterfacespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DxHostedPublicVirtualInterfaceStatus](#dxhostedpublicvirtualinterfacestatus)

---
