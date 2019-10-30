---
title: Rdns
menu:
  docs_{{ .version }}:
    identifier: rdns-linode.kubeform.com
    name: Rdns
    parent: linode.kubeform.com-reference
    weight: 1
menu_name: docs_{{ .version }}
section_menu_id: reference
---

## Rdns
| Field | Type | Description |
| ------ | ----- | ----------- |
| `apiVersion` | string | `linode.kubeform.com/v1alpha1` |
|    `kind` | string | `Rdns` |
| `metadata` | ***[Kubernetes meta/v1.ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#objectmeta-v1-meta)***|Refer to the Kubernetes API documentation for the fields of the `metadata` field.|
| `spec` | ***[RdnsSpec](#rdnsspec)***||
| `status` | ***[RdnsStatus](#rdnsstatus)***||
## Phase(`string` alias)

Appears on:[RdnsStatus](#rdnsstatus)

## RdnsSpec

Appears on:[Rdns](#rdns), [RdnsStatus](#rdnsstatus)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `providerRef` | ***[Kubernetes core/v1.LocalObjectReference](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.13/#localobjectreference-v1-core)***||
| `id` | ***string***||
| `address` | ***string***|The public Linode IPv4 or IPv6 address to operate on.|
| `rdns` | ***string***|The reverse DNS assigned to this address. For public IPv4 addresses, this will be set to a default value provided by Linode if not explicitly set.|
## RdnsStatus

Appears on:[Rdns](#rdns)

| Field | Type | Description |
| ------ | ----- | ----------- |
| `observedGeneration` | ***int64***| ***(Optional)*** Resource generation, which is updated on mutation by the API Server.|
| `output` | ***[RdnsSpec](#rdnsspec)***| ***(Optional)*** |
| `state` | ***kubeform.dev/kubeform/apis/base/v1alpha1.State***| ***(Optional)*** |
| `phase` | ***[Phase](#phase)***| ***(Optional)*** |
---
