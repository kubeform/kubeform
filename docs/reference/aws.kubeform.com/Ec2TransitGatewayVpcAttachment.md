---
title: Ec2TransitGatewayVpcAttachment
menu:
  docs_{{ .version }}:
    identifier: ec2transitgatewayvpcattachment-aws.kubeform.com
    name: Ec2TransitGatewayVpcAttachment
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Ec2TransitGatewayVpcAttachment
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Ec2TransitGatewayVpcAttachment` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Ec2TransitGatewayVpcAttachmentSpec](#ec2transitgatewayvpcattachmentspec)***||
| `status` | ***[Ec2TransitGatewayVpcAttachmentStatus](#ec2transitgatewayvpcattachmentstatus)***||
## Ec2TransitGatewayVpcAttachmentSpec

Appears on:[Ec2TransitGatewayVpcAttachment](#ec2transitgatewayvpcattachment), [Ec2TransitGatewayVpcAttachmentStatus](#ec2transitgatewayvpcattachmentstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `dnsSupport` | ***string***| ***(Optional)*** |
| `ipv6Support` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***||
| `tags` | ***map[string]string***| ***(Optional)*** |
| `transitGatewayDefaultRouteTableAssociation` | ***bool***| ***(Optional)*** |
| `transitGatewayDefaultRouteTablePropagation` | ***bool***| ***(Optional)*** |
| `transitGatewayID` | ***string***||
| `vpcID` | ***string***||
| `vpcOwnerID` | ***string***| ***(Optional)*** |
## Ec2TransitGatewayVpcAttachmentStatus

Appears on:[Ec2TransitGatewayVpcAttachment](#ec2transitgatewayvpcattachment)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Ec2TransitGatewayVpcAttachmentSpec](#ec2transitgatewayvpcattachmentspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[Ec2TransitGatewayVpcAttachmentStatus](#ec2transitgatewayvpcattachmentstatus)

---
