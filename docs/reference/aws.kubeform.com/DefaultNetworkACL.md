---
title: DefaultNetworkACL
menu:
  docs_{{ .version }}:
    identifier: defaultnetworkacl-aws.kubeform.com
    name: DefaultNetworkACL
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## DefaultNetworkACL
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `DefaultNetworkACL` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DefaultNetworkACLSpec](#DefaultNetworkACLSpec)***||
| `status` | ***[DefaultNetworkACLStatus](#DefaultNetworkACLStatus)***||
## DefaultNetworkACLSpec
##### (Appears on:[DefaultNetworkACL](#DefaultNetworkACL), [DefaultNetworkACLStatus](#DefaultNetworkACLStatus))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `defaultNetworkACLID` | ***string***||
| `egress` | ***[[]DefaultNetworkACLSpecEgress](#DefaultNetworkACLSpecEgress)***| ***(Optional)*** |
| `ingress` | ***[[]DefaultNetworkACLSpecIngress](#DefaultNetworkACLSpecIngress)***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## DefaultNetworkACLSpecEgress
##### (Appears on:[DefaultNetworkACLSpec](#DefaultNetworkACLSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `cidrBlock` | ***string***| ***(Optional)*** |
| `fromPort` | ***int***||
| `icmpCode` | ***int***| ***(Optional)*** |
| `icmpType` | ***int***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `ruleNo` | ***int***||
| `toPort` | ***int***||
## DefaultNetworkACLSpecIngress
##### (Appears on:[DefaultNetworkACLSpec](#DefaultNetworkACLSpec))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `cidrBlock` | ***string***| ***(Optional)*** |
| `fromPort` | ***int***||
| `icmpCode` | ***int***| ***(Optional)*** |
| `icmpType` | ***int***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `ruleNo` | ***int***||
| `toPort` | ***int***||
## DefaultNetworkACLStatus
##### (Appears on:[DefaultNetworkACL](#DefaultNetworkACL))
| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DefaultNetworkACLSpec](#DefaultNetworkACLSpec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis.State***| ***(Optional)*** |
---
