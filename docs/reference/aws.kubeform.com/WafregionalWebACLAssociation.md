---
title: WafregionalWebACLAssociation
menu:
  docs_{{ .version }}:
    identifier: wafregionalwebaclassociation-aws.kubeform.com
    name: WafregionalWebACLAssociation
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## WafregionalWebACLAssociation
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `WafregionalWebACLAssociation` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[WafregionalWebACLAssociationSpec](#WafregionalWebACLAssociationSpec)***||
| `status` | ***[WafregionalWebACLAssociationStatus](#WafregionalWebACLAssociationStatus)***||
## WafregionalWebACLAssociationSpec
##### (Appears on:[WafregionalWebACLAssociation](#WafregionalWebACLAssociation), [WafregionalWebACLAssociationStatus](#WafregionalWebACLAssociationStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `resourceArn` | ***string***||
| `webACLID` | ***string***||
## WafregionalWebACLAssociationStatus
##### (Appears on:[WafregionalWebACLAssociation](#WafregionalWebACLAssociation))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[WafregionalWebACLAssociationSpec](#WafregionalWebACLAssociationSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
