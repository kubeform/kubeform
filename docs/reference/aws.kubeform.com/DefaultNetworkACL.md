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
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[DefaultNetworkACLSpec](#defaultnetworkaclspec)***||
| `status` | ***[DefaultNetworkACLStatus](#defaultnetworkaclstatus)***||
## DefaultNetworkACLSpec

Appears on:[DefaultNetworkACL](#defaultnetworkacl), [DefaultNetworkACLStatus](#defaultnetworkaclstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `defaultNetworkACLID` | ***string***||
| `egress` | ***[[]DefaultNetworkACLSpecEgress](#defaultnetworkaclspecegress)***| ***(Optional)*** |
| `ingress` | ***[[]DefaultNetworkACLSpecIngress](#defaultnetworkaclspecingress)***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***| ***(Optional)*** |
## DefaultNetworkACLSpecEgress

Appears on:[DefaultNetworkACLSpec](#defaultnetworkaclspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `cidrBlock` | ***string***| ***(Optional)*** |
| `fromPort` | ***int64***||
| `icmpCode` | ***int64***| ***(Optional)*** |
| `icmpType` | ***int64***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `ruleNo` | ***int64***||
| `toPort` | ***int64***||
## DefaultNetworkACLSpecIngress

Appears on:[DefaultNetworkACLSpec](#defaultnetworkaclspec)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `action` | ***string***||
| `cidrBlock` | ***string***| ***(Optional)*** |
| `fromPort` | ***int64***||
| `icmpCode` | ***int64***| ***(Optional)*** |
| `icmpType` | ***int64***| ***(Optional)*** |
| `ipv6CIDRBlock` | ***string***| ***(Optional)*** |
| `protocol` | ***string***||
| `ruleNo` | ***int64***||
| `toPort` | ***int64***||
## DefaultNetworkACLStatus

Appears on:[DefaultNetworkACL](#defaultnetworkacl)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[DefaultNetworkACLSpec](#defaultnetworkaclspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[DefaultNetworkACLStatus](#defaultnetworkaclstatus)

---
