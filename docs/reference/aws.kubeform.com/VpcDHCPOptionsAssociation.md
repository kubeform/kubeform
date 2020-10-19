---
title: VpcDHCPOptionsAssociation
menu:
  docs_{{ .version }}:
    identifier: vpcdhcpoptionsassociation-aws.kubeform.com
    name: VpcDHCPOptionsAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpcDHCPOptionsAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpcDHCPOptionsAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcDHCPOptionsAssociationSpec](#vpcdhcpoptionsassociationspec)***||
| `status` | ***[VpcDHCPOptionsAssociationStatus](#vpcdhcpoptionsassociationstatus)***||
## Phase(`string` alias)

Appears on:[VpcDHCPOptionsAssociationStatus](#vpcdhcpoptionsassociationstatus)

## VpcDHCPOptionsAssociationSpec

Appears on:[VpcDHCPOptionsAssociation](#vpcdhcpoptionsassociation), [VpcDHCPOptionsAssociationStatus](#vpcdhcpoptionsassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dhcpOptionsID` | ***string***||
| `vpcID` | ***string***||
## VpcDHCPOptionsAssociationStatus

Appears on:[VpcDHCPOptionsAssociation](#vpcdhcpoptionsassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcDHCPOptionsAssociationSpec](#vpcdhcpoptionsassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
