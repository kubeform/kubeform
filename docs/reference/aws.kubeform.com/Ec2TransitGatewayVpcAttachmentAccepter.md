---
title: Ec2TransitGatewayVpcAttachmentAccepter
menu:
  docs_{{ .version }}:
    identifier: ec2transitgatewayvpcattachmentaccepter-aws.kubeform.com
    name: Ec2TransitGatewayVpcAttachmentAccepter
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Ec2TransitGatewayVpcAttachmentAccepter
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ec2TransitGatewayVpcAttachmentAccepter` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Ec2TransitGatewayVpcAttachmentAccepterSpec](#ec2transitgatewayvpcattachmentaccepterspec)***||
| `status` | ***[Ec2TransitGatewayVpcAttachmentAccepterStatus](#ec2transitgatewayvpcattachmentaccepterstatus)***||
## Ec2TransitGatewayVpcAttachmentAccepterSpec


Appears on:[Ec2TransitGatewayVpcAttachmentAccepter](#ec2transitgatewayvpcattachmentaccepter), [Ec2TransitGatewayVpcAttachmentAccepterStatus](#ec2transitgatewayvpcattachmentaccepterstatus)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dnsSupport` | ***string***| ***(Optional)*** |
| `ipv6Support` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `transitGatewayAttachmentID` | ***string***||
| `transitGatewayDefaultRouteTableAssociation` | ***bool***| ***(Optional)*** |
| `transitGatewayDefaultRouteTablePropagation` | ***bool***| ***(Optional)*** |
| `transitGatewayID` | ***string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
| `vpcOwnerID` | ***string***| ***(Optional)*** |
## Ec2TransitGatewayVpcAttachmentAccepterStatus


Appears on:[Ec2TransitGatewayVpcAttachmentAccepter](#ec2transitgatewayvpcattachmentaccepter)


| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Ec2TransitGatewayVpcAttachmentAccepterSpec](#ec2transitgatewayvpcattachmentaccepterspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
