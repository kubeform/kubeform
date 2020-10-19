---
title: NetworkACL
menu:
  docs_{{ .version }}:
    identifier: networkacl-aws.kubeform.com
    name: NetworkACL
    parent: aws.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## NetworkACL
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `aws.kubeform.com/v1alpha1` |
|    `kind` | string | `NetworkACL` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[NetworkACLSpec](#networkaclspec)***||
| `status` | ***[NetworkACLStatus](#networkaclstatus)***||
## NetworkACLSpec

Appears on:[NetworkACL](#networkacl), [NetworkACLStatus](#networkaclstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://v1-18.docs.kubernetes.io/docs/reference/generated/kubernetes-api/v1.18/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `egress` | ***[[]NetworkACLSpecEgress](#networkaclspecegress)***| ***(Optional)*** |
| `ingress` | ***[[]NetworkACLSpecIngress](#networkaclspecingress)***| ***(Optional)*** |
| `ownerID` | ***string***| ***(Optional)*** |
| `subnetIDS` | ***[]string***| ***(Optional)*** |
| `tags` | ***map[string]string***| ***(Optional)*** |
| `vpcID` | ***string***||
## NetworkACLSpecEgress

Appears on:[NetworkACLSpec](#networkaclspec)

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
## NetworkACLSpecIngress

Appears on:[NetworkACLSpec](#networkaclspec)

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
## NetworkACLStatus

Appears on:[NetworkACL](#networkacl)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[NetworkACLSpec](#networkaclspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
## Phase(`string` alias)

Appears on:[NetworkACLStatus](#networkaclstatus)

---
