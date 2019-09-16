---
title: VpcIpv4CIDRBlockAssociation
menu:
  docs_{{ .version }}:
    identifier: vpcipv4cidrblockassociation-aws.kubeform.com
    name: VpcIpv4CIDRBlockAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## VpcIpv4CIDRBlockAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `VpcIpv4CIDRBlockAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[VpcIpv4CIDRBlockAssociationSpec](#vpcipv4cidrblockassociationspec)***||
| `status` | ***[VpcIpv4CIDRBlockAssociationStatus](#vpcipv4cidrblockassociationstatus)***||
## VpcIpv4CIDRBlockAssociationSpec

Appears on:[VpcIpv4CIDRBlockAssociation](#vpcipv4cidrblockassociation), [VpcIpv4CIDRBlockAssociationStatus](#vpcipv4cidrblockassociationstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `cidrBlock` | ***string***||
| `vpcID` | ***string***||
## VpcIpv4CIDRBlockAssociationStatus

Appears on:[VpcIpv4CIDRBlockAssociation](#vpcipv4cidrblockassociation)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[VpcIpv4CIDRBlockAssociationSpec](#vpcipv4cidrblockassociationspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
