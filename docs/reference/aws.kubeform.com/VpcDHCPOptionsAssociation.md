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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcDHCPOptionsAssociationSpec](#VpcDHCPOptionsAssociationSpec)***||
| `status` | ***[VpcDHCPOptionsAssociationStatus](#VpcDHCPOptionsAssociationStatus)***||
## VpcDHCPOptionsAssociationSpec

Appears on:[VpcDHCPOptionsAssociation](#VpcDHCPOptionsAssociation), [VpcDHCPOptionsAssociationStatus](#VpcDHCPOptionsAssociationStatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dhcpOptionsID` | ***string***||
| `vpcID` | ***string***||
## VpcDHCPOptionsAssociationStatus

Appears on:[VpcDHCPOptionsAssociation](#VpcDHCPOptionsAssociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcDHCPOptionsAssociationSpec](#VpcDHCPOptionsAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
