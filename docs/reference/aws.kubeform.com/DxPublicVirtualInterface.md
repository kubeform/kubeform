---
title: DxPublicVirtualInterface
menu:
  docs_{{ .version }}:
    identifier: dxpublicvirtualinterface-aws.kubeform.com
    name: DxPublicVirtualInterface
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DxPublicVirtualInterface
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DxPublicVirtualInterface` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DxPublicVirtualInterfaceSpec](#dxpublicvirtualinterfacespec)***||
| `status` | ***[DxPublicVirtualInterfaceStatus](#dxpublicvirtualinterfacestatus)***||
## DxPublicVirtualInterfaceSpec

Appears on:[DxPublicVirtualInterface](#dxpublicvirtualinterface), [DxPublicVirtualInterfaceStatus](#dxpublicvirtualinterfacestatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
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
| `routeFilterPrefixes` | ***[]string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vlan` | ***int64***||
## DxPublicVirtualInterfaceStatus

Appears on:[DxPublicVirtualInterface](#dxpublicvirtualinterface)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DxPublicVirtualInterfaceSpec](#dxpublicvirtualinterfacespec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DxPublicVirtualInterfaceStatus](#dxpublicvirtualinterfacestatus)

---
