---
title: Route53ZoneAssociation
menu:
  docs_{{ .version }}:
    identifier: route53zoneassociation-aws.kubeform.com
    name: Route53ZoneAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Route53ZoneAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `Route53ZoneAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[Route53ZoneAssociationSpec](#Route53ZoneAssociationSpec)***||
| `status` | ***[Route53ZoneAssociationStatus](#Route53ZoneAssociationStatus)***||
## Route53ZoneAssociationSpec
##### (Appears on:[Route53ZoneAssociation](#Route53ZoneAssociation), [Route53ZoneAssociationStatus](#Route53ZoneAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `vpcID` | ***string***||
| `vpcRegion` | ***string***| ***(Optional)*** |
| `zoneID` | ***string***||
## Route53ZoneAssociationStatus
##### (Appears on:[Route53ZoneAssociation](#Route53ZoneAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[Route53ZoneAssociationSpec](#Route53ZoneAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
